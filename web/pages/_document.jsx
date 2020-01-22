import Document, { Head, Main, NextScript } from "next/document"
import { ServerStyleSheet } from "styled-components"

export default class MyDocument extends Document {
  static async getInitialProps(ctx) {
    const sheet = new ServerStyleSheet()
    const originalRenderPage = ctx.renderPage

    try {
      ctx.renderPage = () =>
        originalRenderPage({
          enhanceApp: App => props => sheet.collectStyles(<App {...props} />)
        })

      const initialProps = await Document.getInitialProps(ctx)
      return {
        ...initialProps,
        styles: (
          <>
            {initialProps.styles}
            {sheet.getStyleElement()}
          </>
        )
      }
    } finally {
      sheet.seal()
    }
  }

  render() {
    return (
      <html lang="en">
        <Head>
          <link
            href="https://fonts.googleapis.com/css?family=Montserrat:100,100italic,200,200italic,300,300italic,regular,italic,500,500italic,600,600italic,700,700italic,800,800italic,900,900italic"
            rel="stylesheet"
          />
          <link
            href="https://fonts.googleapis.com/css?family=Roboto&display=swap"
            rel="stylesheet"
          ></link>
          <meta name="description" content="Manage your aumo here!" />
          <meta
            name="viewport"
            content="width=device-width, initial-scale=1.0"
          />
          <meta name="apple-mobile-web-app-title" content="Aumo Admin" />
          <meta
            name="apple-mobile-web-app-status-bar-style"
            content="default"
          />
          <meta name="apple-mobile-web-app-capable" content="yes" />
          <meta name="mobile-web-app-capable" content="yes" />
          <meta name="apple-mobile-web-app-title" content="Aumo Admin" />
          <meta name="application-name" content="Aumo Admin" />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </html>
    )
  }
}
