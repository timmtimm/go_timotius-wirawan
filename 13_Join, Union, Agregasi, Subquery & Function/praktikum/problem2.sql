-- 1
SELECT * FROM transactions WHERE user_id=1
UNION
SELECT * FROM transactions WHERE user_id=2;

-- 2
SELECT SUM(total_price) FROM transactions WHERE user_id=1;

-- 3
SELECT COUNT(DISTINCT transactions.id) FROM transactions INNER JOIN transaction_details ON transaction_details.product_id=2;

-- 4
SELECT * FROM products INNER JOIN product_types ON products.id = product_types.id;

-- 5
SELECT transactions.*, products.name, users.name FROM transactions
INNER JOIN products ON products.id = transactions.id
INNER JOIN users ON users.id = transactions.id;

-- 6
DELIMITER $$

CREATE TRIGGER deleteTransactionDetail
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END$$

DELIMITER ;

--7 
DELIMITER $$

CREATE TRIGGER updateTransactionAfterTransactionDetail
AFTER DELETE ON transaction_details FOR EACH ROW
BEGIN
    UPDATE transactions
    SET
        total_qty = transactions.total_qty - OLD.qty
    WHERE transactions.id = OLD.transaction_id;
END$$

DELIMITER ;


-- 8
SELECT * FROM products WHERE id NOT IN (
    SELECT DISTINCT product_id FROM transaction_details
);