CREATE TABLE users
(
    id              SERIAL PRIMARY KEY      NOT NULL UNIQUE,
    name            VARCHAR(255)            NOT NULL,
    username        VARCHAR(255)            NOT NULL UNIQUE,
    password_hash   VARCHAR(255)            NOT NULL 
);

CREATE TABLE merchants
(
    id              SERIAL PRIMARY KEY                              NOT NULL UNIQUE,
    user_id         INT REFERENCES users (id) ON DELETE CASCADE     NOT NULL,
    merchant_id     INT                                             NOT NULL UNIQUE,
    description     TEXT ,
);

CREATE TABLE advert_items
(
    id              SERIAL PRIMARY KEY                                          NOT NULL UNIQUE,
    title           VARCHAR(255)                                                NOT NULL,
    description     TEXT ,
    images_path     VARCHAR(500),
    merchant_id     INT REFERENCES merchants (merchant_id) ON DELETE CASCADE    NOT NULL
);

CREATE TABLE merchant_adverts
(
    id              SERIAL PRIMARY KEY                                                  NOT NULL UNIQUE,
    merchant_id     INT REFERENCES  merchants (id) ON DELETE CASCADE                    NOT NULL,
    advert_id       INT REFERENCES advert_items (id) ON DELETE CASCADE                  NOT NULL,
);

