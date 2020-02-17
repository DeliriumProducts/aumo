import {
  Button,
  Icon,
  Layout,
  List,
  Modal,
  Spinner,
  Tab,
  TabView,
  Text
} from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { View } from "react-native"
import styled from "styled-components/native"
import PAvatar from "../../../components/Avatar"
import Order from "../../../components/Order"
import { Context } from "../../../context/context"
import { actions } from "../../../context/providers/provider"
import Routes from "../../../navigation/routes"

export default ({ navigation }) => {
  const ctx = React.useContext(Context)
  const [loading, setLoading] = React.useState(false)
  const [tabIdx, setTabIdx] = React.useState(0)

  const logout = async () => {
    try {
      setLoading(true)
      await aumo.auth.logout()
      setLoading(false)
      ctx.dispatch({ type: actions.SET_USER, payload: null })
    } catch (error) {
      setLoading(false)
    }
  }

  return (
    <>
      <MainLayout level="1">
        <ProfileContainer
          style={{ flexDirection: "row", justifyContent: "space-between" }}
        >
          <View style={{ flexDirection: "row" }}>
            <Avatar
              size="giant"
              source={{ uri: ctx?.state?.user?.avatar }}
              fallbackSource={require("../../../assets/Avatar.png")}
            />
            <View style={{ marginLeft: 10 }}>
              <Text category="h2">{ctx?.state?.user?.name}</Text>
              <Text appearance="hint" category="s1">
                {ctx?.state?.user?.email}
              </Text>
            </View>
          </View>
          <Button
            disabled={loading}
            size="medium"
            status="basic"
            appearance="ghost"
            icon={style => <Icon name="log-out-outline" {...style} />}
            onPress={logout}
          />
        </ProfileContainer>
        <View style={{ width: "90%", alignSelf: "center" }}>
          <Stats>
            <Stat hint="Receipts" value={ctx?.state?.user?.receipts.length} />
            <Stat hint="Orders" value={ctx?.state?.user?.orders.length} />
            <Stat hint="Points" value={ctx?.state?.user?.points} />
          </Stats>
          <EditButton
            icon={style => <Icon name="edit-outline" {...style} />}
            onPress={() => {
              navigation.push(Routes.UserEdit)
            }}
          >
            EDIT PROFILE
          </EditButton>
        </View>
        <Modal
          backdropStyle={{
            backgroundColor: "rgba(0, 0, 0, 0.5)"
          }}
          onBackdropPress={() => {}}
          visible={loading}
        >
          <ModalContainer level="1">
            {loading && <Spinner size="giant" />}
          </ModalContainer>
        </Modal>
      </MainLayout>
      <TabView selectedIndex={tabIdx} onSelect={setTabIdx} style={{ flex: 1 }}>
        <Tab
          title="Orders"
          icon={style => <Icon {...style} name="bell-outline" />}
        >
          <List
            data={ctx.state.user.orders}
            renderItem={({ item: order }) => (
              <View style={{ margin: 10 }}>
                <Order product={order.product} key={order.order_id} />
              </View>
            )}
          />
        </Tab>
        <Tab
          title="Receipts"
          icon={style => <Icon {...style} name="file-text-outline" />}
        >
          <Layout>
            {ctx.state.user?.receipts.map(a => (
              <Text key={a.receipt_id}>{a.receipt_id}</Text>
            ))}
          </Layout>
        </Tab>
      </TabView>
    </>
  )
}

const Stat = ({ hint, value }) => {
  return (
    <StatContainer>
      <Text category="s2">{value}</Text>
      <Text appearance="hint" category="c2">
        {hint}
      </Text>
    </StatContainer>
  )
}

const MainLayout = styled(Layout)`
  margin-horizontal: -16px;
  padding-horizontal: 16px;
  padding-top: 16px;
`

const ModalContainer = styled(Layout)`
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  padding: 16px;
`

const ProfileContainer = styled(View)`
  flex-direction: row;
  justify-content: space-between;
`

const Stats = styled(View)`
  flex-direction: row;
  margin-top: 24px;
`

const Avatar = styled(PAvatar)`
  margin-horizontal: 8px;
`

const EditButton = styled(Button)`
  margin-vertical: 16px;
  border-radius: 10px;
`

const StatContainer = styled(View)`
  align-items: center;
  width: 100%;
  justify-content: center;
  flex: 1;
`
