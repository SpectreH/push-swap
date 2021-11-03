#!/bin/bash

numbers=""
numbersArr=()
for i in {1..5}; do
  declare -i numberToAppend="$((RANDOM % 100))"
  repeats=no
  while [[ "$repeats" = "no" ]] 
  do
    repeats=yes 
    for y in "${numbersArr[@]}"; do
      if [ "$y" = "$numberToAppend" ]; then
      repeats=no 
      numberToAppend=$(($y + 1))
      fi
    done
  done
  numbers+="$numberToAppend"
  numbersArr+=("$numberToAppend")
  if [ "$i" != 5 ]; then
  numbers+=" " 
  fi
done

go run . "$numbers"
