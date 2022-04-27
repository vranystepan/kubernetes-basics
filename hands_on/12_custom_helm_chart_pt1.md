## custom Helm charts for your application pt. I

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. in your wirking directory, create a new directory with arbitrary name e.g. `chart`

    ```bash
    mkdir chart
    cd chart
    ```

3. create a manifest file for our new Helm chart `Chart.yaml`

    ```yaml
    apiVersion: v2
    name: training-app
    description: A Helm chart for training session
    type: application
    version: 0.0.1
    appVersion: "1.16.0"
    ```

4. create a new `templates` directory

    ```bash
    mkdir templates
    ```

5. and create a new file `deployment.yaml` in this directory (`templates/deployment.yaml`)

    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: app
      labels:
        app: app
    spec:
      replicas: 3
      selector:
        matchLabels:
          app: app
      template:
        metadata:
          labels:
            app: app
        spec:
          containers:
            - name: app
              image: 314595822951.dkr.ecr.eu-west-1.amazonaws.com/training/application:working
              ports:
                - containerPort: 8080
    ```

6. perform a quick check of your chart with `helm template` command

    ```bash
    helm template training-app .
    ```

7. let's parametrize the image name, change `templates/deployment.yaml` to this

    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: app
      labels:
        app: app
    spec:
      replicas: 3
      selector:
        matchLabels:
          app: app
      template:
        metadata:
          labels:
            app: app
        spec:
          containers:
            - name: app
              image: {{ .Values.image }}
              ports:
                - containerPort: 8080
    ```

8. and create a new file `values.yaml` with following contents

    ```yaml
    image: 314595822951.dkr.ecr.eu-west-1.amazonaws.com/training/application:working
    ```

9. perhaps we can make the parametrization more granular, change `templates/deployment.yaml` to this

    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: app
      labels:
        app: app
    spec:
      replicas: 3
      selector:
        matchLabels:
          app: app
      template:
        metadata:
          labels:
            app: app
        spec:
          containers:
            - name: app
              image: {{ .Values.image.name }}:{{ .Values.image.tag }}
              ports:
                - containerPort: 8080
  
11. and update `values.yaml` accordingly

    ```yaml
    image:
      name: 314595822951.dkr.ecr.eu-west-1.amazonaws.com/training/application
      tag: working
    ```

12. check the rendered resources with `helm template` command again

    ```bash
    helm template training-app .
    ```

    > please note that we don't need to specify `--values` flag,
    > since this `values.yaml` file is located directly in the chart
    > stucture - it's used automatically. We call it default values.    

13. let's install this helm chart to the Kubernetes

    ```bash
    helm upgrade --install training-app .
    ```

14. darn, it did not work. Let's parametrize names a bit

    `templates/deployment.yaml`:

    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: {{ .Values.name }}
      labels:
        app: {{ .Values.name }}
    spec:
      replicas: 3
      selector:
        matchLabels:
          app: {{ .Values.name }}
      template:
        metadata:
          labels:
            app: {{ .Values.name }}
        spec:
          containers:
            - name: app
              image: {{ .Values.image.name }}:{{ .Values.image.tag }}
              ports:
                - containerPort: 8080
    ```

    `values.yaml`:

    ```yaml
    name: training-app
    image:
      name: 314595822951.dkr.ecr.eu-west-1.amazonaws.com/training/application
      tag: working 
    ```

15. and try to install this chart again

    ```bash
    helm upgrade --install training-app .
    ```

