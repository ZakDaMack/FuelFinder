-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS store WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create a table
CREATE TABLE IF NOT EXISTS store.stationdata (
    id uuid PRIMARY KEY,
    site_id text,
    company text,
    e5 double,
    e10 double,
    b7 double,
    address text,
    postcode text,
    latitude double,
    longitude double,
    created_at timestamp
);