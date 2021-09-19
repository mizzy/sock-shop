resource "aws_lb" "frontend" {
  name = "sockshop"
  idle_timeout = 30
  subnets = [
    aws_subnet.public_subnet_1.id,
    aws_subnet.public_subnet_2.id,
  ]
  security_groups = [aws_security_group.elb_allowed_ports.id]
}
