USE auth;

CREATE OR REPLACE VIEW view_user AS
SELECT a.id                               AS id,
       a.username                         AS username,
       a.password                         AS password,
       a.name_family                      AS name_family,
       a.name_given                       AS name_given,
       GROUP_CONCAT(s.name SEPARATOR ' ') AS scope,
       a.active                           AS active,
       a.created                          AS created
FROM user_scope
         INNER JOIN user a
         ON user_scope.user_id = a.id
         INNER JOIN scope s
         ON user_scope.scope_id = s.id
GROUP BY a.id;
