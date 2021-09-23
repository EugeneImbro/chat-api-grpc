CREATE TABLE users
(
    id       serial       not null unique,
    nickname varchar(255) not null unique
);