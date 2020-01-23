import Head from "next/head"
import withAuth from "../hocs/withAuth.js"

const Users = () => (
  <>
    <Head>
      <title>Aumo Users</title>
    </Head>
  </>
)

export default withAuth(Users)
