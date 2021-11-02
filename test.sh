#!/bin/bash

numbers=""
for i in {1..5}; do
  numbers+="$((RANDOM % 100))" 
  if [ "$i" != 5 ]; then
  numbers+=" " 
  fi
done

go run . "$numbers"
