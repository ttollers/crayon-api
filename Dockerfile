FROM node:carbon

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./

RUN npm install
# If you are building your code for production
# RUN npm install --only=production

# Bundle app source
COPY ./src ./src

COPY temp/translated_final.json /usr/src/book/translated_final.json
COPY temp/native_final.json /usr/src/book/native_final.json

EXPOSE 8080
CMD [ "npm", "start" ]