FROM node:22-alpine3.20 AS build
WORKDIR /web
RUN apk update && apk upgrade
RUN apk add git
COPY . /web
RUN npm install
RUN npm run build

FROM node:22-alpine3.20 AS release
WORKDIR /web
COPY --from=build /web/.output ./
EXPOSE 3000
CMD [ "node", ".output/server/index.mjs" ]

FROM node:22-alpine3.20 AS debug
WORKDIR /web
COPY . /web
EXPOSE 3000
RUN npm install
CMD [ "npm", "run", "dev" ]
