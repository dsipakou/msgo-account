ALTER TABLE currencies
  ADD COLUMN rate numeric(7, 4) NULL,
  DROP COLUMN is_default;

ALTER TABLE rates
  DROP CONSTRAINT unique_date_rate;
