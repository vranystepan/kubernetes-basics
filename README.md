# Ataccama Kubernetes training

## requirements

- Mac OS, Linux, WSL or Powershell
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/docs/intro/install/)
- web browser
- curl

### check
- execute `aws --version`, CLI should return the version
- execute `kubectl version`, CLI should return the version
- execute `helm version`, CLI should return the version

### AWS cli configuration
Please follow configuration instruction in the [dedicated
section](./docs/00_aws_configuration.md) of this repository.


### infrastructure
This repository is using EKS cluster created by Terraform. You
can find these definitions in the separete [GitHub repository](https://github.com/vranystepan/kubernetes-basics-infrastructure).

