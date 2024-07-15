#!/bin/bash

args=()
args+=("同步")
args+=("异步——1")
args+=("异步——2")

#echo ${args[*]}

sh batch.sh ${args[*]}
