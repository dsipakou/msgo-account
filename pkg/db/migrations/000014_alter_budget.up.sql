ALTER TABLE budget
  ADD COLUMN is_completed boolean DEFAULT false, 
  ALTER COLUMN budget_date DROP NOT NULL;
