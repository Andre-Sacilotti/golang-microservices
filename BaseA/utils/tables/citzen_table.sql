CREATE TABLE natural_person (
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    name        VARCHAR(255),
    cpf         VARCHAR(11)     UNIQUE,
    birthdate   TIMESTAMP
 );

