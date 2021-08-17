ALTER TABLE transactions
  DROP CONSTRAINT fk_budget,
  DROP COLUMN budget_id;

DROP TABLE budget;
