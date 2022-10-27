CREATE DATABASE library;

USE DATABASE library;

CREATE TABLE Book (
    id SERIAL PRIMARY KEY,
    name varchar(256),
    author varchar(256),
    synopsis text,
    launch_date date,
    copy_qnt int
);