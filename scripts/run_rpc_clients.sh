#!/bin/bash
set -e
CURDIR=$(cd $(dirname $0); pwd)

source $CURDIR/env.sh

repo=("grpc" "kitex" "tchannel")
ports=(8000 8002 8004)

# 默认为127.0.0.1
ip=${IP:-"10.222.1.129"}

# build
source $CURDIR/build.sh

# benchmark
for b in ${body[@]}; do
  for c in ${concurrent[@]}; do
    for d in ${field[@]}; do
      for e in ${latency[@]}; do
        for ((i = 0; i < ${#repo[@]}; i++)); do
          rp=${repo[i]}
          addr="${ip}:${ports[i]}"

          # run client
          echo "Client [$rp] running with [$taskset_client]"
          $cmd_client $output_dir/bin/${rp}_bencher -addr="$addr" -b=$b -c=$c -n=$n --sleep=$sleep -field=$d -latency=$e -isStream=0 -isTCPCostTest=0 | $tee_cmd
        done
      done
    done
  done
done

finish_cmd
