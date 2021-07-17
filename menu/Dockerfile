FROM node:14

ARG AZ_COSMOS_MENU_CONNECTION_STRING
ENV AZ_COSMOS_MENU_CONNECTION_STRING=${AZ_COSMOS_MENU_CONNECTION_STRING}

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install --production

COPY . .
RUN npm run build

EXPOSE 8082
CMD [ "node", "dist/index.js" ]
