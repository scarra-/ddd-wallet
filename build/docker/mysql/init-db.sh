#!/bin/bash

mysql -uroot -proot -h mysql < /docker-entrypoint-initdb.d/init.sql