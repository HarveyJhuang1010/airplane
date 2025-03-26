#!/bin/bash
GITROOT=$(git rev-parse --show-toplevel)
echo $GITROOT

if [ -z "$Mode" ] || [ "$Mode" == "enum" ]; then
  info "deleting old enum files..."
  find "$GITROOT/internal/enum" -type f \( -name '*_enumer.go' -o -name 'zzz_enumer_*.go' \) -delete
  info "generating enum files..."
  if ! command -v enumer &> /dev/null ; then
    info "install enumer"
    go install github.com/dmarkham/enumer@v1.5.9
  fi

  if go generate $GITROOT/internal/enum/... ; then
    info "generate enum files successfully"
  else
    error "failed to generate enum files"
    exit 1
  fi
fi

if [ -z "$Mode" ] || [ "$Mode" == "mock" ]; then
  if ! command -v mockgen &> /dev/null ; then
    info "install mockgen"
    go install github.com/golang/mock/mockgen@v1.6.0
  fi
  info "deleting old mockgen files..."
  find $GITROOT/internal/domain/entities/bo/mock -type f -delete
  info "generate mockgen file"
  go generate ./internal/domain/entities/bo/...
fi

if [ -z "$Mode" ] || [ "$Mode" == "swag" ]; then
  if ! command -v swag &> /dev/null ; then
    info "install swag"
    go install github.com/swaggo/swag/cmd/swag@latest
  fi
  info "generate swag file"
  swag init
fi
