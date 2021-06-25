CREATE TABLE transactions
(
  id        serial        not null unique,
  user_id   int           not null,
  category  varchar(255)  not null,
  amount    int           not null
)
