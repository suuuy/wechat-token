FROM codingcorp-docker.pkg.coding.net/registry/public/node-yarn-nginx:20181024

COPY docker/nginx/conf.d /etc/nginx/conf.d
COPY docker/nginx/proxy.conf /etc/nginx/
COPY docker/run.sh /usr/local/bin/
COPY dist/wechat-token-linux-amd64  /usr/local/bin/wechat-token

CMD [ "run.sh" ]
