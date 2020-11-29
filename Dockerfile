FROM alpine:latest

WORKDIR /opt/marketauth

ADD auth_srv /opt/marketauth/auth_srv
RUN chmod a+x /opt/marketauth/auth_srv
ADD conf /opt/marketauth/conf
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE 18185
EXPOSE 18186

ENTRYPOINT ["./auth_srv"]