const axios = require("axios")

class ProductAPI {
  static opts = {
    withCredentials: true
  }

  constructor(backend) {
    this.backend = backend
  }

  async getAll() {
    return await axios.get(`${this.backend}/products`, this.opts)
  }

  async get(id) {
    return await axios.get(`${this.backend}/products/${id}`, this.opts)
  }

  async create(product) {
    return await axios.post(`${this.backend}/products`, product, this.opts)
  }

  async edit(id, product) {
    return await axios.put(`${this.backend}/products/${id}`, product, this.opts)
  }

  async delete(id) {
    return await axios.delete(`${this.backend}/products/${id}`, this.opts)
  }
}

export default ProductAPI
