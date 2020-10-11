# Builder
FROM node:latest as builder
WORKDIR /app
RUN npm install @angular/cli -g
COPY package*.json ./
RUN npm install
COPY .  .
RUN npm run build -- --output-path=./dist/out --base-href /

# Prod image
FROM nginx:alpine
RUN rm -rf /usr/share/nginx/html/*
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/dist/out/ /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]