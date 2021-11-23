INSERT INTO course (id,course_name,teacher_id,category_id,image_url) VALUES ('8c6c787f-1ec1-4be2-b7f0-161e0c3a5345','Inggris','73aa9774-5f93-40b4-b510-4e465f97cfcd','0679c05c-db7b-4278-968a-3ec0388673c8','http://192.168.100.62:8080/data/course/bahasa_inggris.png');


INSERT INTO course_qualification (course_id,course_level,min_score,course_material_total,course_exam_total) VALUES ('8c6c787f-1ec1-4be2-b7f0-161e0c3a5345','SD kelas 6',3,3,5);


INSERT INTO course_detail (course_id,overview_text,description_text,image_url) VALUES ('8c6c787f-1ec1-4be2-b7f0-161e0c3a5345','English Book','Mengingat era Globalisasi saat ini, kebutuhan akan Bahasa Inggris memang sangatlah tinggi, seperti yang kita semua ketahui, bahasa pengantar yang dipergunakan masyarakat global sekarang adalah Bahasa Inggris, untuk itu pentingnya memberikan pelajaran Bahasa Inggris sejak dini memang dirasa tepat.','http://192.168.100.62:8080/data/course/bahasa_inggris.png');


INSERT INTO course_material (id,course_id,material_index,title) VALUES ('9b807042-2d9f-4961-a95e-af7b3cde2d0e','8c6c787f-1ec1-4be2-b7f0-161e0c3a5345',1,'Bab 1 : Numbers');
INSERT INTO course_material (id,course_id,material_index,title) VALUES ('80cb982e-edad-4d11-a77e-f4f892d20d3d','8c6c787f-1ec1-4be2-b7f0-161e0c3a5345',2,'Bab 2 : Asking Personal Information');
INSERT INTO course_material (id,course_id,material_index,title) VALUES ('6045e993-0242-4881-b395-464cbca226c8','8c6c787f-1ec1-4be2-b7f0-161e0c3a5345',3,'Bab 3 : Reading Comprehension');


INSERT INTO course_material_detail (course_material_id,position,title,type_material,content,image_url) VALUES ('9b807042-2d9f-4961-a95e-af7b3cde2d0e',1,'',1,'','http://192.168.100.62:8080/data/material/bahasa_inggris_bab_1_1.png');

INSERT INTO course_material_detail (course_material_id,position,title,type_material,content,image_url) VALUES ('80cb982e-edad-4d11-a77e-f4f892d20d3d',1,'Read the conversation below!',0,'A: Hello, how are you?\n B: I am fine, thanks.\n A: What is your name?\n B: My name is Sasha.\n A: Where do you come from?\n B: I come from Jakarta.\n A: Where do you live?\n B: I live in Garut\n A: Who do you live with?\n B: I live with my family.\n A: How old are you?\n B: I am fifteen years old.\n A: Where and when were you born?\n B: I was born in Jakarta on March 2nd, 2000\n A: Where do you study?\n B: I study in SMPN Garut 1.\n A: What is your fathers name?\n B: My Fathers name is Tony.\n A: What is your mothers name?\n B: My mothers name is Abby.\n A: How many brothers do you have?\n B: I have one brother\n A: How many sisters do you have?\nB : I dont have any sister.\n A: What is your telephone number?\n B: My telephone number is 253 698,\n A: What is your hobby?\n B: My hobby is studying.\n A: What is your favorite color?\n B: My favorite color is red,\n A: What is your favorite subject?\n B: My favorite subject is Mathematics.\n A: What is your favorite food?\n B: My favorite food is fried rice.\n A: What is your favorite drink?\n B: My favorite drink is orange juice.\n A: What is your favorite fruit?\n B: My favorite fruit is apple.\n A: What is your favorite vegetable?\n B: My favorite vegetable is carrot.','');

INSERT INTO course_material_detail (course_material_id,position,title,type_material,content,image_url) VALUES ('6045e993-0242-4881-b395-464cbca226c8',1,'Read the text below!',0,'On Sunday, Tom gets up at 10 oclock. Then he reads his newspaper in the kitchen. He hasbreakfast at 11.30 and then he telephones his mother in Scotland','');
INSERT INTO course_material_detail (course_material_id,position,title,type_material,content,image_url) VALUES ('6045e993-0242-4881-b395-464cbca226c8',2,'',0,'In the afternoon, at 1.00, Tom plays tennis with his sister and after that, they eat lunch in arestaurant. At 6.00, Tom swims for one hour and then he goes by bike to his brothers house.They talk and listen to music.','');
INSERT INTO course_material_detail (course_material_id,position,title,type_material,content,image_url) VALUES ('6045e993-0242-4881-b395-464cbca226c8',3,'',0,'Tom watches television in the evening and drinks a glass of Jack Daniels whiskey. He goes to bed at 11.30.','');


