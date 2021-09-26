DROP TABLE IF EXISTS financial_operations;

CREATE TABLE financial_operations (
  id                    BIGINT(50) UNSIGNED NOT NULL AUTO_INCREMENT,
  cpf                   VARCHAR(11) NOT NULL,
  operation_type        ENUM("credit_card_purchase", "debit_card_purchase", "credit_bureau_consultation") NOT NULL,
  realized_at           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  operation_value       FLOAT,
  modification_time     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  insertion_time        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  PRIMARY KEY (id),
  UNIQUE KEY unique_id (id),
);
INSERT INTO financial_operations (
    cpf, operation_type, operation_value
    ) VALUES (
    "12312312390", "credit_card_purchase", 1900.97
    )