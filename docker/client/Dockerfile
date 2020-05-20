FROM node:11.15.0-alpine as build
WORKDIR /app
RUN apk update
RUN apk add git
RUN apk add npm
RUN apk add nginx
RUN git clone https://github.com/pshoultz/boneappletea.git
WORKDIR /app/boneappletea/apps/web-app/client
RUN npm install
RUN npm run build

FROM nginx:1.16.0-alpine
COPY --from=build /app/boneappletea/apps/web-app/client/dist/client-v2 /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
