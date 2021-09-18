resource "aws_security_group" "rds_importer_allowed_ports" {
  # This description is from the original cloudformation template, but it seems not appropriate
  description = "ECS Allowed Ports"
  tags = {
    "Name" = "rds-importer-allowed-ports"
  }
}

resource "aws_security_group_rule" "rds_importer_allowed_port_ssh" {
  type              = "ingress"
  security_group_id = aws_security_group.rds_importer_allowed_ports.id
  protocol          = "tcp"
  from_port         = 22
  to_port           = 22
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "rds_importer_allowed_from_elb" {
  type                     = "ingress"
  security_group_id        = aws_security_group.rds_importer_allowed_ports.id
  protocol                 = -1
  from_port                = 0
  to_port                  = 65536
  source_security_group_id = aws_security_group.elb_allowed_ports.id
}

resource "aws_security_group_rule" "rds_importer_allowed_from_self" {
  type                     = "ingress"
  security_group_id        = aws_security_group.rds_importer_allowed_ports.id
  protocol                 = -1
  from_port                = 0
  to_port                  = 65536
  source_security_group_id = aws_security_group.rds_importer_allowed_ports.id
}
