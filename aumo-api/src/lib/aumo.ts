import { AxiosResponse } from 'axios';

export interface Product {
  id: string;
  name: string;
  price: number;
  image: string;
  description: string;
  stock: number;
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
}

export type Role = 'Admin' | 'Customer';

export interface Order {
  order_id: string;
  user_id: string;
  product_id: string;
  product: Product;
}

export interface Receipt {
  receipt_id: string;
  content: string;
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

export type BaseResponse<T = {}> = AxiosResponse<
  MessageResponse | ErrorResponse | ValidationErrorResponse | T
>;
