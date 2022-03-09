resource "http" "basic-http-example" {
  url = "https://"
  request_headers = {
    content-type = "text/plain"
  }
}
