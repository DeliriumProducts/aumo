---
documentclass: article
title: Aumo - The Digital Receipts of the Future
fontfamily: utopia
author:
  - "Simo Aleksandrov"
  - "Lyubo Lyubchev"
lang: en
papersize: a4
graphics:
  - a4paper
keywords: [gamification, eco, internet of things]
abstract: |
  Tons of paper receipts are produced and then immediately thrown away, for the creation of which are used trees and a human toxic chemical, requiring thousands of decares of forests to be cut down and resulting in spreading of diseases.

  **Aumo** is a mobile application, accompanied by a hardware device and a web server, which aims at removing paper receipts by replacing them with a digital equivalent. Receipt printers of shops and restaurants will be equipped with **Aumo**. Clients will take their digital receipts by approaching their phone (through our mobile application) to **Aumo**, establishing a connection via NFC (Near-Field Communication) technology.
  Incentive for using the digital receipt, as opposed to the paper alternative, will be points which users receive when choosing **Aumo** over the traditional option. Points can be exchanged for bonuses, which can either be discounts or physical items, provided by the shop or restaurant.

  The project is of applicational nature, it is still under development and belongs to the IT field. The idea was conceived by Simo Aleksandrov and was realised by both of the authors.
  \newpage
header-includes: |
  \usepackage{sectsty}
  \sectionfont{\LARGE\underline\bfseries\centering}
  \subsectionfont{\large\bfseries\centering}
  \subsubsectionfont{\normalsize\bfseries\centering}
  \usepackage{wrapfig}
  \usepackage{float}
figPrefix:
  - "Figure."
  - "Figures."
---

\newpage

# Introduction

Some of the materials needed for creating a receipt are:

- BPA (Bisphenol A)
- Wood

The former is a chemical, which is toxic to the human skin and can lead to diseases such as cancer, diabetes type 2, obesity and others
The paper needed to create a receipt requires a lot of trees - statistics show that every year 60 000+ decares of forests are cut
Over 60 000+ decares of forests are cut down every year just for the creation of paper receipts.
Removing the need for both of these materials will help us and our planet.

With this project we replace the paper receipts with a digital alternative. In order to achieve this, we set ourselves some goals:

- Planning our architecture and choosing the correct technologies
- Creating a device, which will act as a middleman between POS Systems and our mobile app
- Developing a backend server
- Desiging a beautiful and easy to use graphical interface
- Mobile app for the clients of restaurants and stores
- Incentive for our users (gamification - win points and rewards)
- Admin panel for adding rewards

# Gallery

\begin{figure}[H]
\includegraphics[width=0.5\textwidth, height=7cm]{images/iphone.png}
\centering
\caption{Mobile App on iPhone X}
\end{figure}

\begin{figure}[H]
\includegraphics[width=0.5\textwidth, height=7cm]{images/render.png}
\centering
\caption{The Device}
\end{figure}

\begin{figure}[H]
\includegraphics[width=0.3\textwidth, height=7cm]{images/AumoLogo.png}
\centering
\caption{Aumo's Logo}
\end{figure}

# Features

We provide the following features:

- Blazing fast
- Transfering information (the digital receipt) through the use of NFC (Near-Field Communication)
- Online Store for coupons/rewards, provided from the corresponding shop/restaurant
- User panel (Mobile app)
- Admin panel for managing the online store
- History of all claimed receipts
- Displaying all daily/weekly/monthly expenses

# How it works

Every POS terminal will be connected with **Aumo** - a small computer (Raspberry Pi), placed inside a 3D printed case.
The device will act as a middleman and intercept the required information from the POS Terminal and based on the user's choice the receipt will either be sent to the printer (if the user decides to get a traditional receipt, for example if they don't have the app or a smartphone) or to their mobile device through the use of an NFC module.
**Aumo** is equipped with an NFC module, which it uses for receiving the receipt from the POS Terminal. The mobile app receives the receipt when it approaches **Aumo**, thanks to the on board NFC chip that majority of smartphones have.

# Мотив (Геймификация)

При всяко използване на **Aumo** клиента получава точки, те могат да бъдат обменени за бонуси, промоции или предметни награди предоставени от търговския обект, като по този начин клиентите биват мотивирани да използват нашето приложение.

