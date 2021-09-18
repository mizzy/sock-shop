# CloudFormation for Sock Shop

The JSON files in this directory is copied from [Microservices Demo: Sock Shop](https://microservices-demo.github.io/deployment/ecs.html) and [markfink-splunk/sock-shop: Deployments of the Weaveworks Sock Shop application instrumented with SignalFx.](https://github.com/markfink-splunk/sock-shop) .

## How to create a stack

Before creating a stak, create a key pair named "sock-shop".

Then run this command:

```
$ AWS_PROFILE=sock-shop-cloudformation aws cloudformation create-stack \
  --stack-name sock-shop \
  --template-body file://cfn-stack-app-only.yaml \
  --capabilities CAPABILITY_IAM \
  --timeout-in-minutes 60
```

## How to update a stack

```
$ AWS_PROFILE=sock-shop-cloudformation  aws cloudformation update-stack \
  --stack-name sock-shop \
  --template-body file://cfn-stack-app-only.yaml \
  --capabilities CAPABILITY_IAM
```
