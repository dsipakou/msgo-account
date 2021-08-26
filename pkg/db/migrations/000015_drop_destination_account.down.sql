ALTER TABLE transactions
  ALTER COLUMN category_id SET NOT NULL,
  DROP CONSTRAINT account_or_category,
  DROP COLUMN dest_account_id;
