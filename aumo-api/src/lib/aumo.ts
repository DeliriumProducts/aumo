export interface Product {
  id: string;
  name: string;
  price: number;
  image: string;
  description: string;
  stock: number;
  shop_id: number;
  shop: Shop;
}

export interface User {
  id?: string;
  name: string;
  email: string;
  password: string;
  avatar: string;
  points: number;
  role: Role;
  orders: Order[];
  receipts: Receipt[];
  is_verified: boolean;
  shops: Shop[];
}

export type Role = 'Admin' | 'Customer' | 'Shop Owner';

export interface Order {
  order_id: string;
  user_id: string;
  product_id: string;
  product: Product;
}

export interface Shop {
  id: number;
  name: string;
  image: string;
  owners: User[];
  products: Product[];
}

export interface Receipt {
  receipt_id: string;
  content: string;
  total: number;
  shop_id: number;
  shop: Shop;
}

export interface MessageResponse {
  message: string;
}

export interface ErrorResponse {
  err: string;
}

export interface ValidationErrorResponse {
  errors: string[];
}
