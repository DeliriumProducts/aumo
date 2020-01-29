export interface Product {
  id: number;
  name: string;
  price: number;
  image: string;
  description: string;
  stock: number;
}

export interface User {
  id?: number;
  name: string;
  email: string;
  password: string;
  avatar: string;
  points: number;
  role: Role;
  orders: Order[];
  receipts: Receipt[];
}

export type Role = 'Admin' | 'Customer';

export interface Order {
  order_id: number;
  user_id: number;
  product_id: number;
  product: Product;
}

export interface Receipt {
  receipt_id: number;
  content: string;
}
