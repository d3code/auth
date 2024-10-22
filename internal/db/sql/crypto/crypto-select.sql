SELECT BIN_TO_UUID(id),
       slug,
       symbol,
       name,
       description,
       date_launched,
       website,
       coin_market_cap_id,
       coin_market_cap_name,
       coin_market_cap_active,
       coin_market_cap_rank,
       coin_market_cap_logo_url,
       coin_market_cap_date_added,
       coin_market_cap_date_first_data,
       coin_market_cap_date_last_data,
       infinite_supply,
       self_reported_circulating_supply,
       self_reported_market_cap,
       updated
FROM cryptocurrency
WHERE (symbol LIKE ? OR name LIKE ?)
  AND (coin_market_cap_active = TRUE OR coin_market_cap_active != ?)
ORDER BY -coin_market_cap_rank DESC
LIMIT ? OFFSET ?
