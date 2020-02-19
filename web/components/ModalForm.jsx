import { Modal, Form, Input } from "antd"

export default Form.create()(
  // eslint-disable-next-line
  class extends React.Component {
    render() {
      const {
        visible,
        onCancel,
        onCreate,
        form,
        entity,
        isProduct
      } = this.props
      const { getFieldDecorator } = form

      return (
        <Modal
          visible={visible}
          title="Product Form"
          okText="Submit"
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
                initialValue: entity?.name
              })(<Input />)}
            </Form.Item>
            <Form.Item label="Image">
              {getFieldDecorator("image", {
                rules: [
                  {
                    required: true
                  }
                ],
                initialValue: entity?.image
              })(<Input />)}
            </Form.Item>
            {isProduct ? (
              <>
                <Form.Item label="Price">
                  {getFieldDecorator("price", {
                    rules: [
                      {
                        required: true
                      }
                    ],
                    initialValue: entity?.price
                  })(<Input />)}
                </Form.Item>
                <Form.Item label="Stock">
                  {getFieldDecorator("stock", {
                    rules: [
                      {
                        required: true
                      }
                    ],
                    initialValue: entity?.stock
                  })(<Input />)}
                </Form.Item>
                <Form.Item label="Description">
                  {getFieldDecorator("description", {
                    rules: [
                      {
                        required: true
                      }
                    ],

                    initialValue: entity?.description
                  })(<Input type="textarea" />)}
                </Form.Item>
              </>
            ) : (
              <></>
            )}
          </Form>
        </Modal>
      )
    }
  }
)
