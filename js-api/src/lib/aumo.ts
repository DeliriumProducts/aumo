import { AxiosResponse } from 'axios';

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

export interface MessageResponse {
  Message: string;
}

export interface ErrorResponse {
  Err: string;
}

export interface ValidationErrorResponse {
  Errors: string[];
}

export type BaseResponse<T = {}> = AxiosResponse<
  MessageResponse | ErrorResponse | ValidationErrorResponse | T
>;
