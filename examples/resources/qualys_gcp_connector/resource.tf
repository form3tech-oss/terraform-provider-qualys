resource "qualys_gcp_connector"  "my_dev_connector" {
  name                = "dev_gcp"
  description         = "dev_gcp"
  gcp_config_file     = file("./service_account.json")
}