CREATE TABLE transactions
(
  id        bigserial     PRIMARY KEY,
  user_id   int           NOT NULL,
  category  varchar(255)  NOT NULL,
  amount    int           NOT NULL
);
