FROM node:21-alpine3.19 as angular

WORKDIR /usr/src/app

COPY . .

RUN npm install -g npm@10.5.2

RUN npm install

RUN npm install -g @angular/cli

EXPOSE 4200

CMD ["ng" ,"serve",  "--host", "0.0.0.0"]
