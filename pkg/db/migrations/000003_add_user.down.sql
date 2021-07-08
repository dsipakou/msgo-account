ALTER TABLE transactions
  DROP CONSTRAINT fk_user;

ALTER TABLE accounts
  DROP CONSTRAINT fk_user;

DROP TABLE users;
