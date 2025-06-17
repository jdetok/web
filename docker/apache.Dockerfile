FROM httpd:2.4

# COPY .conf/apache2.conf /usr/local/apache2/conf/httpd.conf

COPY .conf/jdetoweb.conf /usr/local/apache2/conf/extra/jdetoweb.conf
COPY .conf/jdetoweb-ssl.conf /usr/local/apache2/conf/extra/jdetoweb-ssl.conf

RUN mkdir -p /usr/local/apache2/conf/certs

COPY .certs/cloudflare_origin_ca.pem /usr/local/apache2/conf/certs/cloudflare_origin_ca.pem
COPY .certs/cloudflare-origin.pem /usr/local/apache2/conf/certs/cloudflare-origin.pem
COPY .certs/cloudflare-origin.key /usr/local/apache2/conf/certs/cloudflare-origin.key
