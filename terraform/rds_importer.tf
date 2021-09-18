resource "aws_security_group" "rds_importer_allowed_ports" {
  # This description is from the original cloudformation template, but it seems not appropriate
  description = "ECS Allowed Ports"
  tags = {
    "Name" = "rds-importer-allowed-ports"
  }
}
