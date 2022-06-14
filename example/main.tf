terraform {
  required_providers {
    briantest = {
      source  = "registry.cloudiac.org/brian/briantest"
      version = "0.1.2"
    }
  }
}

variable "template_address" {
  description = "GitHub address"
  type        = string
  default     = "https://github.com/brianchenl/terraform-provider-briantest"
}

resource "briantest_demo" "test" {
  instance_name = "brianTest"
  disk_size     = "108"
}

output "Template_address" {
  value = var.template_address
}
