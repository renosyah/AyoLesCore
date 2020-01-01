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
    image_url STRING NOT NULL DEFAULT '',
    teacher_id UUID NOT NULL REFERENCES teacher (id),
    category_id UUID NOT NULL REFERENCES course_category (id),
    PRIMARY KEY (id)
);

CREATE TABLE course_detail (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES course (id),
    overview_text STRING NOT NULL DEFAULT '',
    description_text STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TABLE course_material (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES course (id),
    material_index INT NOT NULL DEFAULT 0,
    title STRING NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TABLE course_material_detail (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_material_id UUID NOT NULL REFERENCES course_material (id),
    position INT NOT NULL DEFAULT 0,
    title STRING NOT NULL DEFAULT '',
    type_material INT NOT NULL DEFAULT 0,
    content STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TABLE classroom (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES course (id),
    student_id UUID NOT NULL REFERENCES student (id),
    date_add TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);

CREATE TABLE classroom_progress (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    classroom_id UUID NOT NULL REFERENCES classroom (id),
    course_material_id UUID NOT NULL REFERENCES course_material (id),
    date_add TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);
