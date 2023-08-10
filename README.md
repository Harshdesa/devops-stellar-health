# devops-stellar-health

Docker hub location:
https://hub.docker.com/repository/docker/harshdesai7979/stellarhealthpython/general


What's missing?
    
    For actual production deployment, ideally an nginx ingress controller with an ingress pointing to the service with a cloudflare url pointing towards the load balancer(associated with the ingress) would be required. It is not complex, but it would be an unnecessary cost at the moment. I can spin it up if required too though, just let me know.


Answers to some of the questions in the interview doc:

Is the output correct? How do you verify?
  -> I need to write unit tests

How do you ensure changes don't break the application?
  -> ideally this would go through a dev, qa, staging then finally production to avoid breaking the application. 
  -> for unauth access, RBAC is useful. 

How might the application scale?
  -> using replicaCount. To ensure scalability, increasign the replica count on the pod will be enough. 

How can this application be deployed?
  -> I have deployed this via k8s, docker and helm. 
  -> It could however, be also deployed by serverless application model like aws lambda and api gateway

What if the application goes down? How will you know?
  -> I have created healthchecks in the deployment file. These k8s healthchecks can further be monitored by prometheus and grafana for performance. 
  -> A combination of Datadog and Pagerduty can also be used for alerting. 
