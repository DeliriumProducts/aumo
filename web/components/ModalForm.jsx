import { Modal, Form, Input } from "antd"

export default Form.create()(
  // eslint-disable-next-line
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form, product } = this.props
      const { getFieldDecorator } = form

      return (
        <Modal
          visible={visible}
          title="Edit product"
          okText="Edit"
          onCancel={onCancel}
          onOk={onCreate}
        >
          <Form layout="vertical">
            <Form.Item label="Name">
              {getFieldDecorator("name", {
                rules: [
                  {
                    required: true
                  }
                ],
                initialValue: product?.name
              })(<Input />)}
            </Form.Item>
            <Form.Item label="Image">
              {getFieldDecorator("image", {
                rules: [
                  {
                    required: true
                  }
                ],
                initialValue: product?.image
              })(<Input />)}
            </Form.Item>
            <Form.Item label="Price">
              {getFieldDecorator("price", {
                rules: [
                  {
                    required: true
                  }
                ],
                initialValue: product?.price
              })(<Input />)}
            </Form.Item>
            <Form.Item label="Stock">
              {getFieldDecorator("stock", {
                rules: [
                  {
                    required: true
                  }
                ],
                initialValue: product?.stock
              })(<Input />)}
            </Form.Item>
            <Form.Item label="Description">
              {getFieldDecorator("description", {
                rules: [
                  {
                    required: true
                  }
                ],

                initialValue: product?.description
              })(<Input type="textarea" />)}
            </Form.Item>
          </Form>
        </Modal>
      )
    }
  }
)
