#!bin/bash
helm repo add appvia https://terraform-controller.appvia.io
helm repo update

helm install -n terraform-system terraform-controller appvia/terraform-controller --create-namespace
kubectl -n terraform-system get pods

kubectl -n terraform-system create secret generic aws \
  --from-literal=AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
  --from-literal=AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
  --from-literal=AWS_SESSION_TOKEN=$AWS_SESSION_TOKEN