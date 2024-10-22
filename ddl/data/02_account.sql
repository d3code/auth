USE auth;

# Declare id variables
SET @uid_a := UUID_TO_BIN(UUID());
SET @uid_root := UUID_TO_BIN(UUID());

# Create account a
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

# Create account root
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

# Assign scopes to account root
INSERT INTO user_scope (user_id, scope_id)
SELECT @uid_root, id
FROM scope
WHERE name = 'admin';
