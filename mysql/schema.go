package mysql

// Schema contains all the MySQL tables in Aumo
const Schema = `
CREATE TABLE IF NOT EXISTS users (
  id int(10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) NOT NULL,
  role varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  avatar varchar(255) NOT NULL,
  points double NOT NULL,
  UNIQUE KEY email (email)
);

CREATE TABLE IF NOT EXISTS products (
  id int (10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) DEFAULT NULL,
  price double DEFAULT NULL,
  image varchar(255) DEFAULT NULL,
  description varchar (255) DEFAULT NULL,
  stock int (10) unsigned DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS receipts (
  receipt_id int (10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  content TEXT NOT NULL,
  user_id int(10) unsigned NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders (
	order_id int (10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id int (10) unsigned NOT NULL,
  product_id int (10) unsigned NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE NO ACTION ON UPDATE NO ACTION,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
`
