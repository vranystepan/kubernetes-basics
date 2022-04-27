## custom Helm charts for your application pt. II

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. and create a new file `configmap.yaml` in this directory (`templates/configmap.yaml`)

    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: {{ .Values.name }}
    data:
      CONFIG_SOME_VALUE_1: value1
    ```

3. and set it as the source of the environment variables in `templates/deployment.yaml`

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
              envFrom:
                - configMapRef:
                    name: {{ .Values.name }}       
    ```

4. but this static ConfigMap is not so practical, adjust `templates/configmap.yaml` a bit

    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: {{ .Values.name }}
    data:
      {{- range $key, $value := .Values.config }}
      {{ $key }}: {{ $value | quote }}
      {{- end }}
    ```

    and create a corresponding section in `values.yaml` but do not specify any
    configuration variables there.

    ```yaml
    name: training-app
    image:
      name: 314595822951.dkr.ecr.eu-west-1.amazonaws.com/training/application
      tag: working
    config: {}
    ```

5. create a new values file for the environment you want to parametrize e.g. `values-dev.yaml`


    ```yaml
    config:
      CONFIG_FIRST_VARIABLE: test1
    ```

6. upgrade helm release

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```

7. view environment variables in the application container

    ```bash
    kubectl exec deploy/app -- env
    ```

8. add more variables to `values-dev.yaml` (e.g. `CONFIG_SECOND_VARIABLE`) and upgrade release again

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```

9. view environment variables in the application container

    ```bash
    kubectl exec deploy/app -- env
    ```

    > it did not change, right? Well configuration of the deployment did not
    > change so there was no need to perform rolling update. We need to fix
    > this!

10. adjust `templates/deployment.yaml`

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
          annotations:
            checksum/config: {{ include (print $.Template.BasePath "/configmap.yaml") . | sha256sum }}
        spec:
          containers:
            - name: app
              image: {{ .Values.image.name }}:{{ .Values.image.tag }}
              ports:
                - containerPort: 8080
              envFrom:
                - configMapRef:
                    name: {{ .Values.name }}       
    ```

11. perform upgrade of the release

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```

12. add more variables to `values-dev.yaml` (e.g. `CONFIG_THIRD_VARIABLE`) and upgrade release again

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```

13. check status of the pods in your namespace and view the environment variables in the application container

    ```bash
    kubectl get pods
    kubectl exec deploy/app -- env
    ```

    > now, all the pods were rolled-out to the new version and new configuration
    > has been loaded.

14. list versions of your helm release

    ```bash
    helm history training-app
    ```

15. and rollback to the previous version

    ```bash
    helm rollback training-app <number of previous revision>
    ```

16. view environment variables in the application container

    ```bash
    kubectl exec deploy/app -- env
    ```

    > so this is how you can rollback anytinhg in case of emergency,
    > but this mechanism is usually driven by the automation and it's
    > more often based on the git history

