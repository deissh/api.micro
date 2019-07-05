# GoLang based Microservices with Istio

[![Build Status](https://travis-ci.org/deissh/api.micro.svg?branch=next)](https://travis-ci.org/deissh/api.micro)
[![codecov](https://codecov.io/gh/deissh/api.micro/branch/next/graph/badge.svg)](https://codecov.io/gh/deissh/api.micro)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/067a92d715bf4e3a8eca6d517566e984)](https://www.codacy.com/app/Deissh/api.micro?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=deissh/api.micro&amp;utm_campaign=Badge_Grade)

## Architecture

Very simpled, without load balancer, istio and etc

![](https://i.imgur.com/YaMIVvW.png)

## Deployment

The DEPLOYING.md (ToDo) outlines deploying the stack to Google Kubernetes Engine (GKE)
on the Google Cloud Platform (GCP), with Istio and all associated telemetry
components: Prometheus, Grafana, Zipkin, Jaeger, Service Graph, and Kiali.
This README outlines deploying the Microservices/PostgreSQL/Redis stack locally to Docker Swarm.

***WARNING***

You need create `.env` file and replace `JWT_SECRET`
```bash
mv .env.example .env
```

### Requirements

- Docker
- Helm
- gcloud CLI
- Istio 1.1.x
- Jinja2 (pip install) - optional

### Build images (Optional)

All Docker images, references in the Docker Swarm and Kubernetes resource files,
for the microservices are available on Docker Hub. To build all images yourself,
modify and use these two scripts.

***You must raplace images name in scripts***

```bash
bash ./scripts/1_build_services.sh
bash ./scripts/2_push_images.sh
```

Also you can remove all images use these script.

```bash
bash ./7_push_images.sh
```

### Deploy in GKE Cluster

Before all you need create cluster and install Istio with plugins. Requires Helm to be available from the command-line, locally. Also you need download and install Istio v1.11.xx

Installing Google Cloud SDK (in Debian and Ubuntu). To another platform you can find in [offical docs](https://cloud.google.com/sdk/docs/quickstarts).

```bash
# add the Cloud SDK distribution URI as a package source
echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list

# import the Google Cloud Platform public key
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -

# update the package list and install the Cloud SDK
sudo apt-get update && sudo apt-get install google-cloud-sdk

# and login
gcloud init
```

Now you can create cluster and setup Istio with Grafana, Kiali, Jaeger and Prometheus.

```bash
bash ./scripts/3_create_gke_cluster.sh
export ISTIO_HOME=/some/you/istio/path
bash ./part4_install_istio.sh
```

### Port Forward to Tools

#### Jaeger
```bash
kubectl port-forward -n istio-system $(kubectl get pod -n istio-system -l app=jaeger -o jsonpath='{.items[0].metadata.name}') 16686:16686
```

#### Grafana
```bash
kubectl port-forward -n istio-system $(kubectl get pod -n istio-system -l app=grafana -o jsonpath='{.items[0].metadata.name}') 3000:3000
```

#### Prometheus
```bash
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=prometheus -o jsonpath='{.items[0].metadata.name}') 9090:9090
```

#### Kiali
```bash
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=kiali -o jsonpath='{.items[0].metadata.name}') 20001:20001
```