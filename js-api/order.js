const axios = require("axios")
const opts = {
  withCredentials: true
}

class OrderAPI {
  constructor(backend) {
    this.backend = backend
  }

  async place(order) {
    return (await axios.post(`${this.backend}/orders`, order, opts)).data
  }

  async getAll() {
    return (await axios.get(`${this.backend}/orders`, opts)).data
  }

  async get(id) {
    return (await axios.get(`${this.backend}/orders/${id}`, opts)).data
  }
}

module.exports = OrderAPI
