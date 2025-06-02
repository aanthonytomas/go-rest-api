-- Create the database if it doesn't exist
SELECT 'CREATE DATABASE sampledb'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'sampledb')\gexec

-- Connect to the database
\c sampledb

-- Create the items table if it doesn't exist
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Insert sample data
INSERT INTO items (name) VALUES ('Item 1'), ('Item 2'), ('Item 3')
ON CONFLICT (id) DO NOTHING;

-- Grant all privileges on the items table to the sampleuser
GRANT ALL PRIVILEGES ON TABLE items TO sampleuser;
GRANT USAGE, SELECT ON SEQUENCE items_id_seq TO sampleuser;
