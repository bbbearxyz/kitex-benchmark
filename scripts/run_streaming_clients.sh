#!/bin/bash
set -e
CURDIR=$(cd $(dirname $0); pwd)

source $CURDIR/env-streaming.sh

repo=("grpc" "kitex" "tcp-streaming" "tchannel")
ports=(8000 8001 8003 8004)

# 默认为127.0.0.1
ip=${IP:-"10.222.1.128"}

# build
source $CURDIR/build.sh

# benchmark
for b in ${body[@]}; do
  for ((i = 0; i < ${#repo[@]}; i++)); do
    rp=${repo[i]}
    addr="${ip}:${ports[i]}"

    # run client
    echo "Client [$rp] running with [$taskset_client]"
    $cmd_client $output_dir/bin/${rp}_bencher -addr="$addr" -b=$b -n=$n -isStream=1 isTCPCostTest=0 | $tee_cmd
  done
done


finish_cmd
