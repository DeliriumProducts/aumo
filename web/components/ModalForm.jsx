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
                ]
              })(<Input value={product?.name} />)}
            </Form.Item>
            <Form.Item label="Image">
              {getFieldDecorator("image", {
                rules: [
                  {
                    required: true
                  }
                ]
              })(<Input value={product?.image} />)}
            </Form.Item>
            <Form.Item label="Price">
              {getFieldDecorator("price", {
                rules: [
                  {
                    required: true
                  }
                ]
              })(<Input value={product?.price} />)}
            </Form.Item>
            <Form.Item label="Description">
              {getFieldDecorator("description", {
                rules: [
                  {
                    required: true
                  }
                ]
              })(<Input type="textarea" value={product?.description} />)}
            </Form.Item>
          </Form>
        </Modal>
      )
    }
  }
)
