import axios from 'axios';
import { BaseResponse, Product } from './aumo';
import { axiosRequest, withAuth } from './axios';
import { options } from './config';

export async function getAllProducts(): Promise<BaseResponse<Product[]>> {
  return (await axios.get(`${options.Backend}/products`, axiosRequest)).data;
}

export async function getProduct(id: string): Promise<BaseResponse<Product>> {
  return (await axios.get(`${options.Backend}/products/${id}`, axiosRequest))
    .data;
}

export async function createProduct(
  product: CreateRequest,
  cookie?: string
): Promise<BaseResponse<Product>> {
  return (
    await axios.post(`${options.Backend}/products`, product, withAuth(cookie))
  ).data;
}

export async function editProduct(
  id: string,
  product: EditRequest,
  cookie?: string
): Promise<BaseResponse<Product>> {
  return (
    await axios.put(
      `${options.Backend}/products/${id}`,
      product,
      withAuth(cookie)
    )
  ).data;
}

export async function deleteProduct(id: string, cookie?: string) {
  return (
    await axios.delete(`${options.Backend}/products/${id}`, withAuth(cookie))
  ).data;
}

interface CreateRequest {
  name: string;
  image: string;
  price: number;
  description: string;
  stock: number;
}

interface EditRequest extends CreateRequest {}
