terraform {
  required_providers {
    terraform-provider-request = {
      source = "github.com/Pixxle/httpRequest"
      version = "0.0.1"
    }
  }
}

provider "terraform-provider-request" {}

data "httpRequest" "basic-http-example" {
  provider = terraform-provider-request

  url = "https://"
  request_headers = {
    content-type = "text/plain"
  }
  query_parameters = {
    hello = "world"
  }
}
