USE auth;

CREATE TABLE user_access_token
(
    id      BINARY(16)                               NOT NULL PRIMARY KEY,
    user_id BINARY(16)                               NOT NULL,
    token   TEXT                                     NOT NULL,
    created DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    expires DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    active  BOOLEAN     DEFAULT TRUE                 NOT NULL,

    -- foreign key constraint
    CONSTRAINT FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE user_refresh_token
(
    id        BINARY(16)                               NOT NULL PRIMARY KEY,
    user_id   BINARY(16)                               NOT NULL,
    secret_id BINARY(16)                               NOT NULL,
    issuer    VARCHAR(512)                             NOT NULL,
    token     TEXT                                     NOT NULL,
    created   DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    expires   DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    active    BOOLEAN     DEFAULT TRUE                 NOT NULL,

    -- foreign key constraint
    CONSTRAINT FOREIGN KEY (user_id) REFERENCES user (id)
);
