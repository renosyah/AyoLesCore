
DROP TABLE IF EXISTS classroom_certificate CASCADE;
DROP TABLE IF EXISTS classroom_exam_progress CASCADE;
DROP TABLE IF EXISTS classroom_progress CASCADE;
DROP TABLE IF EXISTS classroom CASCADE;
DROP TABLE IF EXISTS course_exam_solution CASCADE;
DROP TABLE IF EXISTS course_exam_answer CASCADE;
DROP TABLE IF EXISTS course_exam CASCADE;
DROP TABLE IF EXISTS course_material_detail CASCADE;
DROP TABLE IF EXISTS course_material CASCADE;
DROP TABLE IF EXISTS course_qualification CASCADE;
DROP TABLE IF EXISTS course_detail CASCADE;
DROP TABLE IF EXISTS course CASCADE;
DROP TABLE IF EXISTS banner CASCADE;
DROP TABLE IF EXISTS course_category CASCADE;
DROP TABLE IF EXISTS teacher CASCADE;
DROP TABLE IF EXISTS student CASCADE;

CREATE TABLE student (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL DEFAULT '',
    email STRING NOT NULL DEFAULT '',
    password STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE teacher (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL DEFAULT '',
    email STRING NOT NULL DEFAULT '',
    password STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE course_category (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE banner (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    title STRING NOT NULL DEFAULT '',
    content STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE course (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_name STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    teacher_id UUID NOT NULL REFERENCES teacher (id),
    category_id UUID NOT NULL REFERENCES course_category (id),
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE course_detail (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES course (id),
    overview_text STRING NOT NULL DEFAULT '',
    description_text STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE course_qualification (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES course (id),
    course_level STRING NOT NULL DEFAULT '',
    min_score INT NOT NULL DEFAULT 0,
    course_material_total INT NOT NULL DEFAULT 0,
    course_exam_total INT NOT NULL DEFAULT 0,
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE course_material (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES course (id),
    material_index INT NOT NULL DEFAULT 0,
    title STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
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
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE course_exam (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES course (id),
    type_exam INT NOT NULL DEFAULT 0,
    exam_index INT NOT NULL DEFAULT 0,
    text STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE course_exam_answer (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_exam_id UUID NOT NULL REFERENCES course_exam (id),
    type_answer INT NOT NULL DEFAULT 0,
    label STRING NOT NULL DEFAULT '',
    text STRING NOT NULL DEFAULT '',
    image_url STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE course_exam_solution (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_exam_id UUID NOT NULL REFERENCES course_exam (id),
    course_exam_answer_id UUID NOT NULL REFERENCES course_exam_answer (id),
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE classroom (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES course (id),
    student_id UUID NOT NULL REFERENCES student (id),
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE classroom_progress (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    classroom_id UUID NOT NULL REFERENCES classroom (id),
    course_material_id UUID NOT NULL REFERENCES course_material (id),
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE classroom_exam_progress (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    classroom_id UUID NOT NULL REFERENCES classroom (id),
    course_exam_id UUID NOT NULL REFERENCES course_exam (id),
    course_exam_answer_id UUID NOT NULL REFERENCES course_exam_answer (id),
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);


CREATE TABLE classroom_certificate(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    classroom_id UUID NOT NULL REFERENCES classroom (id),
    hash_id STRING NOT NULL DEFAULT '',
    create_at TIMESTAMPTZ NOT NULL DEFAULT now()::TIMESTAMPTZ,
    PRIMARY KEY (id)
);

