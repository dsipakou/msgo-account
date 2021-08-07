ALTER TABLE currencies
  DROP COLUMN rate,
  ADD COLUMN is_default boolean DEFAULT FALSE;

ALTER TABLE rates
  ADD CONSTRAINT unique_date_rate 
    UNIQUE (currency_id, rate_date);
