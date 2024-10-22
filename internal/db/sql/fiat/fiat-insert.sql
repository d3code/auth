INSERT INTO fiat(id,
                 coin_market_cap_id,
                 sign,
                 symbol,
                 name)
VALUES ()
ON DUPLICATE KEY UPDATE name = VALUES(name)