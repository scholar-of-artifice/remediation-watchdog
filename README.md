# remediation-watchdog

![License](https://img.shields.io/badge/license-MIT-green.svg?labelColor=white)

> **A standalone "Control Plane" service designed to manage the availability of a target application through a closed-loop feedback system.**

## 🧑‍💻 Technologies

<!--technology badges here-->
![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-%23DD0031.svg?logo=redis&logoColor=white)
![Apache Kafka](https://img.shields.io/badge/Apache%20Kafka-000?logo=apachekafka&logoColor=fff)
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
A basic `Go` application that validates incoming requests and prodces events to the `dazzling-remora` message bus.

### `dazzling-remora`
A Kafka broker running `KRaft` mode that serves as the system's durable message backbone. It allows for asynchronous processing and provides the primary metrics for system saturation.

### `eager-marmot`
A `Go` application which consumes tasks from `dazzling-remora` and executes the final persistence to `bashful-yak`. This service is designed to be horizontally scaled to manage spikey traffic loads.

### `bashful-yak`
A simple `Redis` instance which stores from `eager-marmot`. For this application, complicated data models are not the focus.

### `calm-lynx`
The system's control plane. It observes SLIs across the entire pipeline, including web latency in `absurd-iguana`, consumer lag in `dazzling-remora`, and connection health in `bashful-yak`. It automatically executes remediation scripts to resolve identified bottlenecks.

### `fearless-eagle`
The source of truth for all system telemetry. It provides the high-cardinality metrics that `calm-lynx` uses to evaluate the health of the data pipeline.

### `giant-wasp`
A synthetic traffic generator used to simulate real-world user demand. It is the primary tool used to test the system's auto-scaling and remediation limits under stress.

## Kubernetes Reliability Primitives

This project utilizes native Kubernetes features to ensure baseline reliability.

### Liveness Probes
Automatically restarts `eager-marmot` if the Go runtime enters a deadlocked state.

### Readiness Probes
Prevent `absurd-iguana` from accepting traffic until it has successfully established connection to `dazzling-remora`.

### Resource Quota
Simulated "noisy neighbor" scenarios are managed by setting strict resource limits, allowing for the testing of `calm-lynx` under resource pressure.

## ⚡️ Quick Start Guide

Here is a link to my quickstart guide. I think it is better practice to keep directions for such things in specific documents.

## 📚 Documentation
- Reliability Hooks
- SLIs