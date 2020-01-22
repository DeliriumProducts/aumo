const axios = require("axios")

class OrderAPI {
  static opts = {
    withCredentials: true
  }

  constructor(backend) {
    this.backend = backend
  }

  async place(order) {
    return await axios.post(`${this.backend}/orders`, order, this.opts)
  }

  async getAll() {
    return await axios.get(`${this.backend}/orders`, this.opts)
  }

  async get(id) {
    return await axios.get(`${this.backend}/orders/${id}`, this.opts)
  }
}

export default OrderAPI
