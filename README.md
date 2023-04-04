# telemetry-operator
The Telemetry Operator handles the deployment of all the needed agents for gathering telemetry to assess the full state of a running Openstack cluster.

## Description
This operator deploys a pod containing the number of pods needed to retrieve information from the APIs exposed by the Openstack services.

The pod currently runs three containers:

- ceilometer-central-agent: This component is mainly configured by the polling.yaml file and it polls some service APIs (defined in the configuration file) every fixed amount of time. When it gets the information, it publishes it to RabbitMQ for consumptions

- ceilometer-notification-agent: This component suscribes to RabbitMQ and consumes the metric events generated by the central agent. One consumed, it sends the metrics via a tcp_socket to sg-core.

- sg-core: This component exposes a tcp_socket that the notification agent can use to feed it metrics data. Then, it converts the data to Prometheus format and exposes it via a REST endpoint that Prometheus is able to scrape.

The result is that the metrics that are polled from the Openstack services APIs are exposed to an external Prometheus to scrape.

### Running on the cluster
1. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/telemetry-operator:tag
```

2. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/telemetry-operator:tag
```

NOTE: please note that using the usual Operator development cycle based on ./bin/manager does not work because that does not bundle the required templates into the container. Only docker-build does that.

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### Local pre-commit install

MacOS installation:

```sh
brew install kube-linter
brew install kubebuilder
setup-envtest use -p env
# and add output to your ~/.bashrc
brew install pre-commit
pre-commit install
```

EL linux installation:

```
#TODO
```

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
