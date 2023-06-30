#!/bin/sh
service_name="${1}"

if [ "$service_name" = "" ]; then
  echo "Service name couldn't be empty."
  exit 1
fi

# shellcheck disable=SC2120
generate_doc() {
  go install github.com/swaggo/swag/cmd/swag@v1.8.7
  cd ./"${service_name}/cmd" && swag init --parseDependency=true --output "../docs"
}

generate_doc

echo "${service_name} document generated successfully."
