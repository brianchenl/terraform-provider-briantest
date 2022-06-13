  terraform {
    required_providers {
      briantest = {
        source = "registry.cloudiac.org/brian/briantest"
        version = "0.1.1"
      }
    }
  }

resource "briantest_demo" "test" {
  instance_name  = "brianTest"
  disk_size = "108"
}
