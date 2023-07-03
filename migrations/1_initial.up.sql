BEGIN;

-- CREATE USER my_user WITH PASSWORD 'my_password';
-- GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO my_user;


CREATE TABLE IF NOT EXISTS orders (
    order_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255) NOT NULL,
    entry VARCHAR(255) NOT NULL,
    locale VARCHAR(255) NOT NULL,
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255) NOT NULL,
    delivery_service VARCHAR(255) NOT NULL,
    shardkey VARCHAR(255) NOT NULL,
    sm_id INTEGER NOT NULL,
    date_created TIMESTAMP NOT NULL,
    oof_shard VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS delivery (
    order_id VARCHAR(255) UNIQUE REFERENCES orders(order_uid),
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    zip VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL 
);

CREATE TABLE IF NOT EXISTS payment (
    order_id VARCHAR(255) REFERENCES orders(order_uid),
    transaction VARCHAR(255) NOT NULL,
    request_id VARCHAR(255) NOT NULL,
    currency VARCHAR(255) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL,
    payment_dt INT NOT NULL,
    bank VARCHAR(255) NOT NULL,
    delivery_cost FLOAT NOT NULL,
    goods_total INTEGER NOT NULL,
    custom_fee FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS order_items (
    order_id VARCHAR(255) REFERENCES orders(order_uid),
    chrt_id INTEGER NOT NULL,
    item_track_number VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    rid VARCHAR(255) NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    sale INTEGER NOT NULL,
    size VARCHAR(255) NOT NULL,
    total_price INTEGER NOT NULL,
    nm_id INTEGER NOT NULL,
    brand VARCHAR(255) NOT NULL,
    status INTEGER NOT NULL
);

COMMIT;