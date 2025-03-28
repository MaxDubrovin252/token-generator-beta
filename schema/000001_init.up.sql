CREATE TABLE tokens(
    id serial not null unique,
    token varchar(255) not null unique,
    message varchar(255) not null
);