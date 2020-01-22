const axios = require("axios")

class ReceiptAPI {
  static opts = {
    withCredentials: true
  }

  constructor(backend) {
    this.backend = backend
  }

  async claim(id) {
    return await axios.get(`${this.backend}/receipts/${id}`, this.opts)
  }

  async create(receipt) {
    return await axios.post(`${this.backend}/receipts`, receipt, this.opts)
  }
}

export default ReceiptAPI
