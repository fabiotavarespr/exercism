#!/usr/bin/env bash

main () {

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