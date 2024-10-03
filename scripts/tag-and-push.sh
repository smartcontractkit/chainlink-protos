#!/bin/bash

# Function to create and push tag
create_and_push_tag() {
  local dir=$1
  local package_json="$dir/package.json"
  local service_name=$(grep '"name"' "$package_json" | sed -E 's/.*"name": *"([^"]+)".*/\1/' | sed 's/@chainlink\///')
  local version=$(grep '"version"' "$package_json" | sed -E 's/.*"version": *"([^"]+)".*/\1/')

  if [ "$service_name" != "null" ] && [ "$version" != "null" ]; then
    local tag="${service_name}/v${version}"
    if git rev-parse "$tag" >/dev/null 2>&1; then
      echo "Tag $tag already exists. Skipping..."
    else
      echo "Creating tag: $tag"
      git tag "$tag"
      echo "Pushing tag: $tag"
      git push origin "$tag"
    fi
  else
    echo "Skipping $dir: Missing name or version in package.json"
  fi
}

# Find all directories containing a package.json
find . -name 'node_modules' -prune -o -path './package.json' -prune -o -name 'package.json' -exec dirname {} \; | while read -r dir; do
  echo "processing $dir"
  create_and_push_tag "$dir"
done
