# GoLang based Microservices with Istio
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors)

[![Build Status](https://travis-ci.org/nekko-ru/api.svg?branch=next)](https://travis-ci.org/nekko-ru/api)
[![codecov](https://codecov.io/gh/nekko-ru/api/branch/next/graph/badge.svg)](https://codecov.io/gh/nekko-ru/api)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/067a92d715bf4e3a8eca6d517566e984)](https://www.codacy.com/app/nekko-ru/api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=nekko-ru/api&amp;utm_campaign=Badge_Grade)

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

And create volume

```bash
gcloud compute disks create --size 200GB storage-1
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

You need create account

```bash
echo -n 'admin' | base64
# YWRtaW4=
echo -n '1f2d1e2e67df' | base64
# MWYyZDFlMmU2N2Rm
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: kiali
  namespace: istio-system
  labels:
    app: kiali
type: Opaque
data:
  username: YWRtaW4=
  passphrase: MWYyZDFlMmU2N2Rm
EOF
```

```bash
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=kiali -o jsonpath='{.items[0].metadata.name}') 20001:20001
```
## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
<table>
  <tr>
    <td align="center"><a href="https://vk.com/keelvel"><img src="https://avatars0.githubusercontent.com/u/21129524?v=4" width="100px;" alt="deissh"/><br /><sub><b>deissh</b></sub></a><br /><a href="#infra-deissh" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/nekko-ru/api/commits?author=deissh" title="Tests">⚠️</a> <a href="https://github.com/nekko-ru/api/commits?author=deissh" title="Code">💻</a></td>
  </tr>
</table>

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!