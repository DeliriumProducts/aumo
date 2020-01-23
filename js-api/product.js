const axios = require('axios');
const opts = {
  withCredentials: true
};

class ProductAPI {
  constructor(backend) {
    this.backend = backend;
  }

  async getAll() {
    return await axios.get(`${this.backend}/products`, opts);
  }

  async get(id) {
    return await axios.get(`${this.backend}/products/${id}`, opts);
  }

  async create(product) {
    return await axios.post(`${this.backend}/products`, product, opts);
  }

  async edit(id, product) {
    return await axios.put(`${this.backend}/products/${id}`, product, opts);
  }

  async delete(id) {
    return await axios.delete(`${this.backend}/products/${id}`, opts);
  }
}

export default ProductAPI;
