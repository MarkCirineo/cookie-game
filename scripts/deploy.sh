#! /bin/bash

# default ENV is dev
env=dev

while test $# -gt 0; do
    case "$1" in
        -env)
            shift
            if test $# -gt 0; then
                env=$1
            fi
            # shift
            ;;
        *)
        break
        ;;
    esac
done

cd ../../cookie-game
source .env
go build -o cmd/cookie-game/cookie-game.exe cmd/cookie-game/main.go
cmd/cookie-game/cookie-game.exe -env $env &