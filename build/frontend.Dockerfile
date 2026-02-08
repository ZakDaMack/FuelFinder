# build stage
FROM node:lts-alpine AS web-build
WORKDIR /app
COPY web/package*.json ./
RUN npm config set registry https://registry.npmjs.org/
RUN npm ci
COPY web .
ENV VITE_APP_API_URL="/api/"
RUN npm run build

# production stage
# this ensures that only the relevant files are copied over
FROM nginx:stable-alpine AS fuelfinder-web
COPY --from=web-build /app/dist /usr/share/nginx/html
EXPOSE 80
HEALTHCHECK CMD curl --fail http://localhost || exit 1
CMD ["nginx", "-g", "daemon off;"]
