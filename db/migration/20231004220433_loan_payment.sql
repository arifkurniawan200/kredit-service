-- +goose Up
-- +goose StatementBegin
CREATE TABLE loan_payment (
  id INT AUTO_INCREMENT PRIMARY KEY,
  customer_loan_id INT NOT NULL,
  payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  amount DECIMAL(10, 2) NOT NULL,
  status ENUM('Pending', 'Paid', 'Late') NOT NULL,
  due_date TIMESTAMP NOT NULL,
  late_fee DECIMAL(10, 2) DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
  FOREIGN KEY (customer_loan_id) REFERENCES customer_loan(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE loan_payment;
-- +goose StatementEnd
