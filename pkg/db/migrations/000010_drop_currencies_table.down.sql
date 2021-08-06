ALTER TABLE transactions
  DROP CONSTRAINT fk_currency,
  DROP COLUMN currency_id;
DROP TABLE currencies;
