package mysql

const Schema = `
CREATE TABLE IF NOT EXISTS users (
  id int(10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  avatar varchar(255) NOT NULL,
  points double NOT NULL,
  UNIQUE KEY email (email)
);

CREATE TABLE IF NOT EXISTS products (
  id int (10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  name varchar (255) DEFAULT NULL,
  price double DEFAULT NULL,
  image varchar (255) DEFAULT NULL,
  description varchar (255) DEFAULT NULL,
  stock int (10) unsigned DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS receipts (
  id int (10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  content varchar(255) NOT NULL,
  user_id int(10) unsigned NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders (
    user_id int (10) unsigned NOT NULL,
    product_id int (10) unsigned NOT NULL,
    quantity int (11) NOT NULL,
    PRIMARY KEY (user_id, product_id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
`
