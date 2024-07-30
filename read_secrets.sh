#!/bin/sh

for var in $(env); do
  if echo "$var" | grep -q '_FILE='; then
    var_name=$(echo "$var" | cut -d'=' -f1 | sed 's/_FILE$//')
    file_path=$(echo "$var" | cut -d'=' -f2)
    if [ -f "$file_path" ]; then
      export "$var_name"="$(cat "$file_path")"
    fi
  fi
done

exec "$@"
