FROM node:18-alpine
RUN mkdir -p /app

WORKDIR /app

COPY . .
COPY package*.json ./

RUN npm install 




EXPOSE 3000

CMD [ "npm", "start" ]