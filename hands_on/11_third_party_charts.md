## Third-party Helm charts

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. add bitnami helm chart repository

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm repo add bitnami https://charts.bitnami.com/bitnami
    ```
    </details>

3. list releases in you namespace, the list should be empty

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm ls
    ```
    </details>

4. install a new release to your namespace

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm install nginx bitnami/nginx
    ```
    </details>

5. try to execute the same command again

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm install nginx bitnami/nginx
    ```
    </details>

6. uninstall the release

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm uninstall nginx
    ```
    </details>

7. and install the same release with upgrade command (and `--install` flag)

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm upgrade --install nginx bitnami/nginx
    ```
    </details>

8. try to execute the same command again

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm upgrade --install nginx bitnami/nginx
    ```
    </details>

    > Upgrade command with `--install` flag is idempotent, it perfectly
    > fit to the automation when you don't always know whether you're
    > installing or upgrading

9. list releases in your namespace

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm ls
    ```
    </details>

10. display details about your release

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm status nginx
    ```
    </details>

11. show history of this release

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm history nginx
    ```
    </details>

12. in the same directory, create a new file `values.yaml`

13. open vendor's documentation and examine possible values https://artifacthub.io/packages/helm/bitnami/nginx?modal=values

14. try to add more replicas (2) to your values file, just copy & paste respective line from the default values to your `values.yaml file`

15. upgrade the release again with the custom values file

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm upgrade --install nginx bitnami/nginx --values values.yaml
    ```

    or

    ```bash
    helm upgrade --install nginx bitnami/nginx -f values.yaml
    ```
    </details>

16. show history of this release

    <details>
    <summary>Click to expand!</summary>

    ```bash
    helm history nginx
    ```
    </details>
