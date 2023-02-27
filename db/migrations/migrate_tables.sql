DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) DEFAULT '',
    email varchar(100) default '',
    address varchar(100)
);