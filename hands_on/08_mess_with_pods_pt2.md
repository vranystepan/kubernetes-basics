## Health probes

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. create a local file with the following contents:

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
              image: 314595822951.dkr.ecr.eu-west-1.amazonaws.com/training/application:sleeping
              ports:
                - containerPort: 8080
                  name: http
              readinessProbe:
                periodSeconds: 10
                httpGet:
                  port: http
                  path: /_health/ready
    ```

    > please note newly intoroduced `readinessProbe` stanza.

3. wait for the finished RollingUpdate and send following request to your ingress

    ```bash
    curl https://<your namespace>.s01.training.eks.rocks/_health/set/notready -H 'User-Agent: workstation' -v
    ```

4. watch what's happening in your namespace

    ```bash
    watch kubectl get pod
    ```

5. try do describe one of your pod

    ```bash
    kubectl describe pod <name of the pod>
    ```

6. now, list endpoints in your namespace

    ```bash
    kubectl get endpoints
    ```

7. open a new shell in not-ready pod and make it work again

    ```bash
    curl http://localhost:8080/_health/set/ready
    ```

8. watch what's happening in your namespace

    ```bash
    watch kubectl get pod
    ```

9. list endpoints in your namespace

    ```bash
    kubectl get endpoints
    ```

10. let's apply the next change:

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
                  name: http
              readinessProbe:
                periodSeconds: 10
                httpGet:
                  port: http
                  path: /_health/ready
              livenessProbe:
                initialDelaySeconds: 60
                periodSeconds: 10
                httpGet:
                  port: http
                  path: /_health/alive
    ```

    > please note newly intoroduced `livenessProbe` stanza.

11. and simulate failure with following http request:

    ```bash
    curl https://<your namespace>.s01.training.eks.rocks/_health/set/notalive -H 'User-Agent: workstation' -v
    ```

12. watch what's happening in your namespace

    ```bash
    watch kubectl get pod
    ```
