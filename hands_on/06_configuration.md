## configuration management

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. create local files with the following contents:

    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: app
    data:
      CONFIG_SOME_VALUE_1: value1
    ```

    ```yaml
    apiVersion: v1
    data:
      CONFIG_SOME_SECRET_VALUE_1: c2VjcmV0MQ==
    kind: Secret
    metadata:
      name: app
    type: Opaque
    ```

    > please note the value of `CONFIG_SOME_SECRET_VALUE_1`, this is base64-encoded
    > string.

3. also, adjust `app` deployment a bit, following snippet needs to be configured on the container level

    ```yaml
              env:
                - name: CONFIG_SOME_SECRET_VALUE_1
                  valueFrom:
                    secretKeyRef:
                      name: app
                      key: CONFIG_SOME_SECRET_VALUE_1
                - name: CONFIG_SOME_VALUE_1
                  valueFrom:
                    configMapKeyRef:
                      name: app
                      key: CONFIG_SOME_VALUE_1
    ```

4. wait a few seconds and open a new shell to your application container


    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl exec -it deploy/app -- bash
    ```
    </details>

5. and list the environment variables


    <details>
    <summary>Click to expand!</summary>

    ```bash
    env | grep "CONFIG_"
    ```
    </details>

6. update the `app` deployment again

    ```yaml
              envFrom:
                - configMapRef:
                    name: app
                - secretRef:
                    name: app
    ```

7. wait a few seconds and open a new shell to your application container


    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl exec -it deploy/app -- bash
    ```
    </details>

8. and list the environment variables


    <details>
    <summary>Click to expand!</summary>

    ```bash
    env | grep "CONFIG_"
    ```
    </details>

9. try to add some value to the ConfigMap, locally or via `edit` command

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl edit cm app
    ```
    </details>

10. open the shell again and list the environment variables


    <details>
    <summary>Click to expand!</summary>

    ```bash
    env | grep "CONFIG_"
    ```
    </details>

    > It's not there, right? We'll discuss this in the Helm section.
    > Stay tuned.

11. update the `app` deployment again and add the volume

    ```yaml
          volumes:
            - name: config
              configMap:
                name: app
    ```

    also, add a volume mount to the container

    ```yaml
              volumeMounts:
                - name: config
                  mountPath: /etc/config
    ```

12. wait a few seconds and open a new shell to your application container

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl exec -it deploy/app -- bash
    ```
    </details>

13. list files in `/etc/config`

    <details>
    <summary>Click to expand!</summary>

    ```bash
    ls /etc/config
    cat /ect/config/*
    ```
    </details>

14. introduce some new change to `app` configmap

15. open a new shell in your pod

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl exec -it deploy/app -- bash
    ```
    </details>

16. and list files in `/etc/config` again

    <details>
    <summary>Click to expand!</summary>

    ```bash
    ls /etc/config
    cat /ect/config/*
    ```
    </details>
