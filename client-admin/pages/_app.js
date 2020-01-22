import App from 'next/app';
import React from 'react';

const theme = {
  colors: {
    primary: '#0070f3'
  }
};

export default class MyApp extends App {
  render() {
    const { Component, pageProps } = this.props;
    return <Component {...pageProps} />;
  }
}
