  terraform {
    required_providers {
      briantest = {
        source = "registry.cloudiac.org/brian/briantest"
        version = "0.1.0"
      }
    }
  }

resource "brian_demo" "test" {
  instance_name  = "brianTest"
  disk_size = "108"
}
