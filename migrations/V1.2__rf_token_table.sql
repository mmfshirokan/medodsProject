Create table medods.refresh_token {
    id uuid PRIMARY KEY,
    user_id uuid FOREIGN KEY REFERENCES medods.user(id),
    hash CHARACTER VARYING(200) NOT NULL,
    expiration timestamp NOT NULL
}