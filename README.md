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

#### Build images (Optional)

All Docker images, references in the Docker Swarm and Kubernetes resource files,
for the microservices are available on Docker Hub. To build all images yourself,
modify and use these two scripts.

```bash
bash ./1_build_images.sh
bash ./2_push_images.sh
```

Also you can remove all images use these script.

```bash
bash ./7_push_images.sh
```
