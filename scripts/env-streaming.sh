#!/bin/bash
set -e
CURDIR=$(cd $(dirname $0); pwd)

if ! [ -x "$(command -v taskset)" ]; then
  echo "Error: taskset is not installed." >&2
  exit 1
fi

# cpu binding
nprocs=$(getconf _NPROCESSORS_ONLN)
if [ $nprocs -lt 4 ]; then
  echo "Error: your environment should have at least 4 processors"
  exit 1
elif [ $nprocs -gt 20 ]; then
  nprocs=20
fi
scpu=$((nprocs > 16 ? 4 : nprocs / 4)) # max is 4 cpus
ccpu=$((nprocs-scpu))
scpu_cmd="taskset -c 0-$((scpu-1))"
ccpu_cmd="taskset -c ${scpu}-$((ccpu-1))"
if [ -x "$(command -v numactl)" ]; then
  # use numa affinity
  node0=$(numactl -H | grep "node 0" | head -n 1 | cut -f "4-$((3+scpu))" -d ' ' --output-delimiter ',')
  node1=$(numactl -H | grep "node 1" | head -n 1 | cut -f "4-$((3+ccpu))" -d ' ' --output-delimiter ',')
  scpu_cmd="numactl -C ${node0} -m 0"
  ccpu_cmd="numactl -C ${node1} -m 1"
fi

# GO
GOEXEC=${GOEXEC:-"go"}
GOROOT=$GOROOT

USER=$(whoami)
REPORT=${REPORT:-"$(date +%F-%H-%M)"}
n=10
#body=(64 128 256 512 1024 1024 8192 262144 1048576 10485760 104857600)
#小包
body=(64 128 256 512 1024)

nice_cmd=''
tee_cmd="tee -a output/${REPORT}.log"
function finish_cmd() {
  # to csv report
  ./scripts/reports/to_csv.sh output/"$REPORT.log" > output/"$REPORT".csv
  echo "output reports: output/$REPORT.log, output/$REPORT.csv"
  cat ./output/"$REPORT.csv"
}
if [ "$USER" == "root" ]; then
    nice_cmd='nice -n -20'
fi
cmd_server="${nice_cmd} ${scpu_cmd}"
cmd_client="${nice_cmd} ${ccpu_cmd}"
output_dir=$CURDIR/../output
pb_dir=$CURDIR/../protobuf
grpc_dir=$CURDIR/../protobuf

