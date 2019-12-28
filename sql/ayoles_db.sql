DROP DATABASE IF EXISTS ayoles_db CASCADE;

CREATE DATABASE ayoles_db;

USE ayoles_db;

CREATE TABLE student (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL DEFAULT '',
    email STRING NOT NULL DEFAULT '',
    password STRING NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TABLE teacher (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL DEFAULT '',
    email STRING NOT NULL DEFAULT '',
    password STRING NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TABLE course_category (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TABLE banner (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    title STRING NOT NULL DEFAULT '',
    content STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TABLE course (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_name STRING NOT NULL DEFAULT '',
    teacher_id UUID NOT NULL REFERENCES teacher (id),
    category_id UUID NOT NULL REFERENCES course_category (id),
    PRIMARY KEY (id)
);

