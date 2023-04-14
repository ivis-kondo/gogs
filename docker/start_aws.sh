#!/bin/sh
set -x
set -e

# 環境変数置換
sed -i -e "s/\[DG_DB_HOST\]/$DG_DB_HOST/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DG_DB_APPNAME\]/$DG_DB_APPNAME/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DG_DB_APPUSER\]/$DG_DB_APPUSER/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DG_DB_APPPW\]/$DG_DB_APPPW/g" /data/gogs/custom/conf/app.ini

/app/gogs/gogs web

