---
documentclass: article
title: Aumo - дигиталните касови бележки на бъдещтето
author:
  - Симо Александров
  - Любо Любчев
lang: bg
papersize: a4
keywords: [gamification, eco, internet of things]
abstract: |
  Тонове касови бележки биват създадени и веднага изхвърелни, като за изработката им се използва **BPA (Bisphenol A)**, химикал вреден за човека. Заедно за изработката на тази хартия е също нужна дървесна маса, което означава, че хиляди декари гори биват отсичани годишно.

  **Aumo** е мобилно приложение, придружено с хардуерно устройство и уеб съврър, което цели да премахне хартиените касови бележки, като ги замести с дигитални. Касовите апарати на магазини и заведения ще бъдат оборудвани с **Aumo**. Клиентите ще да получат техните дигитални касови бележки при допира на тяхното мобилно устройство (през мобилното ни приложение) с **Aumo** чрез NFC (Near-Field Communication) технология.
  За мотив да се използва дигиталната касова бележека пред хартиения еквавилент, потребителите ще бъдат възнаграждавани с точки, всеки път когато клиентът предпочете **Aumo** пред традиционнтата касова бележка. Тези точки могат да бъдат използвани за бонуси под формата на намаления или материални награди осигурени от търговския обект.

  Проектът е с приложен характер, все още е в процес на разработка и е от сферата по информатика и информационни технологии. Идеята е измислена от Симо Александров, а е реализирана от двамата автори.

  Tons of paper receipts are produced and then immediately thrown, for the creation of which is used **BPA (Bisphenol A)**, a human toxic chemical. Thousands of forest decares need to be cut down, as wood is another main component, required for the creation of paper receipts.

  **Aumo** is a mobile application, accompanied by a hardware device and a web server, which aims at removing paper receipts by replacing them with a digital equivalent. Receipt printers of shops and restaurants will be equiped with **Aumo**. Clients will take their digital receipts by approaching their phone (through our mobile application) to **Aumo**, by establishing a connection via NFC (Near-Field Communication) technology.
  Incentive for using the digital receipt, as opposed to the paper alternative, will be points which users receive when choosing **Aumo** over the traditional receipt. Points can be exchanged for bonuses, which can either be discounts or physical items, provided by the shop or restaurant.

  The project has applicational nature, it is still under development and belongs to the IT field. The idea was conceived by Simo Aleksandrov and was realised by both of the authors.
---

# Увод

За изработката на касови бележки се използват множество ресурси. Някой от които включват:

- BPA (Bisphenol A / Бисфенол А)
- Дърво

Първите от които са токсични за човешката кожа, Бисфенол А, може да доведе до заболявания като рак, захарен диабет тип 2, наднормено тегло и други.
Тонове дървета биват отсичани за създаването на хартията на касовите бележки. Статистики показват че се отисчат 60 000+ декара гори годишно само от "Големите 5 държави".
Премахването на тези ресурси ще се подпомогне така и на хората живущи на нашата планета, така и на самата планета.

С тази разработка целим заменянето на хартиените касови бележки с дигитална алтернатива. За да постигнем тази цел трябва да бъдат решени следните задачи:

- Планиране на архитектура и подбор на правилните технологии
- Създаване на устройство, което ще играе ролята на посредник между касови апарти и мобилното приложение
- Разработка на backend съврър
- Оформяне на красив и лесен за използване графичен интерфейс
- Мобилно приложение за клиентите на заведения или магазини
- Мотив за потребителите (система за награди и точки - Gamification)
- Административен панел за добавяне на награди

# Галерия

![](./images/scrot.jpg){ width=50% }
![](./images/irl.jpg){ width=50% }

# Функции

Приложението ни предоставя следните функции:

- Светкавично бързо
- Изпращане на информация (дигитална касова бележка) чрез NFC (Near-Field Communication)
- Виртуален магазин за промоции/награди предоставени от търговсия обект
- Потребителски панел (Мобилно приложение)
- Административен панел за управление на виртуалния магазин
- История от всички касови бележки
- Създаване на списък с разходите извършени за деня/седмицата/месеца съответно

# Как работи

Между всеки касов апарат и компютъра, свързан с него, ще бъде поставено по едно устройство - **Aumo**.
**Aumo** представлява малък компютър (например Raspberry Pi), поставен в кутийка (изработена например от 3D принтер).
Устройството ще играе роля на посредник и ще приема нужната информация от компютъра и в зависимост от избора на клиента,
касовия бон ще бъде изпратен към принтера или към NFC модул.
Към този компютър е също свързан NFC модул, чрез който мобилно приложение ще получава касовата бележка.

# Мотив (Геймификация)

Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.

# Технологии

Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.

# Етапи на развитие

Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.

# Заключение

Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.

# Бъдеще и развитие

Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.
