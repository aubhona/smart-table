#!/usr/bin/env bash

extract_env_vars() {
    local config_file=$1
    local env_file=$2
    local prefix=$3

    echo "🔹 Generating $env_file from $config_file"

    > $env_file

    while IFS= read -r line; do
        line=$(echo "$line" | tr -d '[:space:]"')

        if [[ -z "$line" || "$line" =~ ^# ]]; then
            continue
        fi

        if [[ "$line" =~ ^[a-zA-Z0-9_-]+:$ ]]; then
            section=$(echo "$line" | tr -d ':')
            continue
        fi

        key=$(echo "$line" | awk -F: '{print $1}')
        value=$(echo "$line" | cut -d':' -f2- | sed 's/^ //g')

        if [[ -z "$key" ]]; then
            continue
        fi

        if [[ -n "$section" ]]; then
            key="${prefix}_${section}_${key}"
        else
            key="${prefix}_${key}"
        fi

        key=$(echo "$key" | sed 's/__/_/g')

        key=$(echo "$key" | tr '[:lower:]' '[:upper:]')

        if [[ -n "$value" ]]; then
            echo "$key=$value" >> $env_file
        fi
    done < "$config_file"
}

extract_env_vars "configs/config.yaml" "configs/config.env" "SMART_TABLE"

echo "✅ config.env files updated!"
