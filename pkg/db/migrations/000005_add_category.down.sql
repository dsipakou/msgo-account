ALTER TABLE transactions
  DROP CONSTRAINT fk_category,
  ALTER COLUMN category_id TYPE varchar(255);

ALTER TABLE transactions
  RENAME COLUMN category_id TO category;

DROP TABLE categories;
