CREATE DATABASE outlet_pulsa;

USE outlet_pulsa;

CREATE TABLE product_types ( 
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE operators (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment_methods (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_descriptions (
	id INT(11) AUTO_INCREMENT PRIMARY KEY,
	description LONGTEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
 	id INT(11) AUTO_INCREMENT PRIMARY KEY,
    product_type_id INT NOT NULL,
    operator_id INT NOT NULL,
    product_description_id INT NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_product_type_products FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    CONSTRAINT FK_operator_products FOREIGN KEY (operator_id) REFERENCES operators(id),
    CONSTRAINT FK_product_description_products FOREIGN KEY (product_description_id) REFERENCES product_descriptions(id)
);

CREATE TABLE users (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    dob date NOT NULL,
    gender CHAR(1) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
 	id INT(11) AUTO_INCREMENT PRIMARY KEY,
    user_id INT(11) NOT NULL,
    payment_method_id INT(11) NOT NULL,
    status VARCHAR(10) NOT NULL,
    total_qty INT(11) NOT NULL,
    total_price NUMERIC(25, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_user_transaction FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT FK_payment_method_transaction FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    transaction_id INT NOT NULL,
    status VARCHAR(10) NOT NULL,
    qty INT(11) NOT NULL,
    price NUMERIC(25, 2) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_product_transaction_detail FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT FK_transaction_transaction_detail FOREIGN KEY (transaction_id) REFERENCES transactions(id)
);