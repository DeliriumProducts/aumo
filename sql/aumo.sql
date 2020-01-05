CREATE TABLE users (
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

CREATE TABLE products (
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

CREATE TABLE receipts (
  id int (10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  content varchar(255) NOT NULL,
  user_id int(10) unsigned NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE orders (
    userId int (10) unsigned NOT NULL,
    productId int (10) unsigned NOT NULL,
    quantity int (11) NOT NULL,
    PRIMARY KEY (userId, productId),
    FOREIGN KEY (productId) REFERENCES products(id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
