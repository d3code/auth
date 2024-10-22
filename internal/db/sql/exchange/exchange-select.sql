SELECT BIN_TO_UUID(id),
       coin_market_cap_id,
       slug,
       name,
       active,
       listed,
       weekly_visits,
       weekly_volume_spot_usd,
       proof_of_reserves_status,
       proof_of_reserves_audit_status,
       date_first_data,
       date_last_data
FROM exchange
WHERE (name LIKE ?) AND (active = TRUE OR active != ?)
ORDER BY weekly_visits DESC
LIMIT ? OFFSET ?;
