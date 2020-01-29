import axios from 'axios';
import { BaseResponse, Order } from './aumo';
import { axiosRequest } from './axios';
import { options } from './config';

export async function place(order: PlaceRequest): Promise<BaseResponse<Order>> {
  return (await axios.post(`${options.Backend}/orders`, order, axiosRequest))
    .data;
}

interface PlaceRequest {
  product_id: number;
}

export default {
  place
};
