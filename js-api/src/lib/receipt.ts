import axios from 'axios';
import { BaseResponse, Receipt } from './aumo';
import { withAuth } from './axios';
import { options } from './config';

export async function claimReceipt(
  id: string,
  cookie?: string
): Promise<BaseResponse<Receipt>> {
  return (
    await axios.get(`${options.Backend}/receipts/${id}`, withAuth(cookie))
  ).data;
}

export async function createReceipt(
  receipt: ReceiptRequest,
  cookie?: string
): Promise<BaseResponse<Receipt>> {
  return (
    await axios.post(`${options.Backend}/receipts`, receipt, withAuth(cookie))
  ).data;
}

interface ReceiptRequest {
  content: string;
}
