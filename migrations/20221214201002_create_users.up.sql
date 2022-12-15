create table information(
    id bigserial not null primary key,
    email varchar not null unique,
    encryptedPassword varchar not null
);