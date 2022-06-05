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

    > And here we go again, when you look at `spec.template.spec`, what do you see there? Yup, it's just pod + metadata. Same reason, job controller will ultimately create a pod since that's the only way to run containers in Kubernetes.

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

    > Try to count the failed pods. Why it does not run infinitely? That's because we have backoffLimit specified.

9. view the logs

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl logs job/job
    ```
    </details>

    > please note that we don't need to specify the pod's name, Kubernetes is able to derive it from the job.

10. proceed to the [next section](04_periodically_running_job.md)
