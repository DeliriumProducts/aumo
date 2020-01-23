FROM node:13-alpine

WORKDIR /usr/app

COPY package*.json ./
COPY yarn.lock ./

RUN npm install -g yarn
RUN yarn

COPY . .

RUN yarn build

EXPOSE 3000

CMD yarn start
