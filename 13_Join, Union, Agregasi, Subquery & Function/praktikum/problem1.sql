-- 1.a
INSERT INTO operators (name) VALUES ('telkomsel');
INSERT INTO operators (name) VALUES ('indosat ooredo');
INSERT INTO operators (name) VALUES ('xl');
INSERT INTO operators (name) VALUES ('by.u');
INSERT INTO operators (name) VALUES ('smartfren');

-- 1.b
INSERT product_types (type) VALUES ('prabayar');
INSERT product_types (type) VALUES ('pascabayar');
INSERT product_types (type) VALUES ('normal');

-- 1.f
INSERT product_descriptions (description) VALUES ('ini produk sangatlah bagus');
INSERT product_descriptions (description) VALUES ('produk ini best seller kami');
INSERT product_descriptions (description) VALUES ('siapa cepat dia dapat');
INSERT product_descriptions (description) VALUES ('tidak ada bandingannya');
INSERT product_descriptions (description) VALUES ('mengutamakan harga paket yang sangat terjangkau');
INSERT product_descriptions (description) VALUES ('internet menjangkau semua kawasan');
INSERT product_descriptions (description) VALUES ('internet sangat stabil');
INSERT product_descriptions (description) VALUES ('kecepatan no 1');


-- 1.c
INSERT products (product_type_id, operator_id, product_description_id, name) VALUES (1, 3, 1, "kartu perdana");
INSERT products (product_type_id, operator_id, product_description_id, name) VALUES (1, 3, 2, "perdana no.1");

-- 1.d
INSERT products (product_type_id, operator_id, product_description_id, name) VALUES (2, 1, 3, "simpati");
INSERT products (product_type_id, operator_id, product_description_id, name) VALUES (2, 1, 4, "loop");
INSERT products (product_type_id, operator_id, product_description_id, name) VALUES (2, 1, 5, "pilihanku");

-- 1.e
INSERT products (product_type_id, operator_id, product_description_id, name) VALUES (3, 4, 6, "normal banget");
INSERT products (product_type_id, operator_id, product_description_id, name) VALUES (3, 4, 7, "perdana");
INSERT products (product_type_id, operator_id, product_description_id, name) VALUES (3, 4, 8, "ekslusif");

-- 1.g
INSERT payment_methods (name, status) VALUES ('qris', 1);
INSERT payment_methods (name, status) VALUES ('BCA VA', 1);
INSERT payment_methods (name, status) VALUES ('Gopay', 1);

-- 1.h
INSERT users (name, status, dob, gender) VALUES ('budi', 1, '2008-11-11', 'M');
INSERT users (name, status, dob, gender) VALUES ('rama', 1, '2000-10-22', 'F');
INSERT users (name, status, dob, gender) VALUES ('windah', 1, '2002-03-04', 'M');
INSERT users (name, status, dob, gender) VALUES ('basudara', 1, '2002-02-11', 'F');
INSERT users (name, status, dob, gender) VALUES ('ilham', 1, '2001-04-23', 'M');

-- 1.i
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (1, 1, 'paid', 3, 50000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (2, 2, 'paid', 3, 60000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (3, 3, 'paid', 3, 75000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (4, 1, 'paid', 3, 50000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (5, 2, 'paid', 3, 60000);

INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (1, 1, 'paid', 3, 50000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (2, 2, 'paid', 3, 60000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (3, 3, 'paid', 3, 75000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (4, 1, 'paid', 3, 50000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (5, 2, 'paid', 3, 60000);

INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (1, 1, 'paid', 3, 50000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (2, 2, 'paid', 3, 60000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (3, 3, 'paid', 3, 75000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (4, 1, 'paid', 3, 50000);
INSERT transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES (5, 2, 'paid', 3, 60000);

-- i.j
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 1, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 1, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 1, 'paid', 1, 15000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 2, 'paid', 1, 40000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 2, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 2, 'paid', 1, 10000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 3, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 3, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 3, 'paid', 1, 25000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 4, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 4, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 4, 'paid', 1, 15000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 5, 'paid', 1, 40000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 5, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 5, 'paid', 1, 10000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 6, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 6, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 6, 'paid', 1, 15000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 7, 'paid', 1, 40000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 7, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 7, 'paid', 1, 10000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 8, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 8, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 8, 'paid', 1, 25000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 9, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 9, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 9, 'paid', 1, 15000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 10, 'paid', 1, 40000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 10, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 10, 'paid', 1, 10000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 11, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 11, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 11, 'paid', 1, 15000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 12, 'paid', 1, 40000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 12, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 12, 'paid', 1, 10000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 13, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 13, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 13, 'paid', 1, 25000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 14, 'paid', 1, 25000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 14, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 14, 'paid', 1, 15000);

insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (1, 15, 'paid', 1, 40000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (2, 15, 'paid', 1, 10000);
insert transaction_details (product_id, transaction_id, status, qty, price) VALUES (3, 15, 'paid', 1, 10000);

-- 2.a
SELECT * FROM users WHERE gender='M';

-- 2.b
SELECT * FROM products WHERE id=3;

-- 2.c (belum selesai)
SELECT * FROM users WHERE CREATED_AT >= DATE_ADD(CURRENT_TIMESTAMP, INTERVAL -7 DAY) AND name LIKE '%a%';

-- 2.d
SELECT COUNT(gender) FROM users WHERE gender='F';

-- 2.e
SELECT * FROM users ORDER BY name;

-- 2.f
SELECT * FROM products LIMIT 5;

-- 3.a
UPDATE products SET name='product dummy' WHERE id=1;

-- 3.b
UPDATE transaction_details SET qty=3 WHERE id=1;

-- 4.a
DELETE FROM transaction_details WHERE product_id=1;
DELETE FROM products WHERE id=1;

-- 4.b
DELETE FROM transaction_details WHERE product_id=2;
DELETE FROM products WHERE product_type_id=1;