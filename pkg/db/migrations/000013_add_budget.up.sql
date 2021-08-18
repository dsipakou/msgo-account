CREATE TABLE budget
(
  id          bigserial     PRIMARY KEY,
  title       varchar(40)   NOT NULL,
  amount      numeric(9, 4) NOT NULL,
  budget_date   date          NOT NULL,
  description text
);

ALTER TABLE transactions
  ADD COLUMN budget_id int NULL,
  ADD CONSTRAINT fk_budget
    FOREIGN KEY (budget_id)
    REFERENCES budget(id);
