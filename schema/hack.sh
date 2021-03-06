#!/bin/sh

function sqldef_dry() {
  ./mysqldef -h localhost -u root goca --dry-run < schema.sql
}

function sqldef_apply() {
  ./mysqldef -h localhost -u root goca  < schema.sql
}

function sqlboiler() {
  ./sqlboiler mysql --no-tests --no-hooks --wipe
}

$1
