# Scripts and Tools

A collection of useful node and network management scripts.

## Block Height Estimator

The Block Height Estimator script is a Bash script that allows you to estimate the future block height of a Cosmos chain based on a given time duration and block time. It utilizes the current block height from `noriad` status and provides flexibility in specifying the block time in various units (seconds, minutes, hours, or days) as well as an optional node parameter for `noriad`.

### Usage

To estimate the future block height, execute the script with the following command:

`./estimate_block_height.sh <time-duration> [block-time] [node]`
  * `<time-duration>`: The duration of time for which you want to estimate the block height (e.g., 2d for 2 days, 10h for 10 hours, etc.).
  * `[block-time]` (optional): The time it takes for a block to be produced, specified in seconds, minutes, hours, or days. Default value is 5 seconds if not provided.
  * `[node]` (optional): The node parameter for noriad if required. Omit this parameter if not needed.

### Example
To estimate the block height in 2 days with a block time of 5 seconds for Noria's `oasis-3` testnet, run the following command:

`./estimate_block_height.sh 2d 5 https://archive-rpc.noria.nextnet.zone:443/`

If you want to specify a custom node, provide it as the third argument.

`./estimate_block_height.sh 10h 5 tcp://127.0.0.1:26657`

### Dependencies
  * jq: A lightweight and flexible command-line JSON processor.

### Notes
  * The estimation assumes a constant block production rate and doesn't account for potential variations or disruptions in the network.
  * The script retrieves the current block height by invoking noriad status.
  * Ensure you have the necessary permissions to execute the script and access the required resources.
