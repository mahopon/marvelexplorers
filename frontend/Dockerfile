FROM node:22-alpine

WORKDIR /marvelfrontend

COPY package*.json ./

RUN npm install

COPY . .

RUN mkdir ./build/marvelexplorers

RUN mv ./build/static ./build/marvelexplorers/static

RUN npm install -g serve

EXPOSE 3000

CMD ["serve", "-s", "build", "-l", "3000"]