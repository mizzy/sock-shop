resource "aws_service_discovery_service" "main" {
  count = var.registry == null ? 0 : 1

  name = var.registry.name

  dns_config {
    namespace_id = var.registry.namespace_id
    dns_records {
      ttl  = 10
      type = "A"
    }
  }

  health_check_custom_config {
    failure_threshold = 1
  }
}
