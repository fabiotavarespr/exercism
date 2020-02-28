#!/usr/bin/env bash

function usage() {
	echo "Usage: raindrops <number>" >&2
	exit 1
}

main () {
    #Check param has only one
    if [ ! "$#" -eq 1 ]; then
        usage
    fi

    #Check param is an number
    if ! [[ $1 =~ ^[0-9]+$ ]]; then
        usage
    fi

    rainSound=""
    if (( $1 % 3 == 0 ))
    then
        rainSound="Pling"
    fi

    if (( $1 % 5 == 0 ))
    then
        rainSound+="Plang"
    fi

    if (( $1 % 7 == 0 ))
    then
        rainSound+="Plong"
    fi

    if [[ -z $rainSound ]]
    then
        echo $1
    else
        echo $rainSound
    fi
}

main "$@"