USE auth;

DELIMITER $$
CREATE TRIGGER user_scope_default
    AFTER INSERT
    ON user
    FOR EACH ROW
BEGIN
    INSERT INTO user_scope (user_id, scope_id)
    SELECT new.id, scope.id
    FROM scope
    WHERE user_default = TRUE;
END$$
DELIMITER ;

