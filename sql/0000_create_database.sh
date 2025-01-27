#!/bin/bash

cat *.sql > ./result/database.sql
mysql -u root -p jardinpc < ./result/database.sql
rm ./result/database.sql

