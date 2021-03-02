<h3 align="center"><img src="https://i.imgur.com/VH2cLGq.png" alt="logo"></h3>
<p align="center">📜 The digital receipts of the future.</p>

[![Github Actions Widget]][github actions] [![Github Actions Widget2]][github actions2] [![GoReport Widget]][goreport] [![GoDoc Widget]][godoc]

Tons of paper receipts are produced and then immediately thrown away, for the creation of which are used trees and a human toxic chemical, requiring thousands of decares of forests to be cut down and resulting in spreading of diseases.

**Aumo** is a mobile application, accompanied by a hardware device and a web server, which aims at removing paper receipts by replacing them with a digital equivalent. Receipt printers of shops and restaurants will be equipped with **Aumo**. Clients will take their digital receipts by approaching their phone (through our mobile application) to **Aumo**, establishing a connection via NFC (Near-Field Communication) technology.

Incentive for using the digital receipt, as opposed to the paper alternative, will be points which users receive when choosing **Aumo** over the traditional option. Points can be exchanged for bonuses, which can either be discounts or physical items, provided by the shop or restaurant.

[goreport widget]: https://goreportcard.com/badge/github.com/deliriumproducts/aumo
[goreport]: https://goreportcard.com/report/github.com/deliriumproducts/aumo
[github actions widget]: https://github.com/deliriumproducts/aumo/workflows/Backend%20Workflow/badge.svg
[github actions]: https://github.com/deliriumproducts/aumo/actions
[github actions widget2]: https://github.com/deliriumproducts/aumo/workflows/Build%20Android%20App/badge.svg
[github actions2]: https://github.com/deliriumproducts/aumo/actions
[godoc]: https://godoc.org/github.com/deliriumproducts/aumo
[godoc widget]: https://godoc.org/github.com/deliriumproducts/aumo?status.svg

## Getting started

### Paper

- pandoc
- texlive-latex-base
- texlive-latex-extra
- texlive-fonts-recommended 
- texlive-fonts-extra 
- texlive-lang-cyrillic
- cm-super

```console
# Ubuntu
$ sudo apt install pandoc texlive-latex-base texlive-latex-extra texlive-fonts-recommended texlive-fonts-extra texlive-lang-cyrillic cm-super
# Arch
$ yay -S pandoc texlive-core texlive-latexextra texlive-fontsextra texlive-langcyrillic 
```

Run `make` in the corresponding directory

### Backend

Found in the root of the repo

- go
- mysql
- redis

Fill in the .env file in the root of the project
```bash
MYSQL_DATABASE=aumo
MYSQL_USER=root
MYSQL_PASSWORD=
MYSQL_HOST=localhost
MYSQL_PORT=3306

ADDRESS=:8080
COOKIE_SECRET=good-password

REDIS_URL=localhost:6379
REDIS_DATABASE=0
REDIS_URL_TEST=localhost:6379
REDIS_DATABASE_TEST=1

BACKEND_URL=http://localhost:8080/api/v1
FRONTEND_URL=http://localhost:3000
INITIAL_ADMIN_PASSWORD=123456

# Set ENV to PROD if you want to send confirmation emails, otherwise set it to DEV
ENV=PROD
SMTP_HOST=smtp.foo.com
SMTP_USER=foo@bar.com
SMTP_PASS=foobar
```

Create the needed databases:

```mysql
CREATE DATABASE aumo; -- for dev / prod
CREATE DATABASE aumo_test; -- for tests
```

Run the server:

```console
$ go run cmd/aumo/main.go
```

The server should automatically create the initial admin with the `INITIAL_ADMIN_PASSWORD` password and `admin@deliriumproducts.me` email.

If you want to get some data to work with, run:

```console
$ mysql -u root -p aumo < data.sql
```

### Admin panel

Found in the `web/` directory

- node.js
- yarn

Instal deps:

```console
$ yarn
```

Fill in .env:

```bash
BACKEND_URL=localhost:8080/api/v1
```

Start in dev:

```console
$ yarn dev
```

Start in prod:

```console
$ yarn build && yarn start
```

### Mobile

Found in the `client/` directory

- android sdk
- jdk
- node.js
- yarn

Install deps:

```console
$ yarn 
```

Build for android (only needed the first time OR when installing / changing native code)

```console
$ yarn android
```

Build for iOS (only needed the first time OR when installing / changing native code)

```console
$ yarn ios
```

Start normally after having installed the iOS app or APK:

```console
$ yarn start
```

In order to build a release apk, you need to generate a keystore using:

```console
$ keytool -genkeypair -v -keystore aumo.keystore -alias aumo -keyalg RSA -keysize 2048 -validity 10000
```

... and then move the generated file in `client/android/app`.

Next you need to make a file called `release.keystore.properties` in the `client/android/keystores` directory. It should contain:

```bash
key.store=aumo.keystore
key.alias=aumo
key.store.password=***
key.alias.password=***
```

The final apk can be installed using:

```console
$ yarn android --variant=release
```
