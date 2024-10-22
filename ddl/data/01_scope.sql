USE auth;

INSERT INTO scope (id, name, user_default)
VALUES (UUID_TO_BIN(UUID()), 'admin', FALSE),
       (UUID_TO_BIN(UUID()), 'user', TRUE);
