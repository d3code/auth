USE auth;

CREATE TABLE scope
(
    id           BINARY(16)            NOT NULL PRIMARY KEY,
    name         VARCHAR(255)          NOT NULL,
    user_default BOOLEAN DEFAULT FALSE NOT NULL,

    CONSTRAINT UNIQUE (name)
);

CREATE TABLE user_scope
(
    id       BINARY(16) DEFAULT (UUID_TO_BIN(UUID())) NOT NULL PRIMARY KEY,
    user_id  BINARY(16)                               NOT NULL,
    scope_id BINARY(16)                               NOT NULL,

    -- duplicate scope on account constraint
    CONSTRAINT UNIQUE (user_id, scope_id),

    -- foreign key constraints
    CONSTRAINT FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (scope_id) REFERENCES scope (id) ON DELETE CASCADE
);

