# adhdtracker

My personal take on making a task tracker

# Minimal Viable Product:

 GUI

User(s): Each one gets a

Personal Calendar

Tasks:
    Not Done -> In Progress -> Done

Progress Bar of all Tasks


Basic Makefile to build the service

Execute on the service to run it in ./bin/adhdtracker

## Programming languages:

### Back end - Go & Python
Go: DevOps/SRE world is being more dominated by Go, and I want to learn it.

Python: I never fully learned Python but I am familiar with it.

### Front end - Node.js, React, Javascript
Node.js is used by Grafana and I understand grafana’s visualizations & open source libraries

React is heavily being used more and more

Javascript is a known evil and must be used but I wanna try and limit it.

### Makefile
to build out the backend and then frontend: Reference

##Concepts I wish to incorporate:
Objects

Web GUI

## Infrastructure:
### Components:
Docker: Build a docker image to run the entire thing.

Kubernetes: Once we get to scaling up an infrastructure, will use that as a base.

Terraform: comes in much later when we start provisioning this into Cloud Providers and kubernetes is built up.

### Config Management:
Ansible - Much later

### Cloud Provider:
AWS or GCP - Later

### Backend:
Use SQL as a backend to store - Sqlite for now

Use Redis to cache & optimize - Not needed but looking out to the future

### Tooling:
Metrics, Logging and Visualization

cAdvisor -> Prometheus -> Cortex

FluentBit -> Loki

Grafana

### Tracing:
Tempo: Down the road

Jaeger: Down the road

### External Monitoring
Prometheus blackbox

### QA
E2e tests

Canary deployment

### Data Science:
Based on Users and transaction flow

Which Tasks are getting done faster

What improvements in their life are making changes happen

# References:

Jess ideas:

A personal Sunsuma

The personal task tracker could ask people what’s important to them and then let them color code their tasks based on their values so they can see if they’re doing a lot of stuff to tend to one value and not enough to tend to another

ADHDers tend to hyperfocus and neglect things that are important to them, so would be an amazing fit

Plus it would help with mindfulness. You’d see if you didn’t have a lot of friendship  related tasks and realize maybe that needs some attention BEFORE your friends are all mad at you 

