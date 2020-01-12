CREATE TABLE users(
    id bigserial not null primary key,
    email varchar(50) not null unique,
    encrypted_password varchar (128) not null
)