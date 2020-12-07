#!/bin/bash
# example `./test.sh test/test2.md lol.txt`

# validate amount of args
if [[ $# -ne 2 ]]; then
    >&2 echo "Illegal number of parameters: args[source_file_path, target_path]"
    exit 1
fi

# args
source_file_path=$1
output_file_path=$2

# internal
h1amount=0
h2amount=0

# validate filename
if ! [ "${source_file_path: -3}" == ".md" ]; then
  >&2 echo "input file must have markdown format. [source_file_path: $source_file_path]"
  exit 1
fi

while read line
do
  # reading each line
  if [ "$h2amount" -lt 2 ]; then

    if [[ "$line" == "## "* ]]; then
      h2amount=$(($h2amount+1))
    fi

    if [[ ("$h2amount" -lt 2) && ("$h1amount" -eq 1) ]]; then
        echo "$line" # debug to output
        echo "$line" >> "$output_file_path"
    fi

    if [[ $line == "# "* ]]; then
      h1amount=$(($h1amount+1))
    fi

  fi

done < "$source_file_path"
