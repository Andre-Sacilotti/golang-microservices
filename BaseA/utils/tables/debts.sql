CREATE TABLE debts (
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    debtor_id                   INTEGER     NOT NULL,
    value                       INTEGER     NOT NULL,
    was_negociated              BOOL        NOT NULL DEFAULT FALSE,
    credit_taken_at             TIMESTAMP,
    credit_turned_debit_at      TIMESTAMP,
    FOREIGN KEY (debtor_id) REFERENCES natural_person(id)
 );

