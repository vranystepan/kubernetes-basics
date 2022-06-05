## Get stuff done with one-off Jobs

Standalone jobs are not so often seen but still, they do exist.
One good example: database migrations before updating the Deployment.

---

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. create a local file with the following contents:

    ```yaml
    apiVersion: batch/v1
    kind: Job
    metadata:
      name: job
    spec:
      template:
        spec:
          containers:
            - name: job
              image: nginx
              command:
                - sh
                - -c
                - |
                  curl http://app -H 'User-Agent: job'
          restartPolicy: Never
      backoffLimit: 4
    ```

3. list the pods

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl get pods
    ```
    </details>

4. and get the logs of the pod spawned by this job

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl logs <pod name>
    ```
    </details>

5. optionally get logs from the app service and verify if there's a new event

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl logs deploy/app
    ```
    </details>

6. delete the job

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl delete job job
    ```
    </details>

7. update command in the job so it points to the incorrect hostname

    ```yaml
              command:
                - sh
                - -c
                - |
                  curl http://this-does-not-exist
    ```

8. and watch what's happening in your namespace

    <details>
    <summary>Click to expand!</summary>

    ```bash
    watch kubectl get pods
    ```
    </details>

9. view the logs

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl logs job/job
    ```
    </details>

10. proceed to the [next section](04_periodically_running_job.md)
