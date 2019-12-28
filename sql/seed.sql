INSERT INTO student (name,email,password) VALUES ('reno','reno@gmail.com','12345');
INSERT INTO student (name,email,password) VALUES ('rikka','rikka@gmail.com','12345');
INSERT INTO student (name,email,password) VALUES ('agus','agus@gmail.com','12345');
INSERT INTO student (name,email,password) VALUES ('surya','surya@gmail.com','12345');

INSERT INTO teacher (id,name,email,password) VALUES ('73aa9774-5f93-40b4-b510-4e465f97cfcd','prof agus','profagus@gmail.com','12345');
INSERT INTO teacher (id,name,email,password) VALUES ('a46a67f6-4208-4d68-8df4-1894031664b0','prof surya','profsurya@gmail.com','12345');

INSERT INTO course_category (id,name,image_url) VALUES ('0679c05c-db7b-4278-968a-3ec0388673c8','Math','data/category/math.png');
INSERT INTO course_category (id,name,image_url) VALUES ('d2ab9a06-e866-4db0-a5d5-7ef31c4b25f5','Language','data/category/language.png');
INSERT INTO course_category (id,name,image_url) VALUES ('57fa4880-7dbf-4b27-ba75-6d554de77a89','Art','data/category/art.png');
INSERT INTO course_category (id,name,image_url) VALUES ('c6fef7b3-3bc1-4068-b00a-b58d0ffdb699','Science','data/category/science.png');

INSERT INTO banner (title,content,image_url) VALUES ('End Year Discount!','50 % off just in ayoles!','data/banner/end_year.png');
INSERT INTO banner (title,content,image_url) VALUES ('Enroll More','its more fun learn in ayoles','data/banner/more_enroll.png');
INSERT INTO banner (title,content,image_url) VALUES ('Be the best','take more free course to get started','data/banner/free.png');

INSERT INTO course (id,course_name,teacher_id,category_id) VALUES ('2e847a03-5209-4d2b-9e37-b88e461e9c41','Data Science','73aa9774-5f93-40b4-b510-4e465f97cfcd','c6fef7b3-3bc1-4068-b00a-b58d0ffdb699');
INSERT INTO course (id,course_name,teacher_id,category_id) VALUES ('f0c68980-70a8-492d-8701-9fb15086dd44','Basic Math','73aa9774-5f93-40b4-b510-4e465f97cfcd','0679c05c-db7b-4278-968a-3ec0388673c8');
INSERT INTO course (id,course_name,teacher_id,category_id) VALUES ('d0941ffe-d1bd-415b-b85e-f238ee7ca3ac','Art of Color','a46a67f6-4208-4d68-8df4-1894031664b0','57fa4880-7dbf-4b27-ba75-6d554de77a89');
INSERT INTO course (id,course_name,teacher_id,category_id) VALUES ('2071a647-23bf-45d0-80a9-46328b1310d2','Russian is fun','73aa9774-5f93-40b4-b510-4e465f97cfcd','d2ab9a06-e866-4db0-a5d5-7ef31c4b25f5');
INSERT INTO course (id,course_name,teacher_id,category_id) VALUES ('978cda00-30d9-4229-86e4-0f376e970966','Expert Math','a46a67f6-4208-4d68-8df4-1894031664b0','0679c05c-db7b-4278-968a-3ec0388673c8');
INSERT INTO course (id,course_name,teacher_id,category_id) VALUES ('cfc67c74-6519-4da6-aead-cd529de39513','Human Anatomy','73aa9774-5f93-40b4-b510-4e465f97cfcd','c6fef7b3-3bc1-4068-b00a-b58d0ffdb699');

