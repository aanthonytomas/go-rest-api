#!/bin/bash

# Create the user if it doesn't exist
psql -U postgres -tc "SELECT 1 FROM pg_user WHERE usename = 'sampleuser'" | grep -q 1 || \
    psql -U postgres -c "CREATE USER sampleuser WITH PASSWORD 'samplepass';"

# Create the database if it doesn't exist
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'sampledb'" | grep -q 1 || \
    psql -U postgres -c "CREATE DATABASE sampledb;"

# Run the initialization script
psql -U postgres -d sampledb -f db/init.sql

echo "Database setup complete!"
