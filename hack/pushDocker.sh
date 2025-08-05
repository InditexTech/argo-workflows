#!/bin/bash

repo="$1"
tag="$2"

if [[ "${repo}" == "jfrog" ]]; then
  echo "USING JFROG IF"
  docker push inditex-docker.jfrog.io/production/itxapps/clr-argo:$tag
  docker push inditex-docker.jfrog.io/production/itxapps/clr-workflow-controller:$tag
  docker push inditex-docker.jfrog.io/production/itxapps/clr-argoexec:$tag
else
  echo "USING ELSE"
  docker push ieec1registry1.ecommerce.inditex.grp/itxapps/clr-argo:$tag
  docker push ieec1registry1.ecommerce.inditex.grp/itxapps/clr-workflow-controller:$tag
  docker push ieec1registry1.ecommerce.inditex.grp/itxapps/clr-argoexec:$tag
fi