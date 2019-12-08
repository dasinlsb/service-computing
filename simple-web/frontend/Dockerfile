###############
### BUILD IMAGE
###############
FROM node:alpine as build

# Install dependencies
COPY package.json /app/
COPY yarn.lock /app/
WORKDIR /app
RUN yarn

# Build
COPY . .
RUN yarn build


#################
### RUNTIME IMAGE
#################
FROM nginx:alpine

COPY --from=build /app/build /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
