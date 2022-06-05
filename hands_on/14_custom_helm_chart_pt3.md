## custom Helm charts for your application pt. III

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. and create a new file `configmap.yaml` in this directory (`templates/configmap.yaml`)

    ```yaml
    apiVersion: batch/v1
    kind: Job
    metadata:
      name: {{ .Values.name }}-migrations
      annotations:
        "helm.sh/hook": pre-install,pre-upgrade	
        "helm.sh/hook-weight": "-5"
        "helm.sh/hook-delete-policy": before-hook-creation
    spec:
      template:
        metadata:
          name: "{{ .Values.name }}"
        spec:
          restartPolicy: Never
          containers:
          - name: migrations
            image: "alpine:latest"
            command:
              - sh
              - -c
              - |
                sleep 45
                echo migrations completed
    ```

3. perform upgrade of the release and watch what's happening

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm upgrade --install training-app . --values values-dev.yaml
    ```
    </details>
