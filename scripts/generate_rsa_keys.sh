#!/bin/bash

echo "
FROM alpine:3
RUN apk add --no-cache openssh openssl && apk add --update coreutils
" | docker build -t localhost/generate_rsa_keys -

#  Adding `> /dev/null 2>&1` to `ssh-keygen` and `openssl` simply mutes their terminal outputs.
docker run -it --rm localhost/generate_rsa_keys /bin/sh -c '
    echo -e "🚧 \e[32m\e[1mYour keys will be generated using RSA-256.\e[0m\n"

    ssh-keygen -t rsa -b 2048 -m PEM -P "" -f rsa256.key > /dev/null 2>&1

    openssl rsa -in rsa256.key -pubout -outform PEM -out rsa256.key.pub > /dev/null 2>&1

    echo -e "🚀 \e[32m\e[1mYou can find your keys below (Base64 encoding scheme). Copy only the parts in teal color.\e[0m\n"

    echo -e "🤝 \e[32m\e[1mPublic Key (share it publicly or with trusted entities at your discretion):\e[0m\n"
    echo -e "\e[36m$(base64 -i -w 0 rsa256.key.pub)\e[0m\n"

    echo -e "🚨 \e[32m\e[1mPrivate Key (share with nobody, keep secret and securely stored):\e[0m\n"
    echo -e "\e[36m$(base64 -i -w 0 rsa256.key)\e[0m\n"

    echo -e "🐒 \e[32m\e[1mGo ship some code!\e[0m\n"
'
