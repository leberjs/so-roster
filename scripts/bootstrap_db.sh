#!/bin/sh

if [ -d out ]; then
  echo "Cleaning up old output"
  rm -rf out
fi

mkdir out

echo "Bootstrapping database"

cd out

cmds="
create table athlete (id INTEGER PRIMARY KEY ASC, name TEXT); 
"

echo "$cmds" | sqlite3 ./roster.db

