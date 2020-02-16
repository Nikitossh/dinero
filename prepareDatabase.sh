#!/bin/bash

# This script prepare posgresql database for development
# Written by Nikita Shesterikov at 2019-02-11

PORT=5432
POSTGRES_CONFIG="pgconfig.conf"
POSTGRES_PASSWORD="TodayIsTheBestDay"
POSTGRES_USER="dinero"
POSTGRES_DB="dinero"
PGDATA_LOCAL_STORAGE="/var/lib/postgresql/data"
PGDATA="/var/lib/postgresql/data"
CONTAINER_NAME="pg_dinero"

# get the default config if not present
if [[ ! -f $POSTGRES_CONFIG ]]; then
  echo "Trying to get default configuration for postgres"
  docker run -i --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > $POSTGRES_CONFIG
# You can customize config by hand. Below line MUST be set for usage by other containers
# listen_address = '*'
fi

sleep 1

# Run with custom config
echo "Running postgres workload"
docker run -d \
  --name $CONTAINER_NAME \
  -v "$PWD/$POSTGRES_CONFIG":/etc/postgresql/postgresql.conf \
  -v "$PGDATA_LOCAL_STORAGE":/$PGDATA \
  -e POSTGRES_CONFIG=$POSTGRES_CONFIG \
  -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  -e POSTGRES_USER=$POSTGRES_USER \
  -e POSTGRES_DB=$POSTGRES_DB \
  -p $PORT:5432 \
  postgres 
