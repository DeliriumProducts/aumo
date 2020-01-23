FROM node:13-alpine

WORKDIR /usr/app

COPY package*.json ./
COPY yarn.lock ./

RUN yarn

COPY . .

ARG BACKEND_URL

RUN yarn build

EXPOSE 3000

CMD yarn start