Досега всичко е било "на хартия". Този подход с дигитализация на нещо толкова битово ще изисква много усилия, тъй като ние хората обичаме да стоим в комфортната си зона и често не обичаме промяна. Затова ние решихме да вкараме геймификация в **Aumo**, като по този начин клиентите ще бъдат мотивирани да използват дигитална касова бележка пред традиционната - хартиена, както и да посещават по често съответния търговски обект предлагащ услугата. Така усилията, които трябва да бъдат положени, както от страна на клиенти, така и от страна на собствениците на търговските обекти няма да се усещат и ще бъдат забавни.

# Технологии

## Backend

### go

Като език за програмиране използвахме **go**, тъй като е бърз, гъвкав, лесен за писане и разбиране и може да се компилира към всички операционни системи - macOS, Linux, Windows.

### go-chi

Като библиотека за HTTP сървър ползвахме **go-chi**, поради факта че е тънък слой (wrapper) над стандартната библиотека на **go** - **net/http**. Предоставя лесна абстракция за създаване на REST API. Малък пример за сървър:

```go
package main

import (
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
```

### MariaDB (MySQL)

Използвахме **MySQL** (или по-конкретно **MariaDB**) като база от данни, поради широкото ѝ разпространение в индустрията. Лесна е за използване и конфигуриране.

### Redis

За сесии, кеширане за MySQL заявки използвахме **Redis** заедно с **go-redis**.

### Docker

За deployment (публикуване) на нашото приложение, използвахме технологията за контейнери, по-конкретно - **Docker**. Той ни предоставя еднакъв environment, независимо от операционната система, дистробуция или други фактори. Също така ни улеснява scaling (скалиране) на много иснтанции.

### GORM

Като библиотека за свързване и абстракция от **go** към **MySQL**, се спряхме на **GORM**. Много бързо и лесно успяхме да създадем нашите модели.

Само с няколко реда код, ние можем да имаме потребители в нашата база от данни.

```go
type User struct {
	gorm.Model
	Name     string     `json:"name" gorm:"not null"`
	Email    string     `json:"email" gorm:"unique;not null"`
	Password string     `json:"-" gorm:"not null" gob:"-"`
	Avatar   string     `json:"avatar" `
	Points   float64    `json:"points" gorm:"not null"`
	Orders   []ShopItem `json:"orders" gorm:"many2many:user_shop_item;"`
	Receipts []Receipt  `json:"receipts"`
}

func (a *Aumo) CreateUser(u User) (User, error) {
	if err := a.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
```

### Raspberry Pi

За устройството използахме **Raspberry Pi**, поставено в 3D принтиранa кутиика, направена в CAD системата **Solidworks**. Кутията беше принтирана в Русенския Университет, тъй като те разполагат с 3D принтер. (виж фиг. \ref{fig:rpi} и \ref{fig:aumo-pics})

\begin{figure}[H]
\includegraphics[width=0.5\textwidth]{images/rpi.jpg}
\centering
\caption{Raspberry Pi\label{fig:rpi}}
\end{figure}

\begin{figure}[H]
\includegraphics[width=0.5\textwidth]{images/render.png}
\includegraphics[width=0.5\textwidth]{images/irl.jpg}
\centering
\caption{Устройството\label{fig:aumo-pics}}
\end{figure}

## Frontend

### React и React-Native

За да създадем хубав и гъвкав интерфейс, заедно с reusable компоненти използвахме **React** - библиотека създадена от Facebook. За мобилното приложение използвахме по-конкретно **React-Native**, тъй като можем да пишем един код за всички платформи - iOS и Android. (виж фиг. \ref{fig:aumo-android} и фиг. \ref{fig:aumo-ios})

\begin{figure}[H]
\includegraphics[width=0.5\textwidth, height=6cm]{images/main_s9.jpg}
\includegraphics[width=0.5\textwidth, height=6cm]{images/shop_s9.jpg}
\centering
\caption{Мобилното приложение на Android\label{fig:aumo-android}}
\end{figure}

\begin{figure}[H]
\includegraphics[width=0.5\textwidth, height=6cm]{images/main_ios.png}
\includegraphics[width=0.5\textwidth, height=6cm]{images/shop_ios.png}
\centering
\caption{Мобилното приложение на iOS\label{fig:aumo-ios}}
\end{figure}

# Етапи на развитие

## Избор на тема

Двамата автори бяхме в Стара Загора (участвахме на състезание), седнахме да обядваме. Докато Симо беше на опашката, забеляза, че зад касите стояха кофи за боклук преливащи от касови бележки, които така и не влизат в употреба а само се изхвърля на вятъра един природен ресурс. Така решихме да измислим решение, с което можем да сложим край на този проблем възможно най-скоро - роди се идеята за **Aumo**.

