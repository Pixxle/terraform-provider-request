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

// If the response from http request has the content-type header set to application/json the structure gets unmarshalled
// and is available for use in the body computed field of the httpRequest datasource. In all other cases the terraform
// provider publishes the http response body in the body.value field.
output "myJsonOutput" {
  value = data.httpRequest.basic-http-example.body.about
}
