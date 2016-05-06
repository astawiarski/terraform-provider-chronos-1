# must set this in env var or tfvar

variable "chronos_url" {}

provider "chronos" {
  url = "${var.chronos_url}"
}
