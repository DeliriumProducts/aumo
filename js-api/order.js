const axios = require("axios")
const opts = {
  withCredentials: true
}

class OrderAPI {
  constructor(backend) {
    this.backend = backend
  }

  async place(order) {
    return await axios.post(`${this.backend}/orders`, order, opts)
  }

  async getAll() {
    return await axios.get(`${this.backend}/orders`, opts)
  }

  async get(id) {
    return await axios.get(`${this.backend}/orders/${id}`, opts)
  }
}

module.exports = OrderAPI
