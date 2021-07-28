FROM node:alpine

# app dir
ARG APP_DIR=app
RUN mkdir -p ${APP_DIR}
WORKDIR ${APP_DIR}

# install dependencys
COPY package*.json ./
RUN npm install --production

# copy project files
COPY . .
EXPOSE 8080

# Запуск проекта
CMD ["npm", "start", "--production"]