-- +goose Up
-- +goose StatementBegin
CREATE TABLE customer_loan (
   id INT AUTO_INCREMENT PRIMARY KEY,
   customer_id INT,
   status ENUM('Pending', 'Approved', 'Rejected', 'Paid') NOT NULL,
   loan_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   amount DECIMAL(10, 2) NOT NULL,
   approved_date TIMESTAMP DEFAULT NULL,
   tenor INT NOT NULL,
   interest FLOAT NOT NULL,
   admin_fee FLOAT NOT NULL,
   is_active BOOLEAN DEFAULT true,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP DEFAULT NULL,
   FOREIGN KEY (customer_id) REFERENCES customer(id),
   FOREIGN KEY (tenor) REFERENCES tenor(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customer_loan;
-- +goose StatementEnd
