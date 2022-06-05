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
              envFrom:
                - configMapRef:
                    name: {{ .Values.name }}       
    ```

4. but this static ConfigMap is not so practical, adjust `templates/configmap.yaml` a bit

    ```yaml
    data:
      {{- range $key, $value := .Values.config }}
      {{ $key }}: {{ $value | quote }}
      {{- end }}
    ```

    and create a corresponding section in `values.yaml` (it should a object) but do not specify any configuration variables there.

    <details>
    <summary>Click to expand!</summary>

    ```yaml
    name: training-app
    image:
      name: 314595822951.dkr.ecr.eu-west-1.amazonaws.com/training/application
      tag: working
    config: {}
    ```
    </details>

5. create a new values file for the environment you want to parametrize e.g. `values-dev.yaml` and put some configuration values to the `config` object

    <details>
    <summary>Click to expand!</summary>

    ```yaml
    config:
      CONFIG_FIRST_VARIABLE: test1
    ```
    </details>

6. upgrade helm release

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```
    </details>

7. view environment variables in the application container

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl exec deploy/app -- env
    ```
    </details>

8. add more variables to `values-dev.yaml` (e.g. `CONFIG_SECOND_VARIABLE`) and upgrade release again

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```
    </details>

9. view environment variables in the application container

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl exec deploy/app -- env
    ```
    </details>

    > it did not change, right? Well configuration of the deployment did not
    > change so there was no need to perform rolling update. We need to fix
    > this!

10. adjust `templates/deployment.yaml`

    ```yaml
    spec:
      template:
        metadata:
          annotations:
            checksum/config: {{ include (print $.Template.BasePath "/configmap.yaml") . | sha256sum }}     
    ```

11. perform upgrade of the release

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```
    </details>

12. add more variables to `values-dev.yaml` (e.g. `CONFIG_THIRD_VARIABLE`) and upgrade release again

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```
    </details>

13. check status of the pods in your namespace and view the environment variables in the application container

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl get pods
    kubectl exec deploy/app -- env
    ```
    </details>

    > now, all the pods were rolled-out to the new version and new configuration
    > has been loaded.

14. list versions of your helm release

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm history training-app
    ```
    </details>

15. and rollback to the previous version

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm rollback training-app <number of previous revision>
    ```
    </details>

16. view environment variables in the application container

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl exec deploy/app -- env
    ```
    </details>

    > so this is how you can rollback anytinhg in case of emergency,
    > but this mechanism is usually driven by the automation and it's
    > more often based on the git history

