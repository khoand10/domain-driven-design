DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) DEFAULT '',
    email varchar(100) default '',
    address varchar(100)
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       name VARCHAR(100),
       email varchar(100),
       password varchar(200),
       active BIT
);