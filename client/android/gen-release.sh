#!/usr/bin/env bash

cat > keystores/release.keystore.properties << EOF
key.store=aumo.keystore
key.alias=aumo
key.store.password=${KEYSTORE_PASSWORD}
key.alias.password=${ALIAS_PASSWORD}
EOF
