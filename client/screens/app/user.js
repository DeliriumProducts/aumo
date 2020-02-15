import {
  Avatar as KAvatar,
  Button,
  Icon,
  Layout,
  Modal,
  Spinner,
  Text
} from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { View } from "react-native"
import styled from "styled-components/native"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/provider"

export default () => {
  const ctx = React.useContext(Context)
  const [loading, setLoading] = React.useState(false)

  const logout = async () => {
    try {
      setLoading(true)
      await aumo.auth.logout()
      setLoading(false)
      ctx.dispatch({ type: actions.SET_USER, payload: null })
    } catch (error) {
      console.warn(error)
      setLoading(false)
    }
  }

  return (
    <>
      <MainLayout level="1">
        <Avatar size="giant" source={{ uri: ctx?.state?.user?.avatar }} />
        <ProfileContainer>
          <MainContainer
            style={{ flexDirection: "row", justifyContent: "space-between" }}
          >
            <View>
              <Text category="h2">{ctx?.state?.user?.name}</Text>
              <Text appearance="hint" category="s1">
                {ctx?.state?.user?.email}
              </Text>
            </View>
            <Button
              disabled={loading}
              size="medium"
              status="basic"
              appearance="ghost"
              icon={style => <Icon name="log-out-outline" {...style} />}
              onPress={logout}
            />
          </MainContainer>
          <Stats>
            <Stat hint="Receipts" value={ctx?.state?.user?.receipts?.length} />
            <Stat hint="Orders" value={ctx?.state?.user?.orders?.length} />
            <Stat hint="Points" value={ctx?.state?.user?.points} />
          </Stats>
          <EditButton
            icon={style => <Icon name="edit-outline" {...style} />}
            // onPress={onFollowButtonPress}>
          >
            EDIT PROFILE
          </EditButton>
        </ProfileContainer>
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
  flex-direction: row;
  margin-horizontal: -16px;
  padding-horizontal: 16px;
  padding-top: 16px;
  margin-bottom: 8px;
`

const ModalContainer = styled(Layout)`
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  padding: 16px;
`

const ProfileContainer = styled(View)`
  flex: 1;
  margin-horizontal: 8px;
`

const MainContainer = styled(View)`
  flex-direction: row;
  justify-content: space-between;
`

const Stats = styled(View)`
  flex-direction: row;
  margin-top: 24px;
`

const Avatar = styled(KAvatar)`
  margin-horizontal: 8px;
`

const EditButton = styled(Button)`
  margin-vertical: 16px;
`

const StatContainer = styled(View)`
  align-items: center;
  flex: 1;
`