## Проучване

При установено проучване от нас, не успяхме да открием подобни решения, действащи в момента на пазара. Така се убедихме, че е време да започнем работа върху бъдещият ни продукт - **Aumo**.

## Избиране на технологии и архитектура

През този етап ни минаха доста идеи относно подхода ни с технологиите, като се спряхме на вече гореспоментати. (Виж сек. [Технологии]).

## Архитектура

1. Информацията за касовия бон е изпратена към **Aumo**
2. В зависимост от избора на клиента, касовия бон ще бъде изпратен или към мобилното приложение или към принтера
3. Касовия бон бива криптиран и добавен към профила на потребителя в съвръра

\begin{figure}[H]
\includegraphics[height=10cm]{images/arch.png}
\centering
\caption{Архитектурата на Aumo\label{fig:aumo-arch}}
\end{figure}

## Изработване

Започнахме работа върху проекта, по време на Русенския хакатон (TeenHack Ruse 2019) провел се в началото на октомври.

- Тогава успяхме да създадем REST API написан на **go** чрез, който извършаме CRUD операции.
- Изградихме базова концепция за нашия графичен интерфейс - **Figma**
- Създадохме нашето мобилно приложение с **React Native**
- Свързахме мобилното приложение със сървъра ни

# Описание на приложението

Сорс кода на проекта може да бъде намерен в GitHub[^1]. След като е изтеглен проекта, той може да бъде конфигуриран много лесно, благодарение на **Docker**.

```sh
docker bulid -t aumo .
docker run aumo
```

Мобилното приложение може да бъде стартирано чрез **npm** или **yarn**

```sh
npm run start
# or
yarn start
```

След като са стартирани, ще може да бъде намерено в **Expo** на вашето мобилно устройство, а сървъра на порт 3000.

[^1]: https://github.com/DeliriumProducts/aumo

# Заключение

**Aumo - дигиталните касови бележки на бъдещeто** ще спаси тонове хартия, като съответно ще бъдат запазени хиляди декари гори и спасяването на потенциално изчезващи живонски видове (тези, които живеят по дърветата). Предполага се, че търговските обекти ще започнат да печелят повече благодарение на геймификацията влючена в **Aumo**.
Хората ще спрат да губят своите касови бележки и ще могат лесно и бързо да си правят отчет за деня или седмицата къде какви пари отиват.
Намалява се рискът за различни заболявания, като например рак, и диабет, тъй като BPA вече не е фактор в този вид бележки.

Въпреки предизвикателността на поставените от нас задачи, ние успяхме да преодолеем почти всички. Подбрахме правилната за нуждите ни технология, която да може бързо да се справя с поставените от нас задачи.

Приложението е все още в процес на разработка, като към момента съществува само базов прототип, но се очаква до края на декември 2019 год. то да бъде почти завършено.

# Technologies used

|         Client         |   Server    |     Miscellaneous     |
| :--------------------: | :---------: | :-------------------: |
|         babel          |     go      |        Testing        |
|      react-native      |   go-chi    |          NFC          |
|        next.js         | upper.io/db |        Caching        |
|        react.js        |             |    MySQL / MariaDB    |
|          JSX           |             |         Redis         |
|       CSS-in-JS        |             |         HTTP          |
|        Webpack         |             |         REST          |
|          PWA           |             |     Raspberry Pi      |
|          SPA           |             |       Sessions        |
| React Native UI Kitten |             |        Cookies        |
|       Ant design       |             |        Docker         |
|   styled-components    |             | Git / Version Control |
|                        |             | Continous Integration |
|                        |             |  Domain Name System   |
|                        |             | Linux / Shell scripts |
|                        |             |         Nginx         |
|                        |             |          SSH          |
|                        |             |   TLS / SSL / HTTPS   |
|                        |             |     Docker Swarm      |

# Бъдеще и развитие

Както вече споменахме, проектът е все още в процес на разработка и доразвитие, съответно това означава, че има много потенциал за нови реализации. Някой от
тях, които трябва да осъществим са:

- Осъществяване на връзка между POS компютъра и **Aumo**
- Интегриране на NFC модула
- Достигане до потенциални клиенти.
- Административен панел за добавяне на продукти

С две думи можем да кажем, че според нас проектът има огромен потенциал да се развие.
