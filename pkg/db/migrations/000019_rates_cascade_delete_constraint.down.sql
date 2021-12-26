ALTER TABLE rates
  DROP CONSTRAINT fk_currency,
  ADD CONSTRAINT fk_currency
    FOREIGN KEY (currency_id)
    REFERENCES currencies(id);
