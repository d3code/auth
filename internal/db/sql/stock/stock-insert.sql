INSERT INTO stock(id,
                  market_id,
                  symbol,
                  name_listed,
                  name_display,
                  address,
                  phone,
                  industry,
                  sector,
                  region,
                  description,
                  market_cap,
                  url,
                  ceo)
VALUES ()
ON DUPLICATE KEY UPDATE name_listed = VALUES(name_listed),
                        address = VALUES(address),
                        phone = VALUES(phone),
                        industry = VALUES(industry),
                        sector = VALUES(sector),
                        region = VALUES(region),
                        description = VALUES(description),
                        market_cap = VALUES(market_cap),
                        url = VALUES(url),
                        ceo = VALUES(ceo);