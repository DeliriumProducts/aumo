import {
  Avatar,
  Button,
  Form,
  Icon,
  Input,
  List,
  Modal,
  Popconfirm,
  Tooltip
} from "antd"
import aumo from "aumo"

export default Form.create()(
  // eslint-disable-next-line
  class extends React.Component {
    state = {
      shopOwners: []
    }

    async componentDidUpdate(prevProps) {
      if (this.props.shop?.id !== prevProps.shop?.id) {
        try {
          const { owners: shopOwners } = await aumo.shop.getShop(
            this.props.shop.id
          )
          this.setState({ shopOwners })
        } catch (error) {}
      }
    }

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
                  placeholder="Email"
                  prefix={
                    <Icon type="user" style={{ color: "rgba(0,0,0,.25)" }} />
                  }
                  enterButton={<Button type="primary">Add</Button>}
                />
              )}
            </Form.Item>
            <Form.Item label="Current owners of this shop">
              <List
                dataSource={this.state.shopOwners}
                renderItem={owner => (
                  <List.Item key={owner.id}>
                    <List.Item.Meta
                      avatar={<Avatar src={owner.avatar} />}
                      title={owner.name}
                      description={owner.email}
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
