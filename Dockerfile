# build stage
FROM node:lts-alpine as frontend-build
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend .
RUN npm run build

# production stage
# this ensures that only the relevant files are copied over
# FROM nginx:stable-alpine as frontend
# COPY --from=frontend-build /app/dist /usr/share/nginx/html
# EXPOSE 80
# HEALTHCHECK CMD curl --fail http://localhost || exit 1
# CMD ["nginx", "-g", "daemon off;"]


FROM node:lts AS scraper-build
WORKDIR /usr/app
COPY scraper .
RUN npm i
RUN npx tsc

FROM node:lts AS scraper
WORKDIR /usr/app
COPY --from=scraper-build /usr/app/dist .
COPY scraper/package*.json ./
COPY data_configs ./configs
RUN npm i --omit=dev
ENV CONFIG_DIR='./configs'
CMD [ "node", "index.js" ]


FROM node:lts AS queryable
COPY queryable .
RUN npm i
RUN npx tsc
CMD [ "node", "dist/index.js" ]