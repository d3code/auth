SELECT COUNT(*)
FROM fiat
WHERE (symbol LIKE ? OR name LIKE ?)
