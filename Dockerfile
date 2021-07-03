# syntax=docker/dockerfile:1

# Go 
FROM golang:1.16.5-alpine
WORKDIR /server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o /server/build/server ./main

CMD /server/build/server

# build stage - Vue
FROM node:lts-alpine as build-stage
WORKDIR /frontend
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend .
RUN ls
RUN pwd
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /frontend/dist/ /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

