#!/bin/bash
set -o nounset -o pipefail
command -v shellcheck >/dev/null && shellcheck "$0"

OWNER="babylonchain"
REPO="babylon-contract"
CONTRACTS="babylon_contract btc_staking"
OUTPUT_FOLDER="$(dirname "$0")"

[ -z "$GITHUB_API_TOKEN" ] && echo "Error: Please define GITHUB_API_TOKEN variable." >&2 && exit 1

[ $# -ne 1 ] && echo "Usage: $0 <version>" && exit 1
type curl >&2

TAG="$1"

GH_API="https://api.github.com"
GH_REPO="$GH_API/repos/$OWNER/$REPO"
GH_TAGS="$GH_REPO/releases/tags/$TAG"
AUTH="Authorization: token $GITHUB_API_TOKEN"

# Validate token
curl -o /dev/null -sH "$AUTH" $GH_REPO || { echo "Error: Invalid repo, token or network issue!";  exit 1; }

# Read asset tags
RESPONSE=$(curl -sH "$AUTH" "$GH_TAGS")

for CONTRACT in $CONTRACTS
do
  # Get id of the contract
  ID=$(echo "$RESPONSE" | grep -C3 "name.:.\+$CONTRACT.wasm" | grep -w id | cut -f1 -d, | awk '{print $2}')

  [ -z "$ID" ] && echo "Error: Failed to get asset id, response: $RESPONSE" | awk 'length($0)<100' >&2 && exit 1
  GH_ASSET="$GH_REPO/releases/assets/$ID"

  # Download asset file
  echo -n "Downloading $CONTRACT..." >&2
  curl -s -L -H "Authorization: token $GITHUB_API_TOKEN" -H 'Accept: application/octet-stream' "$GH_ASSET" >"$OUTPUT_FOLDER/$CONTRACT.wasm"
  echo "$TAG" >"$OUTPUT_FOLDER/version.txt"
  echo "done." >&2
done
