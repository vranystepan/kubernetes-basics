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

## Contents

### Basic Kubernetes resources

- [./hands_on/00_single_pod.md](./hands_on/00_single_pod.md)
- [./hands_on/01_deployment.md](./hands_on/01_deployment.md)
- [./hands_on/02_simple_service.md](./hands_on/02_simple_service.md)
- [./hands_on/03_one_off_job.md](./hands_on/03_one_off_job.md)
- [./hands_on/04_periodically_running_job.md](./hands_on/04_periodically_running_job.md)
- [./hands_on/05_ingress.md](./hands_on/05_ingress.md)
- [./hands_on/06_configuration.md](./hands_on/06_configuration.md)

### Basic operations

- [./hands_on/07_mess_with_pods_pt1.md](./hands_on/07_mess_with_pods_pt1.md)
- [./hands_on/08_mess_with_pods_pt2.md](./hands_on/08_mess_with_pods_pt2.md)
- [./hands_on/09_resources.md](./hands_on/09_resources.md)

### Helm basics

- [./hands_on/11_third_party_charts.md](./hands_on/11_third_party_charts.md)
- [./hands_on/12_custom_helm_chart_pt1.md](./hands_on/12_custom_helm_chart_pt1.md)
- [./hands_on/13_custom_helm_chart_pt2.md](./hands_on/13_custom_helm_chart_pt2.md)
- [./hands_on/14_custom_helm_chart_pt3.md](./hands_on/14_custom_helm_chart_pt3.md)
