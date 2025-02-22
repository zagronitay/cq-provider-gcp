# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- 
## Unreleased
### 🚀 Added
### :gear: Changed
### :spider: Fixed
### 💥 Breaking Changes

-->

## [v0.5.1] - 2022-01-03
###### SDK Version: 0.6.1
### :spider: Fixed
* Fixed issues with disabled services [#84](https://github.com/cloudquery/cq-provider-gcp/pull/84)
### :gear: Changed
* Updated to SDK version [v0.6.1](https://github.com/cloudquery/cq-provider-sdk/blob/v0.6.1/CHANGELOG.md#v061---2022-01-03)


## [v0.5.0] - 2021-11-21
###### SDK Version: 0.5.0

### :rocket: Added
* Add support for error classifier in GCP. [#78](https://github.com/cloudquery/cq-provider-gcp/issues/78)

### :gear: Changed
* Upgraded to SDK Version [v0.5.0](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v050---2021-10-21)

## [v0.4.8] - 2021-10-07
###### SDK Version: v0.4.9

### :gear: Changed
* Upgraded to SDK Version [v0.4.9](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)
* Dropped column `post_key_revocation_action_type` from `gcp_compute_instances`


## [v0.4.7] - 2021-10-03
###### SDK Version: v0.4.7

### :rocket: Added
* added migration tests for improved stability [#67](https://github.com/cloudquery/cq-provider-gcp/pull/67)

### :spider: Fixed
* Fixed issues in integration tests [#69](https://github.com/cloudquery/cq-provider-gcp/pull/69) [#70](https://github.com/cloudquery/cq-provider-gcp/pull/70)

### :gear: Changed
* Upgraded to SDK Version [v0.4.7](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v047---2021-09-23)

## [v0.4.6] - 2021-09-14
###### SDK Version: v0.4.4

### :spider: Fixed
* Fixed forbidden error in gcp buckets iam policy fetch [#59](https://github.com/cloudquery/cq-provider-gcp/pull/59)

### :gear: Changed
* Upgraded to SDK Version [v0.4.4](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v044---2021-09-13)

### :rocket: Added
* added e2e intergration tests for improved stability [#51](https://github.com/cloudquery/cq-provider-gcp/pull/51)

## [v0.4.5] - 2021-09-09
###### SDK Version: 0.4.3

### :spider: Fixed
Embed migrations [#58](https://github.com/cloudquery/cq-provider-gcp/pull/58)

## [v0.4.4] - 2021-09-09
###### SDK Version: 0.4.3

### :rocket: Added
* added support for urlmaps resources [#47](https://github.com/cloudquery/cq-provider-gcp/issues/47)

### :spider: Fixed
* iam.service_accounts duplicate primary key fixed. [#53](https://github.com/cloudquery/cq-provider-gcp/pull/53)

### :gear: Changed
* Upgraded to SDK Version [0.4.3](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)

## [v0.4.3] - 2021-08-13
###### SDK Version: 0.3.2

### :spider: Fixed
* Add Ignore error to all resources. [#45](https://github.com/cloudquery/cq-provider-gcp/pull/45)

## [v0.4.2] - 2021-08-12
###### SDK Version: 0.3.2

### :spider: Fixed
* Fixed common errors in gcp provider. [#43](https://github.com/cloudquery/cq-provider-gcp/pull/43)

## [v0.4.1] - 2021-08-11
###### SDK Version: 0.3.2
### :gear: Changed
* Upgraded to SDK Version [0.3.2](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v032---2020-08-11)

## [v0.4.0] - 2021-07-27
###### SDK Version: 0.3.0

### :rocket: Added

* Added a changelog :)
* Added support for passing credentials in `config.hcl` [#35](https://github.com/cloudquery/cq-provider-gcp/pull/35) [#36](https://github.com/cloudquery/cq-provider-gcp/pull/36) 

### :gear: Changed
* Upgraded to SDK Version [0.3.0](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)
* **Breaking Change**: default CloudQuery "id" from `id` to `cq_id` [#41](https://github.com/cloudquery/cq-provider-sdk/pull/41)


### :spider: Fixed
* Fixed Domain registrations bad path. [#38](https://github.com/cloudquery/cq-provider-gcp/pull/38)

## [v0.3.6] - 2021-07-15
###### SDK Version: 0.2.8

Base version at which changelog was introduced.

### Supported Resources
- bigquery.datasets
- cloudfunctions.functions
- compute.addresses
- compute.autoscalers
- compute.backend_services
- compute.disk_types
- compute.disks
- compute.firewalls
- compute.forwarding_rules
- compute.images
- compute.instances
- compute.interconnects
- compute.urlmaps
- compute.networks
- compute.projects
- compute.ssl_certificates
- compute.ssl_policies
- compute.subnetworks
- compute.target_https_proxies
- compute.target_ssl_proxies
- compute.vpn_gateways
- crm.projects
- dns.managed_zones
- dns.policies
- domains.registrations
- iam.project_roles
- iam.service_accounts
- kms.keys
- logging.metrics
- logging.sinks
- monitoring.alert_policies
- resource_manager.folders
- resource_manager.projects
- sql.instances
- storage.buckets
