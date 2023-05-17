#!/bin/bash

# Function to convert time duration to seconds
convert_to_seconds() {
  local duration=$1
  local seconds=0

  while [[ $duration =~ ([0-9]+[smhd]) ]]; do
    local value=${BASH_REMATCH[1]}
    local unit=${value: -1}
    local number=${value%?}

    case $unit in
      s) seconds=$((seconds + number));;
      m) seconds=$((seconds + number * 60));;
      h) seconds=$((seconds + number * 60 * 60));;
      d) seconds=$((seconds + number * 24 * 60 * 60));;
      *) ;;
    esac

    duration=${duration#*[smhd]}
  done

  echo $seconds
}

# Check if input parameters are provided
if [ $# -lt 1 ]; then
  echo "Estimates future block height based on time duration and block time."
  echo ""
  echo -e "Usage:\t$0 <time-duration> [block-time] [node]"
  echo -e "Example:\t$0 2d 5s"
  echo -e "Example:\t$0 2d 5s https://archive-rpc.noria.nextnet.zone:443/"
  echo -e "Optional:\tYou may provide the block time and node parameter for noriad. Don\'t forget to add the port for the node endpoint."
  echo -e "Time units:\ts (seconds), m (minutes), h (hours), d (days)"
  exit 1
fi

# Get the time duration from the command-line argument
duration=$1

# Get the block time from the command-line argument (default: 5 seconds)
block_time=${2:-5s}

# Get the node parameter if provided
node=$3

# Get the current block height
current_block_height=$(noriad status --node "$node" | jq -r '.SyncInfo.latest_block_height')

# Convert the time duration and block time to seconds
duration_seconds=$(convert_to_seconds "$duration")
block_time_seconds=$(convert_to_seconds "$block_time")

# Calculate the estimated block height
estimated_block_height=$((current_block_height + duration_seconds / block_time_seconds))

# Print the estimated block height
echo "$estimated_block_height"
