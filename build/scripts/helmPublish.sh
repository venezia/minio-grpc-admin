#!/usr/bin/env sh

if [[ "$#" -ne 1 ]]; then
  echo "Usage: `basename $0`<release version>"
  exit 1
fi

helm repo add --username=${HARBOR_ROBOT_LOGIN} --password=${HARBOR_ROBOT_PASSWORD} cnct ${HELM_REGISTRY_AND_USER}
helm dependency update --debug deployments/helm/sos-api-server
helm package --debug deployments/helm/sos-api-server
curl -u ${HARBOR_ROBOT_LOGIN}:${HARBOR_ROBOT_PASSWORD} -F chart=@sos-api-server-${1}.tgz ${HELM_REGISTRY_PUSH_ENPOINT} --show-error --fail