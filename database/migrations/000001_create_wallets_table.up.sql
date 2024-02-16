CREATE TABLE wallets (
    incremental_id INT AUTO_INCREMENT,
    id VARCHAR(32) UNIQUE,
    owner_id VARCHAR(255) NOT NULL,
    balance INT NOT NULL DEFAULT 0,
    currency VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (incremental_id)
);
