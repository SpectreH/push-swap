#!/bin/bash
if test "$#" -ne 2; then
  echo "Illegal number of parameters"
  exit 0
fi

argument=$1
mode=$2
declare -i counter=$argument
numbers=""
numbersArr=()
for ((i = 1; i <= $1; i++)); do
  declare -i numberToAppend="$((RANDOM % 1000))"
  repeats=no
  while [[ "$repeats" = "no" ]]; do
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

if [ "$mode" = 0 ]; then
  ./push-swap "$numbers"
  echo -n "Instructions counter: "
  ./push-swap "$numbers" | wc -l
else
  ./push-swap "$numbers" | ./checker "$numbers"
fi
