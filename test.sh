#!/bin/bash
if test "$#" -ne 1; then
    echo "Illegal number of parameters"
    exit 0
fi

argument=$1;
declare -i counter=$argument
numbers=""
numbersArr=()
for ((i=1;i<=$1;i++)); do
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
  if [ "$i" != "$counter" ]; then
  numbers+=" " 
  fi
done

./push_swap "$numbers"
