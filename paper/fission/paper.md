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
  An incentive for using the digital receipt, as opposed to the paper alternative, will be points that users receive when choosing **Aumo** over the traditional option. Points can be exchanged for bonuses, which can either be discounts or physical items, provided by the shop or restaurant.

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

With this project, we replace the paper receipts with a digital alternative. To achieve this, we set ourselves some goals:

- Planning our architecture and choosing the correct technologies
- Creating a device, which will act as a middleman between POS Systems and our mobile app
- Developing a backend server
- Designing a beautiful and easy to use graphical interface
- Mobile app for the clients of restaurants and stores
- Incentive for our users (gamification - win points and rewards)
- Admin panel for adding rewards

# Gallery

\begin{figure}[H]
\includegraphics[width=0.5\textwidth, height=7cm]{images/aumo_s9.png}
\centering
\caption{Mobile App on Samsung S9}
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
- Transferring information (the digital receipt) through the use of NFC (Near-Field Communication)
- Online Store for coupons/rewards, provided from the corresponding shop/restaurant
- User panel (Mobile app)
- Admin panel for managing the online store
- History of all claimed receipts
- Displaying all daily/weekly/monthly expenses

# How it works

Every POS terminal will be connected with **Aumo** - a small computer (Raspberry Pi), placed inside a 3D printed case.
The device will act as a middleman and intercept the required information from the POS Terminal and based on the user's choice the receipt will either be sent to the printer (if the user decides to get a traditional receipt, for example, if they don't have the app or a smartphone) or to their mobile device through the use of an NFC module.
**Aumo** is equipped with an NFC module, which it uses for receiving the receipt from the POS Terminal. The mobile app receives the receipt when it approaches **Aumo**, thanks to the onboard NFC chip that the majority of smartphones have.

# Incentive (Gamification)

Every time a user chooses **Aumo** over the traditional paper receipt, they are rewarded with bonus points, which can be exchanged for coupons or physical rewards, thus acting as an incentive to use our mobile app.

Digitalizing paper receipts will take a lot of effort, mainly since we, as humans prefer to stay within our comfort zone. We decided to implement an element of gamification into **Aumo**, thus making it both fun and rewarding to use, as well as helping out the planet.

# Technologies

## Backend

### go

We choose **go** for our backend programming language, due to its flexibility, performance, simplicity and cross-compilation.

### go-chi

We used **go-chi** as our HTTP server because it's a thin layer on top of **go**'s standard library - **net/http**. It provides a simple abstraction for building REST APIs.

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

We used **MariaDB** for our database, due to its popularity and support in the industry. It's both easy to use and setup.

### Redis

For sessions and caching **MySQL** queries we used **Redis** together with **go-redis**.

### Docker

For deploying our app, we used **Docker** containers. It provides the same environment, no matter the OS, Linux distribution and other factors. It also makes scaling to several instances a breeze.

### upper/db

As a library for connecting our backend code to our database, we used **upper/db**. Just with a few lines of code, we can save users to our database.

```go
// User represents a user of aumo
type User struct {
	ID         uuid.UUID `json:"id,omitempty" db:"id,omitempty"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password"`
	Avatar     string    `json:"avatar" db:"avatar"`
	Points     float64   `json:"points" db:"points"`
	Role       Role      `json:"role" db:"role"`
	Orders     []Order   `json:"orders" db:"-"`
	Receipts   []Receipt `json:"receipts" db:"-"`
	IsVerified bool      `json:"is_verified" db:"verified"`
	Shops      []Shop    `json:"shops,omitempty" db:"-"`
}

func (u *userStore) Save(us *aumo.User) error {
	var err error
	_, err = tx.Collection("users").Insert(us)
	return err
}
```

### Raspberry Pi

For the hardware device, we used a **Raspberry Pi**, placed inside a 3D printed case, which was made in **Solidworks**. The case was printed in the Ruse University "Angel Kanchev", because they have a 3D printer (fig. \ref{fig:rpi} and \ref{fig:aumo-pics}).

\begin{figure}[H]
\includegraphics[width=0.5\textwidth]{images/rpi.jpg}
\centering
\caption{Raspberry Pi\label{fig:rpi}}
\end{figure}

\begin{figure}[H]
\includegraphics[width=0.5\textwidth]{images/render.png}
\includegraphics[width=0.5\textwidth]{images/irl.jpg}
\centering
\caption{The device\label{fig:aumo-pics}}
\end{figure}

## Frontend

### React & React-Native

To create a beautiful and flexible UI, as well as have reusable components, we used **React** - a company developed from Facebook. For our mobile app, we used **React-Native**, because we can share one codebase and have a mobile app both for iOS and Android (fig. \ref{fig:aumo-mobile})

\begin{figure}[H]
\includegraphics[width=0.5\textwidth, height=6cm]{images/main.png}
\includegraphics[width=0.5\textwidth, height=6cm]{images/shop.png}
\includegraphics[width=0.5\textwidth, height=6cm]{images/profile.png}
\centering
\caption{The mobile app on both platforms\label{fig:aumo-mobile}}
\end{figure}

# Stages of development

## Choosing a topic

While having lunch at a fast-food restaurant, we noticed that the trash can was overflowing with paper receipts. The receipts weren't getting used at all, and therefore only wasting precious resources. We wanted to figure out a solution as soon as possible, thus **Aumo** was born.

## Research

After careful analysis and research, we weren't able to find similar products on the market. This motivated us to create our future product - **Aumo**.

## Designing our architecture and choosing technologies

We went through several iterations of our architecture and technology stack, but eventually settled on the aforementioned ones (Check out section [Technologies]).

## Architecture

1. The information for the receipt is sent to **Aumo**.
2. Based on the user's choice, the receipt gets sent either to the printer or the mobile app.
3. The digital receipts gets added to the user's profile on our servers.

\begin{figure}[H]
\includegraphics[height=10cm]{images/arch.png}
\centering
\caption{Aumo's Architecture\label{fig:aumo-arch}}
\end{figure}

## Development

We began working on our project during the Ruse Hackathon (TeenHack Ruse 2019), which was in Early October of 2019.

During that hackathon, we managed to:

- Create a REST API, written in **go**, which can do the basic CRUD operations.
- Design a mockup wireframe with **Figma**
- Develop a mobile app with **React-Native**
- Connect the mobile app with our server

# Conclusion

**Aumo - The Digital Receipts of the Future** will save tons of paper, thus also saving thousands of decares of forests and saving species that live on the trees. We also hope that restaurants and shops will gain a bigger revenue, thanks to the gamification incentive.
People will stop losing their receipts and will easily be able to keep a daily/weekly/monthly record of their expenses.
It will also reduce the chances of illnesses, such as cancer and diabetes, since BPA is no longer a necessity in the creation of digital receipts.

Despite the difficulties of our goals, we managed to overcome the majority of them. This is thanks to our very careful choice for our tech stack.

The app is still a work in progress and as of now it only exists as a proof of concept prototype.

# Future

As we mentioned, the project is still a work in progress, which means it has a lot of potential for future development.
Some goals we have in mind are:

- Establishing the connection between the POS terminal and **Aumo**
- Integrating the NFC module
- Reaching potential customers

\newpage

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
