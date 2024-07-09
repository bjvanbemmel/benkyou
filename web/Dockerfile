FROM node:20-slim AS build
WORKDIR /build
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build
RUN npm prune

FROM node:20-slim AS release
COPY --from=build /build/.output /release/.output
EXPOSE 3000
WORKDIR /release
CMD [ "node", "./.output/server/index.mjs" ]

FROM node:20-slim AS debug
WORKDIR /web
COPY . /web
EXPOSE 3000
RUN npm install
CMD [ "npm", "run", "dev" ]
