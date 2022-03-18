terraform {
  required_providers {
    terraform-provider-request = {
      source = "github.com/Pixxle/request"
      version = "0.0.1"
    }
  }
}

provider "terraform-provider-request" {}

data "myhttp" "basic-http-example" {
  provider = terraform-provider-request

  url = "https://"
  request_headers = {
    content-type = "text/plain"
  }
  query_parameters = {
    hello = world
  }
}
