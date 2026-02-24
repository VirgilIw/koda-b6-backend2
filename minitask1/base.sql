CREATE TABLE "books" (
    "id" serial primary key,
    "book_name" varchar(255) not null,
    "book_author" varchar(255) not null,
    "publication_year" int not null
);
-- add new column
ALTER TABLE "books"
ADD COLUMN "category_id" INT;

ALTER TABLE "books"
ADD COLUMN "bookshelf_id" INT;

-- add new contraints
alter table "books"
add constraint "books_category"
foreign key ("category_id")
REFERENCES "category" ("id");

alter table "books"
add constraint "books_bookshelf"
foreign key ("bookshelf_id")
REFERENCES "bookshelf" ("id");

alter table "books"
drop column "book_author";

create table "authors" (
    "id" int generated always as identity primary key,
    "name" varchar(255) not null
)

create table "book_authors" (
    "id" int generated always as identity primary key,
    "book_id" int,
    "author_id" int,
    foreign key ("book_id") references "books"("id"),
    foreign key ("author_id") references "authors"("id"),
    unique("book_id", "author_id")
)

SELECT "id", "book_name", "book_author", "publication_year" from "books";

INSERT INTO books 
(book_name, book_author, publication_year, category_id, bookshelf_id)
VALUES
('The Great Adventure', 'John Smith', 2015, 1, 1),
('Science of Space', 'Neil Carter', 2018, 3, 2),
('World History 101', 'Anna Lee', 2012, 4, 3),
('Learning SQL', 'Mark Johnson', 2020, 5, 4),
('Mystery Island', 'Sarah Connor', 2016, 1, 5),
('Data Structures', 'Robert King', 2019, 5, 1),
('Modern Physics', 'Albert West', 2017, 3, 2),
('Biography of Tesla', 'Emily Stone', 2014, 2, 3),
('Ancient Civilizations', 'David Brown', 2011, 4, 4),
('Creative Writing', 'Laura White', 2021, 2, 5);
-- 
create table "category"(
    "id" int generated always as identity primary key,
    "category_name" varchar(255) not null
);

SELECT "id", "category_name" from "category"

INSERT INTO category (category_name) VALUES
('Fiction'),
('Non-Fiction'),
('Science'),
('History'),
('Technology'),
('Philosophy'),
('Education'),
('Art'),
('Business'),
('Health');


-- BOOKSHELF{
--     INT ID PK
--     STRING SECTION
-- }

create table "bookshelf" (
    "id" int generated always as identity primary key,
    "section" varchar(10)
);

SELECT "id", "section" from "bookshelf";

INSERT INTO bookshelf (section) VALUES
('A1'),
('A2'),
('B1'),
('B2'),
('C1'),
('C2'),
('D1'),
('D2'),
('E1'),
('E2'),
('F1');
-- OFFICER {
--     INT ID PK
--     STRING OFFICER_NAME
-- }

create table "officer" (
    "id" int generated always as identity primary key,
    "officer_name" varchar(255)
);

SELECT "id", "officer_name" from "officer";

INSERT INTO officer (officer_name) VALUES
('Andi'),
('Budi'),
('Citra'),
('Dewi'),
('Anda'),
('Beda'),
('Citrop'),
('Citrol'),
('Citros'),
('Citroa'),
('Citrov');

-- BORROWER {
--     INT ID PK
--     STRING BORROWER_NAME
-- }

create table "borrower" (
    "id" int generated always as identity primary key,
    "borrower_name" varchar(255)
);

SELECT "id", "borrower_name" from "borrower";

INSERT INTO borrower (borrower_name) VALUES
('Rina'),
('Dimas'),
('Sari'),
('Fajar'),
('Tono'),
('Lina'),
('Bagas'),
('Nina'),
('Rafi'),
('Putri');

-- BORROW_TRANSACTION {
--     INT ID PK
--     DATE BORROW_DATE
--     DATE RETURN_DATE
--     INT BORROWER_ID FK
--     INT BOOK_ID FK
--     INT OFFICER_ID FK
-- }

create table "borrow_transaction" (
    "id" int generated always as identity primary key,
    "borrow_date" timestamp not null,
    "return_date" timestamp not null,
    "borrower_id" int,
    "book_id" int,
    "officer_id" int,
    foreign key ("borrower_id") references "borrower" ("id"),
    foreign key ("book_id") references "books" ("id"),
    foreign key ("officer_id") references "officer" ("id")
)

-- 
INSERT INTO borrow_transaction
(borrow_date, return_date, borrower_id, book_id, officer_id)
VALUES
(NOW(), NOW() + INTERVAL '7 days', 1, 1, 1),
(NOW(), NOW() + INTERVAL '5 days', 2, 2, 2),
(NOW(), NOW() + INTERVAL '10 days', 3, 3, 3),
(NOW(), NOW() + INTERVAL '3 days', 4, 4, 4),
(NOW(), NOW() + INTERVAL '14 days', 5, 5, 5),
(NOW(), NOW() + INTERVAL '6 days', 6, 6, 6),
(NOW(), NOW() + INTERVAL '8 days', 7, 7, 7),
(NOW(), NOW() + INTERVAL '9 days', 8, 8, 8),
(NOW(), NOW() + INTERVAL '4 days', 9, 9, 9),
(NOW(), NOW() + INTERVAL '12 days', 10, 10, 10);

SELECT * FROM category;
SELECT * FROM bookshelf;
SELECT * FROM officer;
SELECT * FROM borrower;
SELECT * FROM books;
SELECT * FROM borrow_transaction;
