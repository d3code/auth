USE auth;

# declare uid UUID_TO_BIN(UUID());
SET @uid_a := UUID_TO_BIN(UUID());

SET @uid_root := UUID_TO_BIN(UUID());

# Create account
INSERT INTO user (id,
                  username,
                  password,
                  name_family,
                  name_given)
VALUES (@uid_a,
        'a',
        '$2a$12$EXVcFG71dUpO8hqix7CCE.JQOamkWUX/tQqaom3tSPsVXHwNZ/KU6',
        'User',
        'Standard');

INSERT INTO user (id,
                  username,
                  password,
                  name_family,
                  name_given)
VALUES (@uid_root,
        'root',
        '$2a$12$EXVcFG71dUpO8hqix7CCE.JQOamkWUX/tQqaom3tSPsVXHwNZ/KU6',
        'User',
        'Root');


INSERT INTO user_scope (user_id,
                        scope_id)
SELECT @uid_a, id
FROM scope
WHERE name = 'auth:admin';