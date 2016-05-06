provider "chronos" {
	url = "http://localhost:8081"
}

resource "chronos_job" "job-create-example" {
	name = "job_create_example"
	command = "sleep"
	shell = true
	epsilon = "PT30M"
	#executor = "mesos"
	#executor_flags = "--<mesos-flag>"
	retries = 3
	owner = "mail@address.tld"
	owner_name = "firstname surname"
	description = "simple example who run sleep for 5 sec"
	async = false
	cpus = 1
	disk = 128
	mem = 512
	disabled = false
	uris = [
		"https://test.com/testfile.jpg"
	]
	arguments = ["5"]
	high_priority = ""
	run_as_user = "mail"
	container = {
		type = "DOCKER"
		image = "debian:jessie"
		network = "HOST"
	}
	schedule = "R/2014-03-08T20:00:00.000Z/PT2H"
	schedule_timezone = "GMT"
# not implemented
#	constraints = {
#		constraint = {
#			attribute = "hostname"
#			operation = "GROUP_BY"
#			parameter = "1000"
#		}
#	}
}


resource "chronos_job" "docker-job-create-example" {
	name = "complex_dockerjob"
	command = "while sleep 10; do date -u +%T; done"
	shell = true
	epsilon = "PT60S"
	executor = ""
	executor_flags = ""
	retries = 2
	owner = ""
	async = false
	cpus = 0.5
	disk = 256
	mem = 512
	disabled = false
	uris = []
	schedule = "R/2015-05-21T18:14:00.000Z/PT2M"
	#env  = {}
	arguments = []
	run_as_user = "root"
	container = {
		type = "docker"
		image = "libmesos/ubuntu"
		network = "HOST"
		#volumes = {}
	}
}
