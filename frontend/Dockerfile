FROM node:alpine

#install requirements
RUN npm install -g http-server
WORKDIR /app
COPY package*.json ./
RUN npm install

#run appp
COPY . .
RUN npm run build
EXPOSE 8080
CMD [ "http-server", "dist" ]