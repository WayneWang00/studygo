#!/bin/bash

printf "the complete list is %s\n" "$$"
printf "the complete list is %s\n" "$!"
printf "the complete list is %s\n" "$?"
printf "the complete list is %s\n" "$*"
#printf "the complete list is %s\n" "$@"
printf "the complete list is %s\n" "$#"
printf "the complete list is %s\n" "$0"
printf "the complete list is %s\n" "$1"
printf "the complete list is %s\n" "$2"
