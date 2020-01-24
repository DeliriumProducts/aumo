INSERT INTO `users` VALUES 
    (1,'Simo','simo3003@me.com','123','https://github.com/fr3fou.png', 5000, "Admin"),
    (2,'Pesho','pesho@me.com','1234','https://github.com/pesho.png', 5000, "Customer"),
    (3,'Lyubo','lyubo@me.com','1234234','https://github.com/impzero.png', 5000, "Admin");
INSERT INTO `products` VALUES 
    (1,'Akai 32-inch HD LED LCD',500,'https://azcd.harveynorman.com.au/media/catalog/product/cache/21/image/992x558/9df78eab33525d08d6e5fb8d27136e95/a/k/ak3219nf.jpg',"Enjoy watching your favourite movies and shows in stunning HD quality with the Akai 32-inch HD LED LCD Smart TV.",5),
    (2,'CYBERPOWERPC Gamer Master Gaming PC',500,'https://images-na.ssl-images-amazon.com/images/I/812kz16Md0L._SX466_.jpg',"Cyber PowerPC Gamer Master series is a line of gaming PCs powered by AMD's newest Ryzen CPU and accompanying AM4 architecture.",5),
    (3,'iPhone 11',500,'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone11-select-2019-family?wid=882&amp;hei=1058&amp;fmt=jpeg&amp;qlt=80&amp;op_usm=0.5,0.5&amp;.v=1567022175704',"Shoot amazing videos and photos with the Ultra Wide, Wide, and Telephoto cameras. Capture your best low-light photos with Night mode. ",5);
INSERT INTO `receipts` VALUES 
    (1,"Paconi: 250 leva",1),
    (2,"Lidl: 500 leva",2);
INSERT INTO `orders` VALUES 
    (1, 1, 2),
    (2, 1, 1),
    (3, 1, 3),
    (4, 2, 1);
