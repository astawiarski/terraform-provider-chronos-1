# Chronos Terraform Provider

## Install
```
$ go get 
```

## Usage

### Provider Configuration
Use a [tfvar file](https://www.terraform.io/intro/getting-started/variables.html) or set the ENV variable

```bash
$ export TF_VAR_chronos_url="http://chronos.domain.tld:4040"
```

```hcl
variable "chronos_url" {}

provider "chronos" {
  url = "${var.chronos_url}"
}
```

```bash
$ export TF_VAR_chronos_url="https://chronos.domain.tld:8443"
$ export TF_VAR_chronos_debug="true"
$ export TF_VAR_chronos_timeout="5"

```

```hcl
variable "chronos_url" {}
variable "chronos_debug" {}
variable "chronos_timeout" {}

provider "chronos" {
  url = "${var.chronos_url}"
  debug = "${var.chronos_debug}"
  request_timeout = "${var.chronos_timeout}"
}
```

### Basic Usage
```hcl
resource "chronos_job" "hello-world" {
  name = "Hello World"
  command = "echo Hello World"
  owner = "mail@address.tld"
  owner_name = "firstname surname"
  description = "simple example who display Hello World"
  schedule = "R/2014-03-08T20:00:00.000Z/PT2H"
  schedule_timezone = "GMT"
}
```

### Docker Usage
```hcl
resource "chronos_job" "docker-hello-world" {
  name = "Hello Docker"
  command = "echo Hello Docker"
  owner = "mail@address.tld"
  owner_name = "firstname surname"
  description = "simple example who display Hello Docker"
  container = {
  	type = "DOCKER"
  	image = "debian:jessie"
  	network = "HOST"
  }
  schedule = "R/2014-03-08T20:00:00.000Z/PT2H"
  schedule_timezone = "GMT"
}
```

## Development

### Build
```bash
$ go install
```

### Test
```bash
$ export CHRONOS_URL="http://chronos.domain.tld:8080"
$ ./test.sh
```