INSERT INTO course_exam (id,course_id,type_exam,exam_index,text,image_url) VALUES ('0a5531bb-63d5-4ca7-a46f-d3e28cc7f3bc','8c6c787f-1ec1-4be2-b7f0-161e0c3a5345',0,1,'What time does Tom have breakfast on Sunday?','');
INSERT INTO course_exam (id,course_id,type_exam,exam_index,text,image_url) VALUES ('dbf3a8e7-70d2-42db-a0f0-7282d939ced5','8c6c787f-1ec1-4be2-b7f0-161e0c3a5345',0,2,'Who does he telephone in the morning?','');
INSERT INTO course_exam (id,course_id,type_exam,exam_index,text,image_url) VALUES ('4ca46ce3-3322-4c1e-855d-84dc99018317','8c6c787f-1ec1-4be2-b7f0-161e0c3a5345',0,3,'Where does his mother live?','');
INSERT INTO course_exam (id,course_id,type_exam,exam_index,text,image_url) VALUES ('b1394ecd-ee23-4356-af15-1a9e1a3a80cd','8c6c787f-1ec1-4be2-b7f0-161e0c3a5345',0,4,'What time does he play tennis with his sister?','');
INSERT INTO course_exam (id,course_id,type_exam,exam_index,text,image_url) VALUES ('4ce0cf82-8c63-400c-a5dd-188dd50b1e74','8c6c787f-1ec1-4be2-b7f0-161e0c3a5345',0,5,'How long does Tom swim for?','');


INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('26e43c24-c885-44aa-baf3-2352a450e66a','0a5531bb-63d5-4ca7-a46f-d3e28cc7f3bc',0,'A','10.30','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('7f087ad3-9b48-4662-a9c1-d8746a7e2909','0a5531bb-63d5-4ca7-a46f-d3e28cc7f3bc',0,'B','11.00','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('5bd2cb4e-54ac-4cfe-8167-07eee3f1f6b5','0a5531bb-63d5-4ca7-a46f-d3e28cc7f3bc',0,'C','11.30','');

INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('a2916eb3-8700-47ac-9cf3-cab10361ca2e','dbf3a8e7-70d2-42db-a0f0-7282d939ced5',0,'A','His mother','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('2424bfb7-1e0c-4996-aa10-aa5dfeeb0191','dbf3a8e7-70d2-42db-a0f0-7282d939ced5',0,'B','His Plumber','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('acdb6018-b2d7-4bbc-9b13-c4980592f3a7','dbf3a8e7-70d2-42db-a0f0-7282d939ced5',0,'C','His Wife','');

INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('04066439-e73e-4f52-a9e8-bea9d25f7443','4ca46ce3-3322-4c1e-855d-84dc99018317',0,'A','Kuvukiland','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('b57b777a-3a74-451f-808d-5683000b0003','4ca46ce3-3322-4c1e-855d-84dc99018317',0,'B','Scotland','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('6f38ec4c-5bcb-4aea-845a-9e9ed6c945ea','4ca46ce3-3322-4c1e-855d-84dc99018317',0,'C','Australia','');

INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('36ae91ff-c737-4188-8aee-3cb2291d2f66','b1394ecd-ee23-4356-af15-1a9e1a3a80cd',0,'A','afternoon, at 1.00','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('742372cd-dd70-4405-84c8-5a49952706c3','b1394ecd-ee23-4356-af15-1a9e1a3a80cd',0,'B','afternoon, at 1.30','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('1acba7f3-bbe5-42c2-8859-7db349c01a56','b1394ecd-ee23-4356-af15-1a9e1a3a80cd',0,'C','afternoon, at 2.00','');

INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('10fabd0c-6e5b-418b-a4c5-998578996240','4ce0cf82-8c63-400c-a5dd-188dd50b1e74',0,'A','for half hour','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('9a9c45cd-cbf5-413a-9002-1fa42dca0858','4ce0cf82-8c63-400c-a5dd-188dd50b1e74',0,'B','for two hour','');
INSERT INTO course_exam_answer (id,course_exam_id,type_answer,label,text,image_url) VALUES ('c9ceed22-c09b-43b5-ad7b-ed5abc0f34f7','4ce0cf82-8c63-400c-a5dd-188dd50b1e74',0,'C','for one hour','');


INSERT INTO course_exam_solution (course_exam_id,course_exam_answer_id) VALUES ('0a5531bb-63d5-4ca7-a46f-d3e28cc7f3bc','5bd2cb4e-54ac-4cfe-8167-07eee3f1f6b5');
INSERT INTO course_exam_solution (course_exam_id,course_exam_answer_id) VALUES ('dbf3a8e7-70d2-42db-a0f0-7282d939ced5','a2916eb3-8700-47ac-9cf3-cab10361ca2e');
INSERT INTO course_exam_solution (course_exam_id,course_exam_answer_id) VALUES ('4ca46ce3-3322-4c1e-855d-84dc99018317','b57b777a-3a74-451f-808d-5683000b0003');
INSERT INTO course_exam_solution (course_exam_id,course_exam_answer_id) VALUES ('b1394ecd-ee23-4356-af15-1a9e1a3a80cd','36ae91ff-c737-4188-8aee-3cb2291d2f66');
INSERT INTO course_exam_solution (course_exam_id,course_exam_answer_id) VALUES ('4ce0cf82-8c63-400c-a5dd-188dd50b1e74','c9ceed22-c09b-43b5-ad7b-ed5abc0f34f7');





