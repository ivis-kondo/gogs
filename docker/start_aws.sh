#!/bin/sh
set -x
set -e

# 環境変数置換
sed -i -e "s/\[DG_DB_HOST\]/$DG_DB_HOST/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DG_DB_APPNAME\]/$DG_DB_APPNAME/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DG_DB_APPUSER\]/$DG_DB_APPUSER/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DG_DB_APPPW\]/$DG_DB_APPPW/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DOMAIN\]/$DOMAIN/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DG_GIT_TOKEN\]/$DG_GIT_TOKEN/g" /data/gogs/custom/conf/app.ini
sed -i -e "s/\[DG_TEMPLATE_BRANCH\]/$DG_TEMPLATE_BRANCH/g" /data/gogs/custom/conf/app.ini

/app/gogs/gogs web
