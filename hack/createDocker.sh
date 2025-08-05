#!/bin/bash

repo="$1"
tag="$2"

if [[ "${repo}" == "jfrog" ]]; then
  echo "USING JFROG IF"
  docker build \
      --output "type=docker" \
      --platform="linux/amd64" \
      --target argocli \
      --provenance=false \
      --tag inditex-docker.jfrog.io/production/itxapps/clr-argo:$tag .
  docker build \
      --output "type=docker" \
      --platform="linux/amd64" \
      --target workflow-controller \
      --provenance=false \
      --tag inditex-docker.jfrog.io/production/itxapps/clr-workflow-controller:$tag  .
  docker build \
        --output "type=docker" \
        --platform="linux/amd64" \
        --target argoexec \
        --provenance=false \
        --tag inditex-docker.jfrog.io/production/itxapps/clr-argoexec:$tag  .
else
  echo "USING ELSE"
  docker build \
      --output "type=docker" \
      --platform="linux/amd64" \
      --target argocli \
      --provenance=false \
      --tag ieec1registry1.ecommerce.inditex.grp/itxapps/clr-argo:$tag .
  docker build \
      --output "type=docker" \
      --platform="linux/amd64" \
      --target workflow-controller \
      --provenance=false \
      --tag ieec1registry1.ecommerce.inditex.grp/itxapps/clr-workflow-controller:$tag  .
  docker build \
        --output "type=docker" \
        --platform="linux/amd64" \
        --target argoexec \
        --provenance=false \
        --tag ieec1registry1.ecommerce.inditex.grp/itxapps/clr-argoexec:$tag  .
fi