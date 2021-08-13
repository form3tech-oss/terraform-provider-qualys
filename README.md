# terraform-provider-qualys

A Terraform provider for Qualys CloudView Resources

## References

Based on original work from here - [https://code.stanford.edu/xuwang/terraform-provider-qualys]

Qualys CloudView API Docs - [https://qualysguard.qg2.apps.qualys.com/cloudview-api/swagger-ui.html#/]

## Example

```shell
export QUALYS_URL=https://qualysguard.qg2.apps.qualys.com
export QUALYS_USERNAME=<your qualys login>
export QUALYS_PASSWORD=<your qualys password>
```

```hcl
provider "qualys" {
    // recommend to use environment variables to avoid exposing your credentials
    // base_url = "<your qualys platform url>"
    // username = "<your qualys login>"
    // password = "<your qualys password>"
}

data "qualys_gcp_connector"  "my_test_connector" {
    connector_id = "<connector-uuid>"
}

output "my_test_connector_project_id" {
    value = data.qualys_gcp_connector.my_test_connector.project_id
}

resource "qualys_gcp_connector"  "my_dev_connector" {
    name           = "dev_gcp"
    description    = "dev_gcp"
    config_file     = file("./dev_gcp.json")
}
```