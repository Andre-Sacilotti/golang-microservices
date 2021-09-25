CREATE TABLE address(
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    postal_code         VARCHAR(8),
    address             VARCHAR(255),
    number              VARCHAR(30),
    complement          VARCHAR(255),
    neighbourhood       VARCHAR(255),
    city                VARCHAR(255),
    state               VARCHAR(255),
    country             VARCHAR(255),
    citzen_id           BIGSERIAL,
    FOREIGN KEY (citzen_id) REFERENCES natural_person(id)
);

