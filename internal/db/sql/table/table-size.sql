SELECT table_schema                                            AS 'database',
       ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS size_in_mb
FROM information_schema.tables
GROUP BY table_schema