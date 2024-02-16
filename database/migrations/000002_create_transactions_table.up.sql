CREATE TABLE transactions (
    incremental_id INT AUTO_INCREMENT,
    id VARCHAR(32) UNIQUE,
    wallet_id VARCHAR(255) NOT NULL,
    amount INT NOT NULL DEFAULT 0,
    type VARCHAR(255) NOT NULL,
    origin_id VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (incremental_id)
);
