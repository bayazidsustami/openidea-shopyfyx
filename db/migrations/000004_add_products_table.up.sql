CREATE TYPE condition AS ENUM ('new', 'second');

CREATE TABLE IF NOT EXISTS products (
	product_id   SERIAL PRIMARY KEY,
	product_name VARCHAR(255) NOT NULL,
	condition    condition,
	tags         VARCHAR(255) NOT NULL,
	is_available BOOLEAN NOT NULL,
	image_url    VARCHAR(255) NOT NULL,
	user_id      INTEGER NOT NULL,
	deleted_at   TIMESTAMP,
	created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_stocks (
    product_stock_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE products ADD FOREIGN KEY (user_id) REFERENCES users (user_id);

ALTER TABLE product_stocks ADD FOREIGN KEY (product_id) REFERENCES products (product_id);