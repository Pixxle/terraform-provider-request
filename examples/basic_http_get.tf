terraform {
  required_providers {
    terraform-provider-request = {
      source = "github.com/Pixxle/request"
      version = "0.0.1"
    }
  }
}

provider "terraform-provider-request" {}

resource "request_myhttp" "basic-http-example" {
  url = "https://"
  request_headers = {
    content-type = "text/plain"
  }
}
