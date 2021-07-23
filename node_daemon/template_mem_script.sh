#!/bin/sh
dump_dir="./mem_pods_dump/"
pod_uid="%s"
log_file="mem_logs"
max_usage_file="memory.kmem.max_usage_in_bytes"

mkdir ${dump_dir}
chmod 0777 ${dump_dir}
cp -r /sys/fs/cgroup/memory/kubepods.slice/kubepods-pod${pod_uid}.slice ${dump_dir}
cd ${dump_dir}
cat ./kubepods-pod${pod_uid}.slice/${max_usage_file} >> ./${log_file}