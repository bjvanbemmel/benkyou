FROM node:20-slim AS build
WORKDIR /build
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build
RUN npm prune

FROM nginx:stable-alpine AS release
COPY --from=build /build/dist /usr/share/nginx/html
COPY ./default.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]

FROM node:20-slim AS debug
WORKDIR /web
COPY . /web
EXPOSE 3000
RUN npm install
CMD [ "npm", "run", "dev" ]
