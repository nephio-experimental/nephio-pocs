WARNING: This repository is published by the [Nephio Authors](https://nephio.org/) but is
neither endorsed nor maintained by the Nephio Technical Steering Committee (TSC). It is intended
to be used for reference only. The Nephio distribution repositories are located in the
[`nephio-project` organization](https://github.com/nephio-project). For more information
[see this page](https://nephio.org/experimental).

Nephio: Proofs of Concepts
==========================

This repository collects information and code for proofs of concepts (PoCs) presented to and
contributed to the [Nephio project](https://nephio.org/).

**Note**: Issues should be opened in the [nephio](https://github.com/nephio-project/nephio) repository,
using the prefix "nephio-pocs: " in the issue title.

Pull requests are welcome. Each PoC has its own directory in this repository, which should
minimally contain a `README.md` file with a description and link, but may contain the full source
code and documentation for the PoC. Additionally, each PoC should be listed in the summary
below in alphabetical order and with a short one-paragraph description relating it directly
to Nephio's goals:

* [Candice](candice/):
  Allows local NETCONF and NETCONF-like workflows (to CNFs and adjacent PNFs)
  to be managed declaratively in Kubernetes.
* [CNCK](cnck/):
  Allows for network functions that are not designed to be cloud native
  (do not self-configure or self-orchestrate) to work better in a cloud native
  environment by handling dependency discovery and configuration.
* [ENO](eno/):
  External Network Operator (ENO) is a framework that enables network automation in Kubernetes.
  Exposes a common API which allows the dynamic orchestration of networks on cluster and fabric
  level and therefore gives the ability to applications to consume a high number of different networks.
  ENO intends to run in workload clusters under Nephio context and will be responsible for the network
  provision inside the cluster as well as in the fabric.
* [free5GC Operator](free5gc-operator/):
  Deploys and manages [free5GC](https://www.free5gc.org/)'s AMF, SMF, and UPF components on
  Kubernetes.
* [Knap](knap/):
  Enables "network-as-a-service" for Kubernetes by allowing network functions
  to specify Multus network requirements in general terms without requiring local cluster-
  or node-specific knowledge. Knap then generates Multus CNI configs automatically.
* [Nephio NF Controllers](nephio-nf-controllers/)
  An example of network function controllers integrated with [kpt](https://kpt.dev/).
* [Nephio Package Deployment Controller](https://github.com/nephio-project/nephio-controller-poc/)
  Shows how to use Porch to render package variants across a set of cluster
  repositories, while injecting cluster specific context and reconfiguring each
  package based on that context.
* [Nephio 5gc Topology Controller](nephio-5gc-controller/)
  Builds on top of the NF controller and the pacakge deployment controller to
  render multiple network functions across multiple, different types of
  clusters.
* [NF Injector
  Controller](https://github.com/henderiw-nephio/nf-injector-controller)
  Builds on the Nephio 5gc Topology controller to inject a `UPFDeployment`
  resource and associated IPAM allocation requests.
* [K8s IPAM Controller](https://github.com/nokia/k8s-ipam)
  Provides an extensible IPAM system, along with a reference implementation. It
  builds on top of the 5gc topology controller and NF injector controller to
  perform IP allocations and inject the results back into the package.
* [Nephio/Porch WebUI](https://github.com/GoogleContainerTools/kpt-backstage-plugins)
  A prototype web UI for Porch that can be [configured
  ](https://github.com/nephio-project/nephio-packages/tree/main/nephio-webui) for use in Nephio.
* [Planter](planter/):
  A meta-scheduler for Kubernetes that enables an all-at-once fire-and-forget declarative
  frontend for the complex lifecycle management of the workload-and-its-cluster interrelationships.
  Addresses the ["bifurcation"](https://www.youtube.com/watch?v=6FULuWvXR84)
  (or "chicken-and-egg") problem that arises from vertically-integrated network function workloads.
* [nf-deploy-fn](https://github.com/henderiw-nephio/nf-deploy-fn)
  Mutates and injects the IP allocationss in `UPFDeployment`. Runs as a kpt function in the kpt pipeline.
* [nad-inject-fn](https://github.com/henderiw-nephio/nad-inject-fn)
  Injects/generates Multus network attachement definitions based on a `UPFDeployment` and IP allocations. Runs as a kpt function in the kpt pipeline

## ONE Summit 2022 Workshop

For the ONE Summit in 2022, we tied together many of these prototypes into a
larger working prototype. After the summit, we will publish a procedure for
provisioning a complete sandbox environment to experiment with this.
