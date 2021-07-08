ALTER TABLE transactions
  DROP CONSTRAINT fk_account,
  ALTER COLUMN amount TYPE int,
  DROP COLUMN account_id,
  DROP COLUMN description,
  DROP COLUMN created_at,
  DROP COLUMN updated_at;

DROP TABLE accounts;
