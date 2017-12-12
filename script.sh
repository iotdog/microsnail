#!/usr/bin/env sh

## Constants

INIT_SVC="init"

##

show_help()
{
  echo "usage: script.sh [cmd] [arg]"
  echo "cmd list:"
  echo "$INIT_SVC - initialize service, arg is service name, eg. script.sh $INIT_SVC newsvc"
}

init_svc()
{
  local svcname=$1
  echo "initialize service: $svcname"
  if [ -d $svcname ]; then
    echo "error: directory $svcname existed"
  else
    mkdir $svcname
    cp template/svcreadme.md.template $svcname/README.md
    sed -i -e "s/(replacehere4svcname)/$svcname/g" $svcname/README.md
    if [ -f $svcname/README.md-e ]; then
      rm $svcname/README.md-e
    fi
    touch $svcname/main.go

    mkdir $svcname/client
    touch $svcname/client/main.go

    mkdir $svcname/config
    cp template/configs.go.template $svcname/config/configs.go
    cp template/debug.json.template $svcname/config/debug.json
    sed -i -e "s/(replacehere4svcname)/$svcname/g" $svcname/config/debug.json
    if [ -f $svcname/config/debug.json-e ]; then
      rm $svcname/config/debug.json-e
    fi
    cp template/preview.json.template $svcname/config/preview.json
    sed -i -e "s/(replacehere4svcname)/$svcname/g" $svcname/config/preview.json
    if [ -f $svcname/config/preview.json-e ]; then
      rm $svcname/config/preview.json-e
    fi
    cp template/production.json.template $svcname/config/production.json
    sed -i -e "s/(replacehere4svcname)/$svcname/g" $svcname/config/production.json
    if [ -f $svcname/config/production.json-e ]; then
      rm $svcname/config/production.json-e
    fi

    mkdir $svcname/handler
    touch $svcname/handler/handler.go
    mkdir $svcname/proto
    cp template/svc.proto.template $svcname/proto/svc.proto
    sed -i -e "s/(replacehere4svcname)/$svcname/g" $svcname/proto/svc.proto
    if [ -f $svcname/proto/svc.proto-e ]; then
      rm $svcname/proto/svc.proto-e
    fi

    mkdir $svcname/wrapper
    touch $svcname/wrapper/middleware.go
  fi
}

if [ $# -eq 0 ]; then
  show_help
elif [ $1 = $INIT_SVC ]; then
  if [ ! $# -eq 2 ]; then
    echo "usage: script.sh init [svcname]"
  else
    init_svc $2
  fi
else
  echo "unknown"
fi
