## Memory and CPU settings

1. make sure you're in the correct namespace [link](./00_single_pod.md)

2. create following resource

    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: app
      labels:
        app: app
    spec:
      replicas: 1
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
              resources:
                requests:
                  cpu: 10000m
                  memory: 128Mi
                limits:
                  cpu: 10000m
                  memory: 128Mi
    ```

3. list all pods in your namespace

    ```bash
    kubectl get pods
    ```

    this is what happens when you don't have enought resources available.

    Try to describe on of the pods

    ```bash
    kubectl describe pod <pod name>
    ```

4. now, let's do some more realistic scenario, update deployment as follows

    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: app
      labels:
        app: app
    spec:
      replicas: 1
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
              resources:
                requests:
                  cpu: 100m
                  memory: 128Mi
                limits:
                  cpu: 100m
                  memory: 128Mi
    ```

    > Please not the volume section, this deployment has a special
    > ephemerail storage that is count towards the memory limits.
    > In the following step we'll try to fill this limit up.

    > In your environment, you can simulate this with your Java
    > applications :D

5. now, try to fill these 100MiB and check the usage

    ```bash
    curl https://<your namespace>.s01.training.eks.rocks/fill
    ```

    or

    ```powershell
    Invoke-WebRequest -Headers https://<your namespace>.s01.training.eks.rocks/fill
    ```

    and then

    ```bash
    kubectl top pods
    ```

    > Please note that it will take some time until
    > values in `top pod` are updated.

6. send a few more requests there and see what happened

    ```bash
    curl https://<your namespace>.s01.training.eks.rocks/fill
    ```

    or

    ```powershell
    Invoke-WebRequest -Headers https://<your namespace>.s01.training.eks.rocks/fill
    ```

    and then

    ```bash
    kubectl get pods
    ```

    or

    ```bash
    kubectl describe pods
    ```

7. Perform cleanup

    ```bash
    kubectl delete deploy app
    kubectl delete svc app
    kubectl delete ing app
    kubectl delete cm app
    kubectl delete secret app
    ```
