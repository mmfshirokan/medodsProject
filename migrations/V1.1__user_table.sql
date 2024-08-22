CREATE TABLE medods.user ( 
    id uuid PRIMARY KEY,
    ip inet NOT NULL,
    name CHARACTER VARYING(100) NOT NULL,
    email CHARACTER VARYING(100) NOT NULL,
    password CHARACTER VARYING(100) NOT NULL,
    UNIQUE(ip, name, email)
);