ALTER TABLE budget
  ADD COLUMN category_id int NULL,
  ADD CONSTRAINT fk_category
    FOREIGN KEY (category_id)
    REFERENCES categories(id);
