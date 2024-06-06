USE coursesplatform;

CREATE SCHEMA coursesplatform;

SHOW TABLES IN coursesplatform;

DESCRIBE TABLE users;
SELECT * FROM users;

DESCRIBE TABLE courses;
SELECT * FROM courses;

DESCRIBE TABLE subscriptions;
SELECT * FROM subscriptions;

INSERT INTO users (id, email, password_hash, type, creation_date, last_updated) VALUES
(1, 'emikohmann@gmail.com', '5f4dcc3b5aa765d61d8327deb882cf99', 'admin', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(2, 'juanperez@gmail.com', '5f4dcc3b5aa765d61d8327deb882cf99', 'normal', '2024-05-30 00:00:00', '2024-05-30 00:00:00');

TRUNCATE TABLE courses;
INSERT INTO courses (id, title, description, category, image_url, creation_date, last_updated) VALUES
(1, 'Python for Data Science, AI & Development', 'What you''ll learn Learn Python - the most popular programming language and for Data Science and Software Development.  Apply Python programming logic Variables, Data Structures, Branching, Loops, Functions, Objects & Classes.  Demonstrate proficiency in using Python libraries such as Pandas & Numpy, and developing code using Jupyter Notebooks.', 'Programming', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://s3.amazonaws.com/coursera-course-photos/fc/c1b8dfbac740999b6256aca490de43/Python-Image.jpg?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&fit=crop&q=50', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(2, 'Computer Science: Programming with a Purpose', 'The basis for education in the last millennium was “reading, writing, and arithmetic;” now it is reading, writing, and computing. Learning to program is an essential part of the education of every student, not just in the sciences and engineering, but in the arts, social sciences, and humanities, as well.', 'Programming', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://s3.amazonaws.com/coursera-course-photos/86/9297002bbc11e8b82d4d350d2ce323/IntroCSlogo.png?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&fit=crop&q=50', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(3, 'Java Programming and Software Engineering Fundamentals Specialization', 'Advance your subject-matter expertise. Learn in-demand skills from university and industry experts. Master a subject or tool with hands-on projects. Develop a deep understanding of key concepts. Earn a career certificate from Duke University.', 'Programming', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://coursera_assets.s3.amazonaws.com/collections/product-images/ibm-full-stack-cloud-developer.jpeg?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(4, 'Introduction to Java Programming', 'Learn the basics of Java programming language including concepts such as variables, data types, control structures, arrays, and methods. Gain hands-on experience through coding exercises and projects.', 'Programming', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://d2j5ihb19pt1hq.cloudfront.net/sdp_page/s12n_logos/Duke_JavaProgrammingIntrotoSoftware_CM195522.jpg?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(5, 'Web Development Fundamentals', 'Learn the fundamentals of web development including HTML, CSS, and JavaScript. Build simple web pages and interactive web applications.', 'Web Development', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://s3.amazonaws.com/coursera-course-photos/3f/63c230ac6b11e78079b5cde52a7a37/C1_square.jpg?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(6, 'Machine Learning Foundations', 'Get started with machine learning by learning foundational concepts and algorithms. Understand supervised and unsupervised learning, regression, classification, clustering, and more.', 'Machine Learning', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://s3.amazonaws.com/coursera-course-photos/83/e258e0532611e5a5072321239ff4d4/jhep-coursera-course4.png?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(7, 'Python for Data Analysis', 'Learn to analyze data using Python. Cover topics such as data manipulation, cleaning, visualization, and statistical analysis using libraries like Pandas, NumPy, and Matplotlib.', 'Data Science', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://d15cw65ipctsrr.cloudfront.net/be/0ce870e9cb11e8b2d233b4be6a9cf3/pythonfluency_1x1_4.png?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(8, 'Cybersecurity Fundamentals', 'Learn the basics of cybersecurity including network security, cryptography, secure coding practices, and threat detection. Understand common cyber threats and how to mitigate them.', 'Cybersecurity', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://d15cw65ipctsrr.cloudfront.net/f7/fb2aa0d85611e8a743d71c4407f947/89066_iconImage_EducationCoursera_aqua.jpg?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(9, 'Introduction to Game Development', 'Learn the basics of game development using popular game engines such as Unity or Unreal Engine. Understand game design principles, scripting, and asset creation.', 'Game Development', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://d15cw65ipctsrr.cloudfront.net/41/4d3d7c05fb42729c9d90352e072ca3/1060x596_GCC-photos_Karrim.png?auto=format%2Ccompress%2C%20enhance&dpr=2&w=320&h=180&q=50&fit=crop&crop=faces', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(10, 'Front-end Web Development Professional Certificate', 'Become proficient in front-end web development by learning HTML, CSS, JavaScript, and popular frameworks like React and Angular. Build responsive and interactive web applications.', 'Web Development', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://d15cw65ipctsrr.cloudfront.net/ed/d25c0d25114924a34754928dbf8273/Front-end-dev-ProCert.png?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(11, 'Back-end Web Development Professional Certificate', 'Master back-end web development by learning server-side programming languages such as Python, Java, or Node.js. Explore database management, API development, and deployment strategies.', 'Web Development', 'https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://d15cw65ipctsrr.cloudfront.net/1e/d2f5a0dc494722980faf74d9fa8fc8/Back-end-dev-ProCert.jpg?auto=format%2Ccompress%2C%20enhance&dpr=2&w=265&h=216&q=50&fit=crop', '2024-05-30 00:00:00', '2024-05-30 00:00:00');

UPDATE courses
SET title = 'Advanced Python for Data Science, AI & Development - Updated',
    last_updated = '2024-05-30 12:00:00'  -- Example of updated timestamp
WHERE id = 1;
