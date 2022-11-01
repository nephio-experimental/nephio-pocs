Knap
====

The Kubernetes Network Attachment Provider enables "network-as-a-service" for Kubernetes.

This is essentially "Multus, Part 2", a way for network function designers to package
additional Multus networks without requiring the cluster- and node-specific administrative
knowledge necessary for writing exact CNI configs. Instead, provider plugins take a few
general hints and generate those configs for you based on local knowledge that is preconfigured
and/or discovered.

See [repository](https://github.com/tliron/knap).
