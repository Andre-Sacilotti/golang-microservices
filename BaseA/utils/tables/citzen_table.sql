CREATE TABLE citizens (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255),
    cpf         VARCHAR(255)     UNIQUE,
    birthdate   TIMESTAMP
 );

