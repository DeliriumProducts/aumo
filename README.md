<h3 align="center"><img src="https://i.imgur.com/VH2cLGq.png" alt="logo"></h3>
<p align="center">📜 The digital receipts of the future.</p>

[![Github Actions Widget]][github actions] [![GoReport Widget]][goreport] [![GoDoc Widget]][godoc]

Tons of paper receipts are produced and then immediately thrown away, for the creation of which are used trees and a human toxic chemical, requiring thousands of decares of forests to be cut down and resulting in spreading of diseases.

**Aumo** is a mobile application, accompanied by a hardware device and a web server, which aims at removing paper receipts by replacing them with a digital equivalent. Receipt printers of shops and restaurants will be equipped with **Aumo**. Clients will take their digital receipts by approaching their phone (through our mobile application) to **Aumo**, establishing a connection via NFC (Near-Field Communication) technology.

Incentive for using the digital receipt, as opposed to the paper alternative, will be points which users receive when choosing **Aumo** over the traditional option. Points can be exchanged for bonuses, which can either be discounts or physical items, provided by the shop or restaurant.

[Preview here!](https://expo.io/@deliriumproducts/aumo)

[goreport widget]: https://goreportcard.com/badge/github.com/deliriumproducts/aumo
[goreport]: https://goreportcard.com/report/github.com/deliriumproducts/aumo
[github actions widget]: https://github.com/tsoding/kgbotka/workflows/CI/badge.svg
[github actions]: https://github.com/deliriumproducts/aumo/actions
[godoc]: https://godoc.org/github.com/deliriumproducts/aumo
[godoc widget]: https://godoc.org/github.com/deliriumproducts/aumo?status.svg

## Getting started

Requirements:

### Backend

- go
- mysql
- redis

### Admin panel

- node.js

### Mobile

- android sdk
- jdk
- node.js

In order to build a release apk, you need to generate a keystore using:

```console
$ keytool -genkeypair -v -keystore aumo.keystore -alias aumo -keyalg RSA -keysize 2048 -validity 10000
```

... and then move the generated file in client/android/keystores. Next you need to make a file called `release.keystore.properties` in the same directory. It should contain:

```bash
key.store=aumo.keystore
key.alias=aumo
key.store.password=***
key.alias.password=***
```

The final apk can be installed using

```console
$ yarn android --variant=release
```
