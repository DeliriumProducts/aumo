import {
  Avatar,
  Button,
  Form,
  Icon,
  Input,
  List,
  Modal,
  Popconfirm,
  Tooltip,
  message
} from "antd"
import aumo from "aumo"

export default Form.create()(
  // eslint-disable-next-line
  class extends React.Component {
    state = {
      shopOwners: [],
      loading: true
    }

    async componentDidUpdate(prevProps) {
      if (this.props.shop?.id !== prevProps.shop?.id) {
        try {
          this.setState({ loading: true })
          const { owners: shopOwners } = await aumo.shop.getShop(
            this.props.shop.id
          )
          this.setState({ shopOwners })
        } catch (error) {
        } finally {
          this.setState({ loading: false })
        }
      }
    }

    render() {
      const { visible, onCancel, onAdd, form, shop, currentUser } = this.props
      const { getFieldDecorator } = form

      return (
        <Modal
          visible={visible}
          title={"Manage shop owners"}
          okText="Submit"
          onCancel={onCancel}
          footer={null}
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
                  onSearch={async () => {
                    form.validateFields(async (err, data) => {
                      if (err) {
                        return
                      }

                      try {
                        await aumo.shop.addOwner(shop.id, data.ownerEmail)
                        message.success(`Successfully added new owner! 🎉`)

                        const { owners: shopOwners } = await aumo.shop.getShop(
                          this.props.shop.id
                        )

                        this.setState({ shopOwners })
                      } catch (err) {
                        if (!err.response) {
                          message.error(`${err}`, 5)
                          return
                        }
                        if (err.response.status === 401) {
                          message.error("Unauthorized.", 1)
                        } else if (err.response.status === 404) {
                          message.error("User not found")
                        } else {
                          message.error("Server error, please try again")
                        }
                        return
                      }
                      form.resetFields()
                    })
                  }}
                />
              )}
            </Form.Item>
            <Form.Item label="Current owners of this shop">
              {this.state.loading ? (
                <Icon type="loading" style={{ fontSize: 24 }} spin />
              ) : (
                <List
                  dataSource={this.state.shopOwners}
                  renderItem={owner => (
                    <List.Item key={owner.id}>
                      <List.Item.Meta
                        avatar={<Avatar src={owner.avatar} />}
                        title={owner.name}
                        description={owner.email}
                      />
                      {owner.email !== currentUser.email && (
                        <Tooltip
                          title={
                            owner.email === currentUser.email
                              ? "You cannot remove yourself!"
                              : "Remove this owner!"
                          }
                        >
                          <Popconfirm
                            onConfirm={async e => {
                              e.stopPropagation()
                              try {
                                await aumo.shop.removeOwner(
                                  shop.id,
                                  owner.email
                                )
                                message.success(
                                  `Successfully deleted shop owner ${owner.name}! 🎉`
                                )

                                const {
                                  owners: shopOwners
                                } = await aumo.shop.getShop(this.props.shop.id)

                                this.setState({ shopOwners })
                              } catch (err) {
                                if (!err.response) {
                                  message.error(`${err}`, 5)
                                  return
                                }
                                if (err.response.status === 401) {
                                  message.error("Unathorized. Try again.", 1)
                                } else {
                                  message.error(
                                    "Server error, please try again"
                                  )
                                }
                                return
                              }
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
                              onClick={async e => {
                                e.stopPropagation()
                              }}
                              style={{
                                right: 10
                              }}
                            ></Button>
                          </Popconfirm>
                        </Tooltip>
                      )}
                    </List.Item>
                  )}
                ></List>
              )}
            </Form.Item>
          </Form>
        </Modal>
      )
    }
  }
)
