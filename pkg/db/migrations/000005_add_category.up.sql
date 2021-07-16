CREATE TABLE categories
(
  id          bigserial           PRIMARY KEY,
  name        varchar(128)        NOT NULL,
  parent      varchar(128)        DEFAULT NULL,
  created_at  timestamptz         DEFAULT NOW(),
  updated_at  timestamptz         DEFAULT NOW(),
  CONSTRAINT  unique_name_parent  UNIQUE (name, parent)
);

ALTER TABLE transactions
  RENAME COLUMN category TO category_id;

ALTER TABLE transactions
  ALTER COLUMN category_id TYPE int USING category_id::integer,
  ADD CONSTRAINT fk_category
    FOREIGN KEY (category_id)
    REFERENCES categories(id)
    ON DELETE CASCADE;
