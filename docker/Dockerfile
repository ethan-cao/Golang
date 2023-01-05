FROM node:latest

ENV MONGO_DB_USERNAME=admin \
    MONGO_DB_PWD=pwd

RUN mkdir -p /home/app

COPY ./hello.js /home/app

CMD ["node", "/home/app/hello.js"]

# keep container running,  
# this is for dev purposes only, you shouldn't need to keep a container alive unless it's running a process
ENTRYPOINT ["tail", "-f", "/dev/null"]