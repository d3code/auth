SELECT BIN_TO_UUID(id), coin_market_cap_id, sign, symbol, name
FROM fiat
WHERE (symbol LIKE ? OR name LIKE ?)
ORDER BY coin_market_cap_id
LIMIT ? OFFSET ?
