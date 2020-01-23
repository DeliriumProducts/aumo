const axios = require('axios');
const opts = {
  withCredentials: true
};

class ReceiptAPI {
  constructor(backend) {
    this.backend = backend;
  }

  async claim(id) {
    return await axios.get(`${this.backend}/receipts/${id}`, opts).data;
  }

  async create(receipt) {
    return await axios.post(`${this.backend}/receipts`, receipt, opts).data;
  }
}

module.exports = ReceiptAPI;
