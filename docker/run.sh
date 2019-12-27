#!/usr/bin/env sh
set -eo pipefail

sed -i  "s/__PORT__/${PORT}/g;" /etc/nginx/conf.d/default.conf
sed -i  "s/__DOMAIN__/${DOMAIN}/g;" /etc/nginx/conf.d/default.conf

nginx -t

/usr/local/bin/wechat-token &

if [ -z ${DEBUG+x} ]; then
    echo "nginx started"
    nginx -g "daemon off;"
else
    echo "nginx started with debug"
    nginx-debug -g "daemon off;"
fi

