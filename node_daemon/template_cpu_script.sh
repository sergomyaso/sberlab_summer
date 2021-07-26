#!/bin/sh
dump_dir="./cpu_pods_dump/"
pod_uid="%s"
log_file="cpu_logs"
max_usage_file="cpu.stat"

mkdir ${dump_dir}
chmod 0777 ${dump_dir}
cp -r /sys/fs/cgroup/cpu/kubepods.slice/kubepods-pod${pod_uid}.slice ${dump_dir}
cd ${dump_dir}
cat ./kubepods-pod${pod_uid}.slice/${max_usage_file} >> ./${log_file}