USE auth;

INSERT INTO scope (id,
                   name,
                   user_default)
VALUES (UUID_TO_BIN(UUID()),
        'auth:admin',
        FALSE),
       (UUID_TO_BIN(UUID()),
        'auth:user',
        TRUE);
