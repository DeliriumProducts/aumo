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

CREATE TABLE IF NOT EXISTS shops (
  shop_id int(10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) NOT NULL,
  image varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS shop_owners (
  shop_id int(10) unsigned NOT NULL,
  user_id varchar(36) NOT NULL,
  PRIMARY KEY (shop_id, user_id),
  FOREIGN KEY (shop_id) REFERENCES shops(shop_id) ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS products (
  id int(10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) DEFAULT NULL,
  price double DEFAULT NULL,
  image varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  stock int(10) unsigned DEFAULT NULL,
  shop_id int(10) unsigned NOT NULL,
  FOREIGN KEY (shop_id) REFERENCES shops(shop_id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS orders (
  order_id varchar(36) NOT NULL PRIMARY KEY,
  user_id varchar(36) NOT NULL,
  product_id int(10) unsigned NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS receipts (
  receipt_id varchar(36) NOT NULL PRIMARY KEY,
  content TEXT NOT NULL,
  total double NOT NULL,
  user_id varchar(36) NULL,
  shop_id int(10) unsigned NOT NULL,
  FOREIGN KEY (shop_id) REFERENCES shops(shop_id) ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
`
