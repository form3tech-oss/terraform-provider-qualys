data "qualys_gcp_connector"  "my_test_connector" {
  connector_id = "<connector-uuid>"
}

output "my_test_connector_project_id" {
  value = data.qualys_gcp_connector.my_test_connector.project_id
}