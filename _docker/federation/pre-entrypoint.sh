#!bin/bash

apk --no-cache add python2
chmod a+x /usr/local/bin/entrypoint.py

/usr/local/bin/entrypoint.py