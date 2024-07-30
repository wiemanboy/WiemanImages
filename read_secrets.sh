#!/bin/sh

echo "Reading secrets from files..."
for var in $(env); do
  echo "Reading env variables"
  echo "$var"
  if echo "$var" | grep -q '_FILE='; then
    echo "Env variable contains _FILE"
    echo "$var"
    var_name=$(echo "$var" | cut -d'=' -f1 | sed 's/_FILE$//')
    file_path=$(echo "$var" | cut -d'=' -f2)
    if [ -f "$file_path" ]; then
      echo "Exporting $var_name from $file_path"
      export "$var_name"="$(cat "$file_path")"
    fi
  fi
done

exec "$@"
