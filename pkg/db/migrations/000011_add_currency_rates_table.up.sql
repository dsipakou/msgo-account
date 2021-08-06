CREATE TABLE rates
(
  id bigserial PRIMARY KEY,
  currency_id integer NOT NULL,
  rate_date date NOT NULL,
  rate numeric(7, 4) NOT NULL,
  description text,
  created_at timestamptz DEFAULT NOW(),
  updated_at timestamptz DEFAULT NOW(),
  CONSTRAINT fk_currency
    FOREIGN KEY (currency_id)
    REFERENCES currencies(id)
);
