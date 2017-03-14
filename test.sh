#!/bin/env bash


for path in "." ".." "notexist"
do
    actually_file_number=`find $path -type f| wc -l`
    go run main.go -root $path 2>&1 >/dev/null
    if [[ $? == 0 ]]
    then
        really_file_number=`cat sha1result |wc -l`
        if [[ $actually_file_number -eq $really_file_number ]]
        then
            echo "Path: ${path} -> passed test..."
        else
            echo "Path: ${path} -> failed test... -> actually number is ${actually_file_number} : really number is ${really_file_number}"
        fi
        >sha1result
    else
      echo "Path: ${path} -> failed test..."
    fi
done


for pattern in ".git" "g." "g+"
do
    actually_file_number=`find . -type f|grep -v -E "${pattern}"| wc -l`
    go run main.go -root . -filter $pattern 2>&1 >/dev/null
    if [[ $? == 0 ]]
    then
        really_file_number=`cat sha1result | wc -l`
        if [[ $actually_file_number -eq $really_file_number ]]
        then
            echo "Pattern: ${pattern} -> passed test..."
        else
            echo "Pattern: ${pattern} -> failed test... -> actually number is ${actually_file_number} : really number is ${really_file_number}"
        fi
        >sha1result
    else
      echo "Pattern: ${pattern} -> failed test..."
    fi
done

for output in "/tmp/output1" "/tmp/output2"
do
  actually_file_number=`find . -type f| wc -l`
  go run main.go -root . -output $output 2>&1 >/dev/null
  if [[ $? == 0 ]]
  then
      really_file_number=`cat ${output} | wc -l`
      if [[ $actually_file_number -eq $really_file_number ]]
      then
          echo "Output: ${output} -> passed test..."
      else
          echo "Output: ${output} -> failed test... -> actually number is ${actually_file_number} : really number is ${really_file_number}"
      fi
  else
    echo "Output: ${output} -> failed test..."
  fi
done
