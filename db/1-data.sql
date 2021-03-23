-- Init users
INSERT INTO "users" (
    "nickname",
    "email",
    "password",
    "created_at",
    "updated_at"
  )
VALUES (
    'Jean',
    'dufour@gmail.com',
    '/RX4Oed0OPWOjHDqkPNkeKoUGnXGF0r1uhaEcX1RJZC',
    '2021-03-22 16:49:03',
    '2021-03-22 16:49:03'
  )
RETURNING "users"."id";
INSERT INTO "users" (
    "nickname",
    "email",
    "password",
    "created_at",
    "updated_at"
  )
VALUES (
    'Myouu',
    'myou@gmail.com',
    '/EuCJTnNy3nUcxI6kzpYxK',
    '2021-03-22 16:49:03',
    '2021-03-22 16:49:03'
  )
RETURNING "users"."id";
