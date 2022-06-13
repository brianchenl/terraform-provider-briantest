terraform {
  required_providers {
    yunjidemo = {
      source  = "brian/briantest"
    }
  }
}

resource "brian_demo" "test" {
  instance_name  = "brianTest"
  disk_size = "108"
}
