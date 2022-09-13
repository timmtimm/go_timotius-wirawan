CREATE DATABASE alta_online_shop;

USE alta_online_shop;

CREATE TABLE product_type ( 
    id int AUTO_INCREMENT PRIMARY KEY,
    type varchar(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE operator (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL,
    type varchar(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE payment_methods_explanation (
    id int AUTO_INCREMENT PRIMARY KEY,
    explanation longtext NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE payment_methods (
    id int AUTO_INCREMENT PRIMARY KEY,
    payment_methods_explanation_id int NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT FK_payment_methods_explanation FOREIGN KEY (payment_methods_explanation_id) REFERENCES payment_methods_explanation(id)
);

CREATE TABLE product (
 	id int AUTO_INCREMENT PRIMARY KEY,
    product_type_id int NOT NULL,
    operator_id int NOT NULL,
    name varchar(255) NOT NULL,
    description longtext NOT NULL,
    price int NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT FK_product_type FOREIGN KEY (product_type_id) REFERENCES product_type(id),
    CONSTRAINT FK_operator FOREIGN KEY (operator_id) REFERENCES operator(id)
);

CREATE TABLE address (
    id int AUTO_INCREMENT PRIMARY KEY,
    street varchar(255) NOT NULL,
    districts varchar(255) NOT NULL,
    province varchar(255) NOT NULL,
    zip_code varchar(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE customer (
    id int AUTO_INCREMENT PRIMARY KEY,
	address_id int NOT NULL,
    dob date NOT NULL,
    status varchar(255) NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT FK_address FOREIGN KEY (address_id) REFERENCES address(id)
);

CREATE TABLE transaction (
 	id int AUTO_INCREMENT PRIMARY KEY,
    customer_id int NOT NULL UNIQUE,
    payment_methods_id int NOT NULL UNIQUE,
    product_id int NOT NULL,
    total int NOT NULL,
    paid boolean NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT FK_customer_transaction FOREIGN KEY (customer_id) REFERENCES customer(id),
    CONSTRAINT FK_payment_methods_transaction FOREIGN KEY (payment_methods_id) REFERENCES payment_methods(id),
    CONSTRAINT FK_product FOREIGN KEY (product_id) REFERENCES product(id)
);

CREATE TABLE user_payment_method_detail (
    id int AUTO_INCREMENT PRIMARY KEY,
    customer_id int NOT NULL,
    payment_methods_id int NOT NULL,
    detail longtext NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT FK_customer_user_payment_method_detail FOREIGN KEY (customer_id) REFERENCES customer(id),
    CONSTRAINT FK_payment_methods_user_payment_method_detail FOREIGN KEY (payment_methods_id) REFERENCES payment_methods(id)
);

CREATE TABLE courier(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    created_at timestamp,
    updated_at timestamp
);

ALTER TABLE courier ADD ongkos_dasar int;

ALTER TABLE courier RENAME TO shipping;

DROP TABLE shipping;