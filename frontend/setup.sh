#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"
npm i
npm run build
rm -rf /usr/container_setup/static_sites/marvelfrontend
mv ./dist ./marvelfrontend
mv ./marvelfrontend /usr/container_setup/static_sites/
docker restart nginx_static