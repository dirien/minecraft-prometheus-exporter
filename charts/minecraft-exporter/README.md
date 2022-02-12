# Minecraft Exporter for Prometheus

![minecraft-exporter](https://dirien.github.io/minecraft-prometheus-exporter/img/minecraft-exporter.jpg)

![Version: 0.4.0](https://img.shields.io/badge/Version-0.4.0-informational?style=for-the-badge)
![Type: application](https://img.shields.io/badge/Type-application-informational?style=for-the-badge)
![AppVersion: 0.12.0](https://img.shields.io/badge/AppVersion-0.12.0-informational?style=for-the-badge)

![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=Prometheus&logoColor=white)
![Minecraft](https://img.shields.io/badge/Minecraft-62B47A?style=for-the-badge&logo=Minecraft&logoColor=white)
![Docker](https://img.shields.io/badge/docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Alpine Linux 3.15.0](https://img.shields.io/badge/alpine_linux_3.15.0-0D597F?style=for-the-badge&logo=alpine-linux&logoColor=white)
![Helm](https://img.shields.io/badge/helm-0F1689?style=for-the-badge&logo=helm&logoColor=white)

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/minecraft-exporter&style=for-the-badge)](https://artifacthub.io/packages/search?repo=minecraft-exporter)

## Description

A Helm chart for prometheus minecraft exporter

## Usage

Please add the minecraft-exporter repository before installing any chart provided by this repository:

```bash
helm repo add minecraft-exporter https://dirien.github.io/minecraft-prometheus-exporter
helm repo update
```

### Installing the Chart

To install the chart with the release name minecraft-exporter run:

```bash
helm install minecraft-exporter minecraft-exporter/minecraft-exporter --version 0.4.0
```

After a few seconds, minecraft-exporter should be running.

To install the chart in a specific namespace use following commands:

```bash
kubectl create ns minecraft-exporter
helm install minecraft-exporter minecraft-exporter/minecraft-exporter --namespace minecraft-exporte --version 0.4.0
```

> **Tip**: List all releases using `helm list`, a release is a name used to track a specific deployment

## Values

| Key                    | Type   | Default                               | Description                                                                                                           |
|------------------------|--------|---------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| affinity               | object | `{}`                                  | Set the affinity for the pod.                                                                                         |
| extraContainers        | list   | `[]`                                  | Extra containers for the exporter pod                                                                                 |
| extraEnv               | list   | `[]`                                  | Extra environment variables                                                                                           |
| extraVolumeMounts      | string | `nil`                                 | Extra Volume Mounts for the exporter container                                                                        |
| extraVolumes           | string | `nil`                                 | Extra Volumes for the pod                                                                                             |
| fullnameOverride       | string | `""`                                  | String to override the default generated fullname                                                                     |
| image.pullPolicy       | string | `"IfNotPresent"`                      | The docker image pull policy                                                                                          |
| image.repository       | string | `"ghcr.io/dirien/minecraft-exporter"` | The docker image repository to use                                                                                    |
| image.tag              | string | `""`                                  | The docker image tag to use @default Chart version                                                                    |
| ingress.annotations    | object | `{}`                                  | Additional annotations                                                                                                |
| ingress.enabled        | bool   | `false`                               | Specifies what type of Ingress should be created                                                                      |
| ingress.hosts          | list   | `["chart-example.local"]`             | Ingress host                                                                                                          |
| ingress.path           | string | `"/"`                                 | Ingress path                                                                                                          |
| ingress.tls            | list   | `[]`                                  | Ingress tls                                                                                                           |
| initContainers         | list   | `[]`                                  | Init Containers for Exporter Pod                                                                                      |
| nameOverride           | string | `""`                                  | String to override the default generated name                                                                         |
| nodeSelector           | object | `{}`                                  | Set the node selector for the pod.                                                                                    |
| options                | list   | `[]`                                  | Flags - for a list visit https://github.com/dirien/minecraft-prometheus-exporter#usage-                               |
| podAnnotations         | object | `{}`                                  | Annotations for the pods                                                                                              |
| rbac.create            | bool   | `true`                                | Specifies whether RBAC resources should be created                                                                    |
| rbac.pspEnabled        | bool   | `true`                                | Specifies to enable pod security policy                                                                               |
| rbac.pspUseAppArmor    | bool   | `true`                                | Specifies to enable use AppArmor                                                                                      |
| replicaCount           | int    | `1`                                   | (int) Numbers of replicas                                                                                             |
| resources              | object | `{}`                                  | Set the resources requests and limits                                                                                 |
| service.annotations    | object | `{}`                                  | Additional annotations                                                                                                |
| service.port           | int    | `9150`                                | Default Service port                                                                                                  |
| service.type           | string | `"ClusterIP"`                         | Specifies what type of Service should be created                                                                      |
| serviceAccount.create  | bool   | `true`                                | Specifies whether a ServiceAccount should be created                                                                  |
| serviceAccount.name    | string | `nil`                                 | The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template |
| serviceMonitor.enabled | bool   | `false`                               | When set true then use a ServiceMonitor to configure scraping                                                         |
| tolerations            | list   | `[]`                                  | Set the tolerations for the pod.                                                                                      |

**Homepage:** <https://github.com/dirien/minecraft-prometheus-exporter/>

## Source Code

* <https://github.com/dirien/minecraft-prometheus-exporter/>

## Maintainers

| Name   | Email               | Url |
|--------|---------------------|-----|
| dirien | engin.diri@ediri.de |     |

## Legal Disclaimer 👮

This project is not affiliated with Mojang Studios, XBox Game Studios, Double Eleven or the Minecraft brand.

"Minecraft" is a trademark of Mojang Synergies AB.

Other trademarks referenced herein are property of their respective owners.