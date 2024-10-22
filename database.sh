#!/bin/bash

export MYSQL_USER="root"
export MYSQL_PWD="password"
export MYSQL_HOST="127.0.0.1"
export MYSQL_PORT="3306"

# Create database
mysql mysql --host="$MYSQL_HOST" --user="$MYSQL_USER" --port="$MYSQL_PORT" <"ddl/schema/create.sql"

# Create tables
for file in ddl/table/*.sql; do
  echo "-- $file"
  mysql --host="$MYSQL_HOST" --user="$MYSQL_USER" --port="$MYSQL_PORT" <"$file"
done

# Create views
for file in ddl/view/*.sql; do
  echo "-- $file"
  mysql --host="$MYSQL_HOST" --user="$MYSQL_USER" --port="$MYSQL_PORT" <"$file"
done

# Create triggers
for file in ddl/trigger/*.sql; do
  echo "-- $file"
  mysql --host="$MYSQL_HOST" --user="$MYSQL_USER" --port="$MYSQL_PORT" <"$file"
done

# Create root data
for file in ddl/data/*.sql; do
  echo "-- $file"
  mysql --host="$MYSQL_HOST" --user="$MYSQL_USER" --port="$MYSQL_PORT" <"$file"
done

