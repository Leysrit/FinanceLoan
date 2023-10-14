-- +migrate Up        
DROP TABLE IF EXISTS limit_loan;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS customers;
-- +migrate Up  
CREATE TABLE IF NOT EXISTS customers (
	CustomerID INT AUTO_INCREMENT PRIMARY KEY ,
	NIK VARCHAR(16) NOT NULL,
	FullName VARCHAR(255) NOT NULL,
	LegalName VARCHAR(255) NOT NULL,
	PlaceOfBirth VARCHAR(100),
	DateOfBirth DATE,
	Salary DECIMAL(10, 2),
	KTPImage VARCHAR(255),
	SelfieImage VARCHAR(255)
);
-- +migrate Up  
CREATE TABLE IF NOT EXISTS limit_loan (
	LimitID INT PRIMARY KEY,
	CustomerID INT NOT NULL,
	TenorMonths INT,
	LimitAmount DECIMAL(10, 2),
	FOREIGN KEY (CustomerID) REFERENCES customers(CustomerID)
		);
-- +migrate Up  
CREATE TABLE IF NOT EXISTS transactions (
	ContractNumber INT PRIMARY KEY,
	CustomerID INT NOT NULL,
	OTR DECIMAL(10, 2),
	AdminFee DECIMAL(10, 2),
	InstallmentAmount DECIMAL(10, 2),
	InterestAmount DECIMAL(10, 2),
	AssetName VARCHAR(255),
	FOREIGN KEY (CustomerID) REFERENCES customers(CustomerID)
);
-- +migrate Up  
INSERT INTO customers (CustomerID, NIK, FullName, LegalName, PlaceOfBirth, DateOfBirth, Salary, KTPImage, SelfieImage)
VALUES
(1, '1234567890', 'Budi Santoso', 'Budi Santoso', 'Jakarta', '1990-05-15', 5000000.00, 'path_to_ktp_image_budi.jpg', 'path_to_selfie_image_budi.jpg'),
(2, '9876543210', 'Annisa Putri', 'Annisa Putri', 'Surabaya', '1985-12-10', 7000000.00, 'path_to_ktp_image_annisa.jpg', 'path_to_selfie_image_annisa.jpg');
-- +migrate Up  
INSERT INTO limit_loan (LimitID, CustomerID, TenorMonths, LimitAmount)
VALUES
	(1, 1, 1, 100000.00), 
	(2, 1, 2, 200000.00), 
	(3, 1, 3, 500000.00), 
	(4, 1, 4, 700000.00), 
	(5, 2, 1, 1000000.00), 
	(6, 2, 2, 1200000.00),
	(7, 2, 3, 1500000.00),
	(8, 2, 4, 2000000.00);
-- +migrate Up  
INSERT INTO transactions (ContractNumber, CustomerID, OTR, AdminFee, InstallmentAmount, InterestAmount, AssetName)
VALUES
	(101, 1, 15000000.00, 100000.00, 400000.00, 5000.00, 'Motor Yamaha XMAX');

-- +migrate Down        
        DROP TABLE IF EXISTS limit_loan;
        DROP TABLE IF EXISTS customers;
        DROP TABLE IF EXISTS transactions;