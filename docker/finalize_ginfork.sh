#!/bin/sh
# Finalize the build

set -x
set -e

# Create git user for Gogs
addgroup -S git
adduser --ingroup git --no-create-home --disabled-password --gecos 'Gogs Git User' git --home /data/git -shell /bin/bash && usermod --password '*' git && passwd --unlock git
echo "export GOGS_CUSTOM=${GOGS_CUSTOM}" >> /etc/profile

echo "test_cache"

# Final cleaning
rm -rf /app/gogs/build
rm /app/gogs/docker/finalize.sh
rm /app/gogs/docker/nsswitch.conf
rm /app/gogs/docker/README.md
