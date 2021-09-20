resource "aws_security_group" "elb_allowed_ports" {
  description = "ELB Allowed Ports"
  tags = {
    "Name" = "elb-allowed-ports"
  }
}

resource "aws_security_group_rule" "elb_allowed_port_http" {
  type              = "ingress"
  security_group_id = aws_security_group.elb_allowed_ports.id
  protocol          = "tcp"
  from_port         = 80
  to_port           = 80
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "elb_allowed_port_zipkin" {
  type              = "ingress"
  security_group_id = aws_security_group.elb_allowed_ports.id
  protocol          = "tcp"
  from_port         = 9411
  to_port           = 9411
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "elb_allow_to_all" {
  type              = "egress"
  security_group_id = aws_security_group.elb_allowed_ports.id
  from_port         = 0
  to_port           = 65536
  protocol          = -1
  cidr_blocks       = ["0.0.0.0/0"]
}
