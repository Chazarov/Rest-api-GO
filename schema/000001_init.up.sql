CREATE TABLE users
(
    id              SERIAL PRIMARY KEY      NOT NULL,
    name            VARCHAR(255)            NOT NULL,
    username        VARCHAR(255)            NOT NULL UNIQUE,
    password_hash   VARCHAR(255)            NOT NULL 
);

CREATE TABLE advert_items
(
    id              SERIAL PRIMARY KEY                                          NOT NULL,
    title           VARCHAR(255)                                                NOT NULL,
    description     TEXT,
    images_path     VARCHAR(500),
    user_id         INT REFERENCES users (id) ON DELETE CASCADE                 NOT NULL
);

CREATE TABLE advert_lists
(
    id              SERIAL PRIMARY KEY                                                  NOT NULL,
    user_id         INT REFERENCES users (id) ON DELETE CASCADE                         NOT NULL,
    advert_id       INT REFERENCES advert_items (id) ON DELETE CASCADE                  NOT NULL
);