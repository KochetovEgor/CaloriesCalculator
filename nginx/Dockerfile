FROM nginx

WORKDIR /etc/nginx

COPY ./nginx.conf nginx.conf
RUN rm -rf conf.d
COPY ./conf.d conf.d