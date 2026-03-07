package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgilIw/koda-b6-backend2/internal/model"
)

type UserRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewUserRepository(db *pgxpool.Pool, rdb *redis.Client) *UserRepository {
	return &UserRepository{
		db:  db,
		rdb: rdb,
	}
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]model.UserDb, error) {
	cacheKey := "users:all"
	valueCache, err := u.rdb.Get(ctx, cacheKey).Result()

	// 	Fungsi	Tujuan
	// json.Marshal	Go struct → JSON
	// json.Unmarshal	JSON → Go struct

	if err == redis.Nil {
		query := `
		select id, fullname, email, password, picture, phone, address, role, created_at 
		from users`
		rows, err := u.db.Query(ctx, query)

		if err != nil {
			return nil, err
		}
		fmt.Println("ERROR:", err)
		defer rows.Close()

		users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.UserDb])

		if err != nil {
			return nil, err
		}

		// ubah data user ke json
		val, err := json.Marshal(users)

		u.rdb.Set(ctx, cacheKey, string(val), time.Minute*15)

		return users, nil
	} else if err != nil {
		return nil, err
	} else {
		users := []model.UserDb{}
		if err := json.Unmarshal([]byte(valueCache), &users); err != nil {
			return nil, err
		}
		return users, nil
	}

}

func (u *UserRepository) GetUserById(ctx context.Context, id int) (model.UserDb, error) {
	cacheKey := fmt.Sprintf("users:%d", id)

	valueCache, err := u.rdb.Get(ctx, cacheKey).Result()

	if err == redis.Nil {
		var user model.UserDb

		query := `select id, fullname, email, password, picture, phone, address, role from users where id = $1`

		row := u.db.QueryRow(ctx, query, id)
		log.Println(err)
		err := row.Scan(
			&user.Id,
			&user.Fullname,
			&user.Email,
			&user.Password,
			&user.Picture,
			&user.Phone,
			&user.Address,
			&user.Role,
		)

		if err != nil {
			return model.UserDb{}, err
		}

		val, err := json.Marshal(user)
		if err != nil {
			return model.UserDb{}, err
		}

		u.rdb.Set(ctx, cacheKey, val, time.Minute*15)

		return user, nil
	}

	if err != nil {
		return model.UserDb{}, err
	}
	var user model.UserDb
	if err := json.Unmarshal([]byte(valueCache), &user); err != nil {
		return model.UserDb{}, err
	}

	return user, nil
}

func (u *UserRepository) CreateUser(ctx context.Context, fullname, email, password string) model.UserDb {
	query := `insert into users(fullname, email, password) values($1, $2, $3) Returning id, email, password, created_at`

	var newUser model.UserDb

	row := u.db.QueryRow(ctx, query, fullname, email, password)
	row.Scan(
		&newUser.Id,
		&newUser.Email,
		&newUser.Password,
		&newUser.Created_At,
	)

	return newUser
}

func (u *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
