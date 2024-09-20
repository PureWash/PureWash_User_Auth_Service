#!/bin/bash

CURRENT_DIR=$(pwd)
OUTPUT_DIR="${CURRENT_DIR}"  # Output directory within the project directory

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR" || { echo "Failed to create output directory"; exit 1; }

# Iterate over directories in trip_protos
for dir in "${CURRENT_DIR}/PureWash_Protos"/*; do
    # Skip non-directories
    [ -d "$dir" ] || continue

    # Compile .proto files in the directory
    for proto_file in "$dir"/*.proto; do
        echo "Compiling $proto_file..."
        protoc -I="${dir}" -I="${CURRENT_DIR}/PureWash_Protos" -I /usr/local/include --go_out="$OUTPUT_DIR" --go-grpc_out="$OUTPUT_DIR" "$proto_file" || { echo "Failed to compile $proto_file"; exit 1; }
    done
done

echo "Compilation completed successfully."
