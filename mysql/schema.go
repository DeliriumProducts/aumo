package mysql

// Schema contains all the MySQL tables in Aumo
const Schema = `
CREATE TABLE IF NOT EXISTS users (
  id varchar(36) NOT NULL PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  avatar varchar(255) NOT NULL,
  points double NOT NULL,
  role varchar(255) NOT NULL,
  verified tinyint(1) NOT NULL,
  UNIQUE KEY email (email)
);

CREATE TABLE IF NOT EXISTS products (
  id int(10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) DEFAULT NULL,
  price double DEFAULT NULL,
  image varchar(255) DEFAULT NULL,
  description varchar (255) DEFAULT NULL,
  stock int(10) unsigned DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS receipts (
  receipt_id varchar(36) NOT NULL PRIMARY KEY,
  content TEXT NOT NULL,
  user_id varchar(36) NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders (
  order_id varchar(36) NOT NULL PRIMARY KEY,
  user_id varchar(36) NOT NULL,
  product_id int(10) unsigned NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE NO ACTION ON UPDATE NO ACTION,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
`
