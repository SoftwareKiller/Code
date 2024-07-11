#!/bin/bash

sync_spark_task() {
    echo "spark-sql -f /esm_e_statement_$1" && sleep 2
}

pids=()

async_spark_task() {
    echo "spark-sql -f /esm_e_statement_$1" && sleep 2 &
    pid=$!
    pids+=($pid)
}

# 对参数列表循环获取
for i in $@; do
    if [ $i == "同步" ]; then
        sync_spark_task $i
    else
        async_spark_task $i
    fi 
done


for pid in "${pids[@]}"; do
    wait $pid
    echo "Waiting for process $pid"
done

exit 0
