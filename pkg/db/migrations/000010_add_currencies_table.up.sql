CREATE TABLE currencies
(
  id          bigserial     PRIMARY KEY,
  code        varchar(5)    NOT NULL,
  sign        varchar(2)    NOT NULL,
  verbal_name varchar(20)   NOT NULL,
  rate        numeric(7, 4) NOT NULL,
  comments    text,
  created_at  timestamptz   DEFAULT NOW(),
  updated_at  timestamptz   DEFAULT NOW()
);

ALTER TABLE transactions
  ADD COLUMN currency_id int NULL,
  ADD CONSTRAINT fk_currency
    FOREIGN KEY (currency_id)
    REFERENCES currencies(id);
