import { Avatar, Button, Icon, Layout, Text } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { StyleSheet, View } from "react-native"
import { Context } from "../../context/context"
import { actions } from "../../context/providers/provider"
export default () => {
  const ctx = React.useContext(Context)
  const [loading, setLoading] = React.useState(false)
  return (
    <>
      <Layout style={styles.header} level="1">
        <Avatar
          style={styles.profileAvatar}
          size="giant"
          source={{ uri: ctx?.state?.user?.avatar }}
        />
        <View style={styles.profileDetailsContainer}>
          <View
            style={{ flexDirection: "row", justifyContent: "space-between" }}
          >
            <View>
              <Text category="h2">{ctx.state.user.name}</Text>
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
              onPress={async () => {
                try {
                  setLoading(true)
                  await aumo.auth.logout()
                  ctx.dispatch({ type: actions.SET_USER, payload: null })
                } catch (error) {
                  console.warn(error)
                } finally {
                  setLoading(false)
                }
              }}
            />
          </View>
          <View style={styles.profileSocialsContainer}>
            <ProfileSocial
              style={styles.profileSocialContainer}
              hint="Receipts"
              value={ctx?.state?.user?.receipts?.length}
            />
            <ProfileSocial
              style={styles.profileSocialContainer}
              hint="Orders"
              value={ctx?.state?.user?.orders?.length}
            />
            <ProfileSocial
              style={styles.profileSocialContainer}
              hint="Points"
              value={ctx?.state?.user?.points}
            />
          </View>
          <Button
            icon={style => <Icon name="edit-outline" {...style} />}
            style={styles.followButton}
            // onPress={onFollowButtonPress}>
          >
            EDIT PROFILE
          </Button>
        </View>
      </Layout>
    </>
  )
}

const ProfileSocial = props => {
  const { style, hint, value, ...viewProps } = props

  return (
    <View {...viewProps} style={[styles.container, style]}>
      <Text category="s2">{value}</Text>
      <Text appearance="hint" category="c2">
        {props.hint}
      </Text>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    alignItems: "center"
  },
  list: {
    flex: 1
  },
  listContent: {
    paddingHorizontal: 8,
    paddingBottom: 8
  },
  header: {
    flexDirection: "row",
    marginHorizontal: -16,
    paddingHorizontal: 16,
    paddingTop: 16,
    marginBottom: 8
  },
  profileAvatar: {
    marginHorizontal: 8
  },
  profileDetailsContainer: {
    flex: 1,
    marginHorizontal: 8
  },
  profileSocialsContainer: {
    flexDirection: "row",
    marginTop: 24
  },
  profileSocialContainer: {
    flex: 1
  },
  followButton: {
    marginVertical: 16
  },
  post: {
    margin: 8
  },
  postHeader: {
    height: 220
  },
  postBody: {
    flexDirection: "row",
    marginHorizontal: -8
  },
  postAuthorContainer: {
    flex: 1,
    justifyContent: "center",
    marginHorizontal: 16
  },
  iconButton: {
    flexDirection: "row-reverse",
    paddingHorizontal: 0
  }
})
