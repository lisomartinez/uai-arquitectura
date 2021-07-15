FROM node:14

ARG AZ_COSMOS_MENU_CONNECTION_STRING
ENV AZ_COSMOS_MENU_CONNECTION_STRING=${AZ_COSMOS_MENU_CONNECTION_STRING}
# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./

# RUN npm install
# If you are building your code for production
RUN npm install --production


# Bundle app source
COPY . .
RUN npm run build

EXPOSE 8082
CMD [ "node", "dist/index.js" ]
