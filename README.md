# ![pipe](./docs/images/pipe.png) Pipeline CRD

[![Go Report Card](https://goreportcard.com/badge/knative/build-pipeline)](https://goreportcard.com/report/knative/build-pipeline)

The Pipeline CRD provides k8s-style resources for declaring CI/CD-style
pipelines.

Pipelines are **Cloud Native**:

- Run on Kubernetes
- Have Kubernetes clusters as a first class type
- Use containers as their building blocks

Pipelines are **Decoupled**:

- One Pipeline can be used to deploy to any k8s cluster
- The Tasks which make up a Pipeline can easily be run in isolation
- Resources such as git repos can easily be swapped between runs

Pipelines are **Typed**:

- The concept of typed resources means that for a resource such as an `Image`,
  implementations can easily be swapped out (e.g. building with
  [kaniko](https://github.com/GoogleContainerTools/kaniko) v.s.
  [buildkit](https://github.com/moby/buildkit))

## Want to start using Pipelines?

- Jump in with [the tutorial!](docs/tutorial.md)
- [Learn about the Concepts](/docs/Concepts.md)
- [Read about how to use it](/docs/using.md)
- Look at [some examples](/examples)

## Want to contribute?

We are so excited to have you!

- See [CONTRIBUTING.md](CONTRIBUTING.md) for an overview of our processes
- See [DEVELOPMENT.md](DEVELOPMENT.md) for how to get started
- [Deep dive](./docs/developers/README.md) into demystifying the inner workings
  (advanced reading material)
- Look at our
  [good first issues](https://github.com/knative/build-pipeline/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
  and our
  [help wanted issues](https://github.com/knative/build-pipeline/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
