#! /bin/sh

set -a
. "./.env"
set +a

goose -dir "./sql/schema" "$@"