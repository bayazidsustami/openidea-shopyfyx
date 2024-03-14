CREATE TABLE IF NOT EXISTS bank_accounts (
    bank_account_id SERIAL PRIMARY KEY,
    bank_account_name VARCHAR(255) NOT NULL,
    bank_account_number VARCHAR(255) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    user_id SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);