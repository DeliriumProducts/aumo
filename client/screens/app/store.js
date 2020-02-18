import { Layout, List } from "@ui-kitten/components"
import React from "react"
import Order from "../../components/Order"
import { Context } from "../../context/context"

export default () => {
  const ctx = React.useContext(Context)

  return (
    <Layout>
      <List
        data={ctx.state.user.orders}
        renderItem={({ item: order }) => (
          <Order product={order.product} key={order.order_id} />
        )}
      />
    </Layout>
  )
}
