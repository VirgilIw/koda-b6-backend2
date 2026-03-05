# MINITASK BACKEND 2

API sederhana menggunakan **Golang** dan **Gin** untuk mengelola **User** dan **Product**.

Data masih disimpan sementara di **memory**, sehingga akan hilang jika server direstart.

---

# Run

```bash
go run cmd/main.go
```

Server berjalan di:

```
http://localhost:8888
```

---

# User Endpoints

### Get All Users

```
GET /users
```

---

### Get User

```
GET /users/:id
```

---

### Create User

```
POST /users
```

Body

```json
{
  "email": "user@mail.com",
  "password": "123456"
}
```

---

### Update User

```
PATCH /users/:id
```

---

### Delete User

```
DELETE /users/:id
```

---

# Product Endpoints

### Get All Products

```
GET /products
```

---

### Get Product By ID

```
GET /products/:id
```

---

### Create Product

```
POST /products
```

Body

```json
{
  "product_name": "Americano",
  "rating": 4.7
}
```

---

### Update Product

```
PATCH /products/:id
```

Body

```json
{
  "product_name": "Latte",
  "rating": 4.8
}
```

---

### Delete Product

```
DELETE /products/:id
```

---

# Example Requests (.http)

```http
### Get Products
GET http://localhost:8888/products


### Create Product
POST http://localhost:8888/products
Content-Type: application/json

{
  "product_name": "Espresso",
  "rating": 4.9
}


### Get Product By ID
GET http://localhost:8888/products/1


### Update Product
PATCH http://localhost:8888/products/1
Content-Type: application/json

{
  "product_name": "Cappuccino",
  "rating": 4.8
}


### Delete Product
DELETE http://localhost:8888/products/1
```

---

Flow aplikasi:

```
router → di(dependency injection) → Handler → Service → Repository → Model
```

---

# Note

Data hanya disimpan di **memory (slice)** dan akan hilang jika server direstart.
