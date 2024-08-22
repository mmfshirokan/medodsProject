SET TIMEZONE='UTC';
Create table medods.refresh_token (
    id uuid PRIMARY KEY,
    user_id uuid,
    hash CHARACTER VARYING(200) NOT NULL,
    expiration timestamp NOT NULL,
    UNIQUE(hash),
    FOREIGN KEY(user_id) 
        REFERENCES medods.user(id)
);