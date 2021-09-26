CREATE EXTENSION pgcrypto;

CREATE TABLE auths (
    id          BIGSERIAL       NOT NULL PRIMARY KEY,
    Username        VARCHAR(255)    UNIQUE,
    Password    VARCHAR(255)
);

alter extension pgcrypto set schema public;

INSERT INTO auths (Username, Password) VALUES ('andre', crypt('1234', gen_salt('bf', 15)))

