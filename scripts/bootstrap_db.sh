#!/bin/sh

if [ -d out ]; then
  echo "Cleaning up old output"
  rm -rf out
fi

mkdir out

echo "Bootstrapping database"

cd out

# The commands to submit - note that each statement must be ";"-terminated.
cmds="
create table athlete (id INTEGER PRIMARY KEY ASC, name TEXT); 
"

# Pipe the commands to `sqlite3` while also passing the database file path.
echo "$cmds" | sqlite3 ./roster.db

