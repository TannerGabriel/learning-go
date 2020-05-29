FROM node:lts-alpine

RUN npm install -g http-server

WORKDIR /app

COPY package*.json ./

RUN npm install

ARG VUE_APP_API_URL
ENV VUE_APP_API_URL $VUE_APP_API_URL

COPY . .

RUN npm run build

EXPOSE 8080
CMD [ "http-server", "dist" ]