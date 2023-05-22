CREATE DATABASE student_service_db;
CREATE TABLE students
(
    id         serial NOT NULL PRIMARY KEY,
    first_name character varying(250),
    last_name  character varying(250),
    age        integer
);