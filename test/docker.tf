resource "chronos_job" "docker-job-create-example" {
	name = "complex_dockerjob"
	command = "while sleep 10; do date -u +%T; done"
	shell = true
	epsilon = "PT60S"
	executor = ""
	executorFlags = ""
	retries = 2
	owner = ""
	async = false
	successCount = 190
	errorcount = 3
	lastsuccess = "2014-03-08T16:57:17.507Z"
	lasterror = "2014-03-01T00:10:15.957Z"
	cpus = 0.5
	disk = 256
	mem = 512
	disabled = false
	uris = []
	schedule = "R/2015-05-21T18:14:00.000Z/PT2M"
	env  = {}
	arguments = []
	runasuser = "root"
	container = {
		type = "docker"
		image = "libmesos/ubuntu"
		network = "HOST"
		###volumes = {}
	},
}
