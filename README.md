[![](https://github.com/mchmarny/disco/actions/workflows/on-push.yaml/badge.svg?branch=main)](https://github.com/mchmarny/disco/actions/workflows/on-push.yaml)
[![](https://github.com/mchmarny/disco/actions/workflows/on-tag.yaml/badge.svg)](https://github.com/mchmarny/disco/actions/workflows/on-tag.yaml)
[![](https://codecov.io/gh/mchmarny/disco/branch/main/graph/badge.svg?token=9HLYDZZADN)](https://codecov.io/gh/mchmarny/disco)
[![version](https://img.shields.io/github/release/mchmarny/disco.svg?label=version)](https://github.com/mchmarny/disco/releases/latest)
[![](https://img.shields.io/github/go-mod/go-version/mchmarny/disco.svg?label=go)](https://github.com/mchmarny/disco)
[![](https://goreportcard.com/badge/github.com/mchmarny/disco)](https://goreportcard.com/report/github.com/mchmarny/disco)
[![](https://img.shields.io/badge/License-Apache%202.0-blue.svg?label=license)](https://github.com/mchmarny/disco/blob/main/LICENSE)

# disco 

Utility for bulk image, license, and vulnerability discovery in containerize workloads on GCP.

> Note: this is a personal project, not an official Google product.

Features:

* Runs as a CLI or a Cloud Run service
* Discover currently deployed container images
  * multiple project and region report with filters
  * deployed image to digest resolution
* Report on vulnerabilities or licenses in these images
  * supports operating system and package-level scans
* Find out if those images are impacted by a specific CVE
* When ran as a service:
  * Creates time-series metrics in Cloud Monitoring (charts, alerts)
  * Exports License and Vulnerability data to BigQuery tables (reports)

![](etc/img/dashboard.png)

## Why

It's easy to end up with a large number of containerized workloads across many GCP projects and regions: Cloud Run, GKE, or even Cloud Functions (yes, those end up running as a container too). You can scan these containers in Artifact Registry using [Container Analysis](https://cloud.google.com/container-analysis/docs/container-analysis) service, but currently it only [covers base OS](https://cloud.google.com/container-analysis/docs/os-overview). It's also not easy to know which of these images (and which versions) are actually being used in active services. Services like Cloud Run also support [multiple revisions](https://cloud.google.com/run/docs/managing/revisions), each potentially using a different version of an image.

`disco` provides an easy way to `disco`ver which of these container images are currently deployed. And, if one of the supported open source scanners is installed, `disco` automatically scans these images for any vulnerabilities, or lists the types of licenses used in those images. 

## Install

You can use `disco` either as CLI or Service. You can find information on how to install `disco` in either of these formats using the following links:

* [CLI](CLI.md) - Available via the most common distribution methods (Homebrew, RPM, DEB, Go, Binary etc).
* [Service](SERVICE.md) - Deploys as a Cloud Run service via Terraform.

## OSS

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmchmarny%2Fdisco.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmchmarny%2Fdisco?ref=badge_large)


## Disclaimer

This is my personal project and it does not represent my employer. While I do my best to ensure that everything works, I take no responsibility for issues caused by this code.
