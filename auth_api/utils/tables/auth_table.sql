CREATE TABLE auths (
    id          BIGSERIAL       NOT NULL PRIMARY KEY,
    Username        VARCHAR(255)    UNIQUE,
    Password    VARCHAR(255)
);

 INSERT INTO auths (Username, Password) VALUES ('andre', '1234')

