ALTER TABLE transactions
  ADD COLUMN transaction_date date NOT NULL DEFAULT CURRENT_DATE;
