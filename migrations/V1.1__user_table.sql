CREATE TABLE medods.user { 
    id uuid PRIMARY KEY,
    emale CHARACTER VARYING(100) UNIQUE NOT NULL,
    ip inet NOT NULL
};