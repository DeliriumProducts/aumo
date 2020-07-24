INSERT INTO `users` VALUES  -- all user's passwords are 123456
    ('24a2068c-a43d-43ae-b967-858953686189','Simo','simo3003@me.com','$2y$12$OzLPsg4GRS2ALckMY7QRIe.49zi2HnwL7kHYCbuMsR8Y4te3G.N9e','https://github.com/fr3fou.png', 5000, "Admin", 1),
    ('19274de1-ac2b-4dd9-adee-20f4ab69e920','Pesho','pesho@me.com','$2y$12$YRKTtih7Cy1kN9ZayAW9Fe.AOfZypr40io8byn4ghfJWWH.aZf2ca','https://github.com/pesho.png', 5000, "Customer", 1),
    ('1cd005fe-213f-4c96-9966-49306b00ca90','Lyubo','lyubo@me.com','$2y$12$w5Bs/vKa0D1uEKlo6AH8QexQujTjtSRIecBEwkSyo/nRzkR4KOdHO','https://github.com/impzero.png', 5000, "Shop Owner", 1);
INSERT INTO `shops` VALUES
    (1, "Paconi", "https://paconi.net/wp-content/uploads/2014/07/fav.png"),
    (2, "Lidl", "https://angliya.com/wp-content/uploads/2019/10/81093392_l-1024x683.jpg");
INSERT INTO `shop_owners` VALUES
    (1, '24a2068c-a43d-43ae-b967-858953686189'),
    (2, '1cd005fe-213f-4c96-9966-49306b00ca90');
INSERT INTO `products` VALUES 
    (1,'Akai 32-inch HD LED LCD',500,'https://azcd.harveynorman.com.au/media/catalog/product/cache/21/image/992x558/9df78eab33525d08d6e5fb8d27136e95/a/k/ak3219nf.jpg',"Enjoy watching your favourite movies and shows in stunning HD quality with the Akai 32-inch HD LED LCD Smart TV.",5, 2),
    (2,'CYBERPOWERPC Gamer Master Gaming PC',500,'https://images-na.ssl-images-amazon.com/images/I/812kz16Md0L._SX466_.jpg',"Cyber PowerPC Gamer Master series is a line of gaming PCs powered by AMD's newest Ryzen CPU and accompanying AM4 architecture.",5, 2),
    (3,'iPhone 11',500,'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone11-select-2019-family?wid=882&amp;hei=1058&amp;fmt=jpeg&amp;qlt=80&amp;op_usm=0.5,0.5&amp;.v=1567022175704',"Shoot amazing videos and photos with the Ultra Wide, Wide, and Telephoto cameras. Capture your best low-light photos with Night mode. ", 1, 1);
INSERT INTO `receipts` VALUES 
    ('acf4e079-688f-473d-b682-d9551a2527d5',"Paconi: 250 leva", 250, '24a2068c-a43d-43ae-b967-858953686189', 1),
    ('5021a0cf-e7cc-4ad1-a0c6-5e5a1c079303',"Lidl: 500 leva", 500, '19274de1-ac2b-4dd9-adee-20f4ab69e920', 2);
INSERT INTO `orders` VALUES 
    ('02ade587-62f4-4570-acfa-c2cb5e09f0c4', '24a2068c-a43d-43ae-b967-858953686189', 2),
    ('4f2e6fea-83f7-49d1-8be2-a3ddcaea0dcb', '19274de1-ac2b-4dd9-adee-20f4ab69e920', 1),
    ('4b75ddf3-334d-4120-9a41-1ced15f0d49c', '19274de1-ac2b-4dd9-adee-20f4ab69e920', 3),
    ('1f7d1e74-b37b-4537-8fcb-756911326249', '1cd005fe-213f-4c96-9966-49306b00ca90', 1);

