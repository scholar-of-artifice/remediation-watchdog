# remediation-watchdog

![License](https://img.shields.io/badge/license-MIT-green.svg?labelColor=white)

> **A standalone "Control Plane" service designed to manage the availability of a target application through a closed-loop feedback system.**

## 🧑‍💻 Technologies

<!--technology badges here-->
![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-%23DD0031.svg?logo=redis&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff)
![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?logo=kubernetes&logoColor=fff)

## 🏗️ High Level Architecture

In this example system which demonstrates how to write simple programs to observe and remediate issues in an orchestrated system. 

[TODO: mermaid chart here eventually. not right now until the project is further along.]

The image above is a basic map of the different components. In this system we have a few parts:

### `absurd-iguana`
A basic `Go` application a which takes in requests and writes data to a database.

### `bashful-yak`
A simple `Redis` instance which stores data. For this application, complicated data models are not the focus.

### `calm-lynx`
This is a `Go` application which monitors the health of `absurd-iguana` and `bashful-yak`. If failures are detected which do not meet certain Service Level Indicators (SLIs) then this app will interfere.

## ⚡️ Quick Start Guide

Here is a link to my quickstart guide. I think it is better practice to keep directions for such things in specific documents.

## 📚 Documentation