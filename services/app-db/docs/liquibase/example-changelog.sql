--liquibase formatted sql

--changeset your.name:1 labels:example-label context:example-context
--comment: example comment
create table user (
    id int primary key auto_increment not null,
    username varchar(50) not null,
    password varchar(50) not null,
    email varchar(50) not null,
    created_at timestamp default current_timestamp not null
);
--rollback DROP TABLE user;
--changeset your.name:2 labels:example-label context:example-context
--comment: example comment
create table otp (
    id int primary key auto_increment not null,
    value varchar(6) not null,
    user_id int not null ,
    created_at timestamp default current_timestamp not null

);
--rollback DROP TABLE otp;
--changeset other.dev:3 labels:example-label context:example-context
--comment: example comment
alter table user add column username varchar(50)
--rollback ALTER TABLE person DROP COLUMN ;

