#!/bin/bash
set -e
CURDIR=$(cd $(dirname $0); pwd)

source $CURDIR/env-streaming.sh

repo=("grpc" "tcp_client")
ports=(8000 8003)

# 默认为127.0.0.1
ip=${IP:-"10.222.1.129"}

# build
source $CURDIR/build_grpc.sh

# benchmark
for ((i = 0; i < ${#repo[@]}; i++)); do
  rp=${repo[i]}
  addr="${ip}:${ports[i]}"

  # run client
  echo "Client [$rp] running with [$taskset_client]"
  $cmd_client $output_dir/bin/${rp}_bencher -addr="$addr" -n=$n -isStream=1 | $tee_cmd
done


finish_cmd
