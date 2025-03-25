#!/bin/sh

repo="$1"
tag="$2"

if [[ "$repo" == "jfrog" ]]; then
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

  docker push inditex-docker.jfrog.io/production/itxapps/clr-argo:$tag
  docker push inditex-docker.jfrog.io/production/itxapps/clr-workflow-controller:$tag
  docker push inditex-docker.jfrog.io/production/itxapps/clr-argoexec:$tag
else
  docker build \
      --output "type=docker" \
      --platform="linux/amd64" \
      --target argocli \
      --provenance=false \
      --tag axinregistry1.central.inditex.grp/itxapps/clr-argo:$tag .
  docker build \
      --output "type=docker" \
      --platform="linux/amd64" \
      --target workflow-controller \
      --provenance=false \
      --tag axinregistry1.central.inditex.grp/itxapps/clr-workflow-controller:$tag  .
  docker build \
        --output "type=docker" \
        --platform="linux/amd64" \
        --target argoexec \
        --provenance=false \
        --tag axinregistry1.central.inditex.grp/itxapps/clr-argoexec:$tag  .
  
  docker push axinregistry1.central.inditex.grp/itxapps/clr-argo:$tag
  docker push axinregistry1.central.inditex.grp/itxapps/clr-workflow-controller:$tag
  docker push axinregistry1.central.inditex.grp/itxapps/clr-argoexec:$tag
fi