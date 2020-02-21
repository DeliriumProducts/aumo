import axios from 'axios';
import { MessageResponse, Product, Shop } from './aumo';
import { withAuth } from './axios';
import { options } from './config';

export async function getAllShops(cookie?: string): Promise<Shop[]> {
  return (await axios.get(`${options.Backend}/shops`, withAuth(cookie))).data;
}

export async function getShop(id: number, cookie?: string): Promise<Shop> {
  return (await axios.get(`${options.Backend}/shops/${id}`, withAuth(cookie)))
    .data;
}

export async function editShop(id: number, shop: EditShopRequest, cookie?: string): Promise<Shop> {
  return (
    await axios.put(
      `${options.Backend}/shops/${id}`,
      shop,
      withAuth(cookie)
    )
  ).data;
}

export async function createShop(
  shop: CreateShopRequest,
  cookie?: string
): Promise<Shop> {
  return (await axios.post(`${options.Backend}/shops`, shop, withAuth(cookie)))
    .data;
}

export async function deleteShop(
  id: number,
  cookie?: string
): Promise<MessageResponse> {
  return (
    await axios.delete(`${options.Backend}/shops/${id}`, withAuth(cookie))
  ).data;
}

export async function addOwner(
  sID: number,
  email: string,
  cookie?: string
): Promise<MessageResponse> {
  return (
    await axios.post(
      `${options.Backend}/shops/${sID}/add-owner`,
      {
        email
      },
      withAuth(cookie)
    )
  ).data;
}

export async function removeOwner(
  sID: number,
  email: string,
  cookie?: string
): Promise<MessageResponse> {
  return (
    await axios.post(
      `${options.Backend}/shops/${sID}/remove-owner`,
      {
        email
      },
      withAuth(cookie)
    )
  ).data;
}

export async function getAllProductsByShop(sID: number, cookie?: string): Promise<Product[]> {
  return (await axios.get(`${options.Backend}/shops/${sID}/products`,
    withAuth(cookie)
  )).data;
}

export async function getProduct(sID: number, pID: number, cookie?: string): Promise<Product> {
  return (await axios.get(`${options.Backend}/shops/${sID}/products/${pID}`,
    withAuth(cookie)
  ))
    .data;
}

export async function createProduct(
  sID: number,
  product: CreateProductRequest,
  cookie?: string
): Promise<Product> {
  return (
    await axios.post(`${options.Backend}/shops/${sID}/products`, product, withAuth(cookie))
  ).data;
}

export async function editProduct(
  sID: number,
  pID: number,
  product: EditProductRequest,
  cookie?: string
): Promise<Product> {
  return (
    await axios.put(
      `${options.Backend}/shops/${sID}/products/${pID}`,
      product,
      withAuth(cookie)
    )
  ).data;
}

export async function deleteProduct(
  sID: number,
  pID: number,
  cookie?: string
): Promise<MessageResponse> {
  return (
    await axios.delete(`${options.Backend}/shops/${sID}/products/${pID}`, withAuth(cookie))
  ).data;
}

interface CreateProductRequest {
  name: string;
  image: string;
  price: number;
  description: string;
  stock: number;
}

interface CreateShopRequest {
  name: string;
  image: string;
}

interface EditProductRequest extends CreateProductRequest { }
interface EditShopRequest extends CreateShopRequest { }

export default {
  getAllShops,
  getShop,
  deleteShop,
  editShop,
  createShop,
  addOwner,
  removeOwner,
  editProduct,
  getProduct,
  deleteProduct,
  getAllProductsByShop,
};
