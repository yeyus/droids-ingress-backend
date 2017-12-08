FROM aledbf/nginx-error-server:0.6
MAINTAINER elyeyus@gmail.com

ADD ./www /var/www/html

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
