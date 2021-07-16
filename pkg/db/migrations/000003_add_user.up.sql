CREATE TABLE users
(
  id          bigserial     PRIMARY KEY,
  name        varchar(128)  NOT NULL,
  email       varchar(128)  NOT NULL,
  password    varchar(128)  NOT NULL,
  created_at  timestamptz   DEFAULT NOW(),
  updated_at  timestamptz   DEFAULT NOW()
);

ALTER TABLE transactions
  ADD CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE;

ALTER TABLE accounts
  ADD CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE;
