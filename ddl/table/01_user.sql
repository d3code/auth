USE auth;

CREATE TABLE user
(
    id          BINARY(16)                               NOT NULL PRIMARY KEY,
    username    VARCHAR(512)                             NOT NULL,
    password    VARCHAR(2048)                            NOT NULL,
    name_given  VARCHAR(2048)                            NOT NULL,
    name_family VARCHAR(2048)                            NOT NULL,
    active      BOOLEAN     DEFAULT TRUE                 NOT NULL,
    created     DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,

    -- unique constraint on username
    CONSTRAINT UNIQUE (username)
);

