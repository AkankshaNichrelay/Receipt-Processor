# Receipt Processor

A mock webservice written in Golang that reads a receipt and calculates reward points.

## Instructions to run application locally

### Running Locally without Taskfile
* Run docker unit tests
  * ``` docker run -p 8080:8080 receipt-processor-build go test ./... ```
* Build docker image
  * ``` docker build -t receipt-processor-build . ```
* Run docker container
  * ``` docker run -p 8080:8080 receipt-processor-build ```

### Running Locally using Taskfile

This services uses [Taskfile](./Taskfile.yml) for ease in running various commands.

* To install for Mac users,
Install Homebrew:
```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Then install task using brew:
```
brew install go-task/tap/go-task
```

* For other OS/ package managers, follow steps for your package manager provided here: https://taskfile.dev/installation/

* Run Docker unit tests
    * All packages
      *  ```task docker-test```
    * A single package
      * ```task docker-test PKG=[path to package]```
* Build Docker image
    * ```task docker-build```
* Run Docker container
    * ```task docker-run```

* Postman (optional. Would need to create postman account)
    * Click on the button below and click on `View collection` to run in web browser. You can also `import a copy` or `Fork Collection`.
    * Set the "receipt-processor_url" to the local url of host. Example: localhost:8080

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/20550046-60ceb700-a977-4c73-a11b-27e0bdb8f04f?action=collection%2Ffork&collection-url=entityId%3D20550046-60ceb700-a977-4c73-a11b-27e0bdb8f04f%26entityType%3Dcollection%26workspaceId%3De2af4782-e8cd-4f79-ac8a-1d5130679b9a#?env%5Blocal%5D=W3sia2V5IjoicmVjZWlwdC1wcm9jZXNzb3JfdXJsIiwidmFsdWUiOiJsb2NhbGhvc3Q6ODA4MCIsImVuYWJsZWQiOnRydWUsInR5cGUiOiJkZWZhdWx0Iiwic2Vzc2lvblZhbHVlIjoibG9jYWxob3N0OjgwODAiLCJzZXNzaW9uSW5kZXgiOjB9XQ==)

