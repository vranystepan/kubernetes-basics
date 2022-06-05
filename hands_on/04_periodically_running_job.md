## Periodically running job

CronJob is the same idea like Crons in the Unix/Linux/whatever systems.
You put there time expression and stuff is happening periodically.
This resource is widely used in PHP-based services but you can see it
even in other (non-PHP) stacks.

Even I'm using this for some infrastructure tasks like backups
of self-managed services. It just works.

---

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. create a local file with the following contents:

    ```yaml
    apiVersion: batch/v1
    kind: CronJob
    metadata:
      name: cronjob
    spec:
      successfulJobsHistoryLimit: 2
      schedule: "* * * * *"
      jobTemplate:
        spec:
          template:
            spec:
              containers:
              - name: cronjob
                image: nginx
                imagePullPolicy: IfNotPresent
                command:
                  - sh
                  - -c
                  - |
                    curl http://app -H 'User-Agent: cronjob'
              restartPolicy: OnFailure
    ```

3. wait a few seconds and list spawned jobs

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl get job
    ```
    </details>

4. also, check the pods spawned by these jobs

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl get pods
    ```
    </details>

5. and try to manually submit a new job with the cronjob's configuration

    ```bash
    kubectl create job --from=cronjob/cronjob cronjob-manual-01
    ```

    > Why would you need this? Sometimes it might happen that
    > Application was broken during the given schedule and you
    > simply want to run it right away.

6. and once again, check the generated resources

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl get jobs
    kubectl get pods
    ```
    </details>

7. delete the cronjob

    <details>
    <summary>Click to expand!</summary>

    ```bash
    kubectl delete cj cronjob
    ```
    </details>
