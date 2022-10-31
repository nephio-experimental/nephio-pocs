Planter
=======

Planter is a meta-scheduler for Kubernetes.

It is an operator that runs in management clusters and delegates workloads to one or more
workload clusters while also allowing the workloads to (re)configure those very workload
clusters on which they are to be deployed.

Planter is intended as a deliberately narrow solution to the ["bifurcation"](https://www.youtube.com/watch?v=6FULuWvXR84)
(or "chicken-and-egg") problem that arises from vertically-integrated workloads. It enables
an all-at-once fire-and-forget declarative frontend for the complex lifecycle management of
the workload-and-its-cluster interrelationships.

See [repository](https://github.com/tliron/planter).
