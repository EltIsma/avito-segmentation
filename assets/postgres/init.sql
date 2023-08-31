
CREATE TABLE storage(  
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(100) NOT NULL UNIQUE
);
/* 2023-08-25 16:26:37 [17 ms] */ 
CREATE TABLE users(  
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL 
);


CREATE TABLE UsersSegment(
user_id INTEGER NOT NULL UNIQUE,
CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES users(id),
slug_name text[] NOT NULL 
); 


CREATE TABLE Report(
id SERIAL PRIMARY KEY,
user_id INTEGER NOT NULL,
slug VARCHAR(100) NOT NULL,
operation VARCHAR(100) NOT NULL,
execution TIMESTAMP not null,
ttl TIMESTAMP NOT NULL  DEFAULT '9999-01-01 00:00:00'
); 


INSERT INTO users (email, name) VALUES ('kuku@mail.ru', 'Christopher');
INSERT INTO users (email, name) VALUES ('den@mail.ru', 'Nolan');