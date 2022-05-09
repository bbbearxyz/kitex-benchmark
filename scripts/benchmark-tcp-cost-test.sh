#!/bin/bash
set -e
CURDIR=$(cd $(dirname $0); pwd)

echo "Checking whether the environment meets the requirements ..."
source $CURDIR/env-tcp-cost-test.sh
echo "Check finished."

repo=("grpc" "kitex" "tcp-cost-test")
ports=(8000 8001 8002)

echo "Building grpc services by exec build.sh..."
source $CURDIR/build.sh
echo "Build finished."
# benchmark
for ((i = 0; i < ${#repo[@]}; i++)); do
  rp=${repo[i]}
  addr="127.0.0.1:${ports[i]}"
  # server start
  echo "Starting server [$rp], if failed please check [output/log/nohup.log] for detail."
  nohup $cmd_server $output_dir/bin/${rp}_reciever >> $output_dir/log/nohup.log 2>&1 &
  sleep 1
  echo "Server [$rp] running with [$cmd_server]"

// 参数有问题

  # run client
  echo "Client [$rp] running with [$cmd_client]"
  $cmd_client $output_dir/bin/${rp}_bencher -addr="$addr" -n=$n -isStream=1 | $tee_cmd

  # stop server
  pid=$(ps -ef | grep ${rp}_reciever | grep -v grep | awk '{print $2}')
  disown $pid
  kill -9 $pid
  echo "Server [$rp] stopped, pid [$pid]."
  sleep 1
done

finish_cmd