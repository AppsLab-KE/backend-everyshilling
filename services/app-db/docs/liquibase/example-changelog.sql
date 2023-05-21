--liquibase formatted sql

--changeset your.name:1 labels:example-label context:example-context
--comment: example comment
create table "user" (
                      id int primary key auto_increment not null,
                      username varchar(50) not null,
                      email varchar(100) not null,
                      password varchar(100) not null
--                        balance DECIMAL(18, 2) NOT NULL

);
--rollback DROP TABLE user;
--changeset your.name:4 labels:example-label context:example-context
--comment: example comment
create table phone (
                       id int primary key auto_increment not null,
                       number varchar(20) not null
);
--rollback DROP TABLE phone;

--changeset your.name:5 labels:example-label context:example-context
--comment: example comment
create table "reset_request" (
                               id int primary key auto_increment not null,
                               tracking_uuid varchar(36) not null,
                               user_id int,
                               otp_code varchar(6),
                               foreign key (user_id) references "user"(id)
);
--rollback DROP TABLE reset_request;

--changeset your.name:6 labels:example-label context:example-context
--comment: example comment
create table "verification_request" (
                                      id int primary key auto_increment not null,
                                      tracking_uuid varchar(36) not null,
                                      user_id int,
                                      otp_code varchar(6),
                                      foreign key (user_id) references "user"(id)
);
--rollback DROP TABLE verification_request;


--changeset your.name:8 labels:example-label context:example-context
--comment: example comment
create table "refresh_token_request" (
                                       id int primary key auto_increment not null,
                                       user_id int,
                                       refresh_token varchar(100) not null,
                                       foreign key (user_id) references "user"(id)
);
--rollback DROP TABLE refresh_token_request;

CREATE TABLE currency (
                          id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
                          code VARCHAR(3) NOT NULL,
                          name VARCHAR(100) NOT NULL
);
--rollback DROP TABLE currency;

--EXCHANGE RATE TABLE
CREATE TABLE exchange_rate (
                               id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
                               source_currency_id INT,
                               target_currency_id INT,
                               rate DECIMAL(18, 4) NOT NULL,
                               FOREIGN KEY (source_currency_id) REFERENCES currency(id),
                               FOREIGN KEY (target_currency_id) REFERENCES currency(id)
);
--rollback DROP TABLE exchange_rate;


-- Transaction table
CREATE TABLE transaction (
                             id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
                             user_id INT,
                             amount DECIMAL(18, 2) NOT NULL,
                             source_currency_id INT,
                             target_currency_id INT,
                             delivery_currency_id INT,
                             total_amount DECIMAL(18, 2) NOT NULL,
                             timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             FOREIGN KEY (user_id) REFERENCES "user"(id),
                             FOREIGN KEY (source_currency_id) REFERENCES currency(id),
                             FOREIGN KEY (target_currency_id) REFERENCES currency(id),
                             FOREIGN KEY (delivery_currency_id) REFERENCES currency(id)
);

--rollback DROP TABLE transaction;

--changeset your.name:7 labels:example-label context:example-context
--comment: example comment
create table "logout_request" (
                                  id int primary key auto_increment not null,
                                  user_id int,
                                  token varchar(100) not null,
                                  foreign key (user_id) references "user"(id)
);
--rollback DROP TABLE logout_request;