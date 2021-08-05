CREATE TYPE transaction_type AS ENUM ('income', 'outcome');

ALTER TABLE transactions
  ADD COLUMN type transaction_type DEFAULT 'outcome';
