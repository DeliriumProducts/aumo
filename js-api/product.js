const axios = require("axios")
const opts = {
  withCredentials: true
}

class ProductAPI {
  constructor(backend) {
    this.backend = backend
  }

  async getAll() {
    return (await axios.get(`${this.backend}/products`, opts)).data
  }

  async get(id) {
    return (await axios.get(`${this.backend}/products/${id}`, opts)).data
  }

  async create(product) {
    return (await axios.post(`${this.backend}/products`, product, opts)).data
  }

  async edit(id, product) {
    return (await axios.put(`${this.backend}/products/${id}`, product, opts))
      .data
  }

  async delete(id) {
    return (await axios.delete(`${this.backend}/products/${id}`, opts)).data
  }
}

module.exports = ProductAPI
