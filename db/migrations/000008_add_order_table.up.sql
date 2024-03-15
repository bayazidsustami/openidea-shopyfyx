CREATE TABLE IF NOT EXISTS orders (
    order_id                VARCHAR(50) PRIMARY KEY,
    product_id              INT NOT NULL,
    quantity                INT NOT NULL,
    bank_account_id         INT NOT NULL,
    payment_proof_image_url VARCHAR(255) NOT NULL,
    created_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE,
    FOREIGN KEY (bank_account_id) REFERENCES bank_accounts(bank_account_id)
);

CREATE INDEX idx_order_id ON orders (order_id);