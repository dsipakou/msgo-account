CREATE TABLE accounts 
(
  id          bigserial     PRIMARY KEY,
  user_id     int           NOT NULL, 
  source      varchar(30)   NOT NULL, 
  amount      numeric(9, 4) NOT NULL,  
  description text,
  created_at  timestamptz   DEFAULT NOW(), 
  updated_at  timestamptz   DEFAULT NOW()
);

ALTER TABLE transactions
  ALTER COLUMN amount     TYPE numeric(9, 4),
  ADD COLUMN account_id   int NOT NULL,
  ADD COLUMN description  text,
  ADD COLUMN created_at   timestamptz DEFAULT NOW(),
  ADD COLUMN updated_at   timestamptz DEFAULT NOW(),
  ADD CONSTRAINT fk_account
    FOREIGN KEY (account_id)
    REFERENCES accounts(id)
    ON DELETE CASCADE;
