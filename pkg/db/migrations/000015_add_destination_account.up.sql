ALTER TABLE transactions
  ALTER COLUMN category_id DROP NOT NULL,
  ADD COLUMN dest_account_id int NULL,
  ADD CONSTRAINT account_or_category
    CHECK (category_id IS NOT NULL or dest_account_id IS NOT NULL);
