#!/bin/bash

echo "🛠️  Running DB migrations..."

psql $POSTGRES_CONN <<EOF
CREATE TABLE IF NOT EXISTS routes (
    id SERIAL PRIMARY KEY,
    origin TEXT,
    destination TEXT,
    total_time FLOAT,
    created_at TIMESTAMP DEFAULT NOW()
);
EOF
