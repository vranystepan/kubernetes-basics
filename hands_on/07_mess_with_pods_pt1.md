## Intoduce broken version of the app

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
              image: 314595822951.dkr.ecr.eu-west-1.amazonaws.com/training/application:breaking
              ports:
                - containerPort: 8080
    ```

3. check the status of all pods

    ```bash
    kubectl get pods
    kubectl logs <name of the new pod>
    kubectl describe pod <name of the new pod>
    ```

    > Please note that at least one pod is still running,
    > this is called RollingUpdate and this is the way
    > how to prevent downtimes in production

4. try to send some requests to your service

    ```bash
    curl https://<your namespace>.s01.training.eks.rocks -H 'User-Agent: workstation'
    ```

    and check how the endpoints lokk

    ```bash
    kubectl get endpoints
    ```

5. now we're gonna deploy slowly starting service. This simulates some startup tasks that sometimes happen in applications.  Create a local file with the following contents:

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
    ```


4. try to send some requests to your service

    ```bash
    curl https://<your namespace>.s01.training.eks.rocks -H 'User-Agent: workstation'
    ```

    or

    ```powershell
    Invoke-WebRequest https://<your namespace>.s01.training.eks.rocks
    ```

    ... we just caused a production outage.
