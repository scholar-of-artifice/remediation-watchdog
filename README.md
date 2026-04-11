# remediation-watchdog

![License](https://img.shields.io/badge/license-MIT-green.svg?labelColor=white)

> **A standalone "Control Plane" service designed to manage the availability of a target application through a closed-loop feedback system.**

## 🧑‍💻 Technologies

<!--technology badges here-->
![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-%23DD0031.svg?logo=redis&logoColor=white)
![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?logo=kubernetes&logoColor=fff)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff)

## ⎈ Orchestration

This project is designed to run in a local `Kubernetes` cluster using `kind` (Kubernetes in Docker). It utilizes native `Kubernetes` primitives to manage lifecycle and reliability.

### Manifests
Found in `/k8s`, defining Deployments, Services, ConfigMaps, etc.

### Namespace
All components run in the `remediation-system` namespace for isolation.

## 🏗️ High Level Architecture

In this example system which demonstrates how to write simple programs to observe and remediate issues in an orchestrated system. 

[TODO: mermaid chart here eventually. not right now until the project is further along.]

The image above is a basic map of the different components. In this system we have a few parts:

### `absurd-iguana`
A basic `Go` application Deployment with a Service endpoint. It produces events to Kafka.

### `dazzling-remora`
A Message Bus using Kafka. For local setup, I am using `KRaft` mode to keep the footprint smaller.

### `eager-marmot`
A `Go` application which consumes events from Kafka and persists data to `Redis`.

### `bashful-yak`
A simple `Redis` instance which stores data. For this application, complicated data models are not the focus.

### `calm-lynx`
This is a `Go` watchdog application which monitors the health of the system. It runs as a Cluster-level controller. If failures are detected which do not meet certain Service Level Indicators (SLIs) then this app will interfere. It interacts with the Kubernetes API to trigger the relevant actions.

### `fearless-eagle`
A Prometheus instance which runs metrics collection for our watchdog.

### `giant-wasp`
A simple script/application which fires requests at `absurd-iguana`. This is mainly to simulate real traffic.

## ⚡️ Quick Start Guide

Here is a link to my quickstart guide. I think it is better practice to keep directions for such things in specific documents.

## 📚 Documentation
- Reliability Hooks
- SLIs