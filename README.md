# CockroachDB Cloud Sample Application

This is a sample application demonstrating the [Service Binding](https://github.com/servicebinding/spec) feature in OpenShift.
The application shows the minimal CRUD service by exposing a set of endpoints over REST, and the Angular based front-end allows user to interact with the application/endpoints.

Under the hood:
* Go Gorilla to expose REST endpoints
* pgx library to connect to CockroachDB Cloud
* [Service Binding Client](https://github.com/RHEcosystemAppEng/sbo-go-library) to load the connection properties and returns the Connection string
* Compatible with [DBaaS](https://github.com/RHEcosystemAppEng/dbaas-operator)


### Run locally

#### Prerequisite

* Go 1.17 or above
* Setup a [CockroachDB Cluster](https://www.cockroachlabs.com/get-started-cockroachdb/), see `Initialize database`

```shell
# Update your cluster info. to the properties under ./test/bindings

# SERVICE_BINDING_ROOT should set to an full path to the service binding folder
SERVICE_BINDING_ROOT=$ABSOLUTE_PATH/test/bindings go run ./cmd/main.go

# Use a browser to access the front-end: http://localhost:8080 
```

### Run in OpenShift

#### Prerequisite

* Go 1.17 or above
* Docker
* OpenShift 4.9 or above
* Setup the [DBaaS](https://github.com/RHEcosystemAppEng/dbaas-operator) in OpenShift and use CockroachDB Could as the database provider.

```shell
# build and push the application to an image registry
$ IMAGE_TAG=<any_image_tag> make docker-build docker-push

# modify ./deploy-crdb-app.yaml to setup your image path

# deploy the application to OpenShift
$ oc apply -f ./deploy-crdb-app.yaml

# follow DBaaS documentation to setup the Service Binding
```

### Initialize database

This application targets [CockroachDB](https://www.cockroachlabs.com/get-started-cockroachdb/) (other PostgreSQL compatible database should also work).

Once a CockroachDB cluster is created. 
During the startup, this sample application will create the `fruit` table with few initial rows automatically. 

In case you want to run your own SQL, you can run the following command with your specific settings:
```yaml
e.g.:
$ cat <some-sql-file> | cockroach sql --url 'postgresql://<username>:<password>@<serverless-host>:26257/defaultdb?sslmode=verify-full&sslrootcert='$HOME'/.postgresql/root.crt&options=--cluster=<routing-id>'

```

### CockroachDB Getting Started
* [Quickstart with CockroachDB Serverless](https://www.cockroachlabs.com/docs/cockroachcloud/quickstart.html)
* [Connect to CockroachDB cluster](https://www.cockroachlabs.com/docs/cockroachcloud/connect-to-a-serverless-cluster)
