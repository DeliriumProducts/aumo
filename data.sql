INSERT INTO `users` VALUES 
    (1,'Simo','simo3003@me.com','123','github.com/fr3fou.png',5000, "Admin"),
    (2,'Pesho','pesho@me.com','1234','github.com/pesho.png',5000, "Customer"),
    (3,'Lyubo','lyubo@me.com','1234234','github.com/impzero.png',5000, "Admin");
INSERT INTO `products` VALUES 
    (1,'TV',500,'image',"it's decent",5),
    (2,'PC',500,'PC ok',"it's bad",5),
    (3,'Phone',500,'iphone ok',"it's meh",5);
INSERT INTO `receipts` VALUES 
    (1,"Paconi: 250 leva",1),
    (2,"Lidl: 500 leva",2);
INSERT INTO `orders` VALUES 
    (1, 1, 2),
    (2, 1, 1),
    (3, 1, 3),
    (4, 2, 1);