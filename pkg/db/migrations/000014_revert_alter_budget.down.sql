ALTER TABLE budget
  DROP COLUMN is_completed,
  ALTER COLUMN budget_date SET NOT NULL;
