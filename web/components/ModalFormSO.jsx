import {
  Form,
  Input,
  Modal,
  Button,
  Icon,
  Avatar,
  List,
  Popconfirm,
  Tooltip
} from "antd"

export default Form.create()(
  // eslint-disable-next-line
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form, shop } = this.props
      const { getFieldDecorator } = form

      return (
        <Modal
          visible={visible}
          title={"Manage shop owners"}
          okText="Submit"
          onCancel={onCancel}
          onOk={onCreate}
        >
          <Form layout="vertical">
            <Form.Item label="Enter new owner email">
              {getFieldDecorator("ownerEmail", {
                rules: [
                  {
                    type: "email",
                    message: "The input is not valid Email!"
                  },
                  {
                    required: true,
                    message: "Please input your Email!"
                  }
                ]
              })(
                <Input.Search
                  placeholder="email"
                  prefix={
                    <Icon type="user" style={{ color: "rgba(0,0,0,.25)" }} />
                  }
                  enterButton={<Button type="primary">Add</Button>}
                />
              )}
            </Form.Item>
            <Form.Item label="Current owners of this shop">
              <List
                dataSource={["dasadasdsadsadsadsad"]}
                renderItem={item => (
                  <List.Item key={item}>
                    <List.Item.Meta
                      avatar={
                        <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
                      }
                      title={<a href="https://ant.design">Georgi Boyadjiev</a>}
                      description="Very bad guy"
                    />
                    <Tooltip title="Remove this owner!">
                      <Popconfirm
                        onConfirm={e => {
                          e.stopPropagation()
                        }}
                        disabled={false}
                        title={`Are you sure?`}
                        placement="bottom"
                        okText="Yes"
                        okType="danger"
                        onCancel={e => e.stopPropagation()}
                      >
                        <Button
                          type="danger"
                          icon="delete"
                          disabled={false}
                          onClick={e => e.stopPropagation()}
                          style={{
                            right: 10
                          }}
                        ></Button>
                      </Popconfirm>
                    </Tooltip>
                  </List.Item>
                )}
              ></List>
            </Form.Item>
          </Form>
        </Modal>
      )
    }
  }
)
