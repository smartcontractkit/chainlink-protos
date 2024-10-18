#!/bin/bash

# Define the plugins and their respective repositories
declare -A plugins
plugins=(
  ["golang"]="https://github.com/asdf-community/asdf-golang.git"
  ["task"]="https://github.com/particledecay/asdf-task.git"
  ["protoc"]="https://github.com/paxosglobal/asdf-protoc.git"
  ["protoc-gen-go-grpc"]="https://github.com/pbr0ck3r/asdf-protoc-gen-go-grpc.git"
  ["protoc-gen-go"]="https://github.com/pbr0ck3r/asdf-protoc-gen-go.git"
  ["nodejs"]="https://github.com/asdf-vm/asdf-nodejs.git"
  ["pnpm"]="https://github.com/jonathanmorley/asdf-pnpm.git"
  ["buf"]="https://github.com/truepay/asdf-buf"
)

# Read the .tool-versions file and install the plugins
while read -r line; do
  plugin=$(echo "$line" | awk '{print $1}')
  version=$(echo "$line" | awk '{print $2}')

  if [[ -n "${plugins[$plugin]}" ]]; then
    echo "Installing $plugin $version..."
    asdf plugin-add "$plugin" "${plugins[$plugin]}"
  else
    echo "No repository URL found for plugin $plugin. Skipping..."
  fi
done < .tool-versions

echo "All plugins installed successfully."
