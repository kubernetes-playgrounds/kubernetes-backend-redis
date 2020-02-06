This project is to showcase how to deploy backend in kubernetes and rely on kube's internal dns to talk to redis service.

# To run

1. make sure you have a running kubernetes cluster (minikube will work, see [here](https://kubernetes.io/docs/tutorials/hello-minikube/) for details on how to run minikube).

2. run backend as pod in kube cluster

   ```
   $ cd k8s && kubectl apply -f backend.yml
   ```

3. run redis as pod/service in kube cluster

   ```
   $ cd k8s && kubectl apply -f redis.yml
   ```

4. verify the backend is talking to redis working by

   ```
   $ kubectl get pods
   $ kubectl logs -f <backend-pod-name-from-output-above>
   $ (in another terminal) kubectl logs -f <server-pod-name-from-output-above>
   ```

   you should see backend prints `PONG` continuously
