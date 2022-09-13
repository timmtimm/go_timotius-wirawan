CREATE DATABASE outlet_pulsa;

USE outlet_pulsa;

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

CREATE TABLE payment_methods (
    id int AUTO_INCREMENT PRIMARY KEY,
    explanation longtext NOT NULL,
    name varchar(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
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

CREATE TABLE customer (
    id int AUTO_INCREMENT PRIMARY KEY,
	address varchar(255) NOT NULL,
    dob date NOT NULL,
    status varchar(255) NOT NULL,
    created_at timestamp,
    updated_at timestamp
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
    CONSTRAINT FK_customer FOREIGN KEY (customer_id) REFERENCES customer(id),
    CONSTRAINT FK_payment_methods FOREIGN KEY (payment_methods_id) REFERENCES payment_methods(id),
    CONSTRAINT FK_product FOREIGN KEY (product_id) REFERENCES product(id)
);