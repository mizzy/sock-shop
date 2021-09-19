resource "aws_lb" "frontend" {
  name         = "sockshop"
  idle_timeout = 30
  subnets = [
    aws_subnet.public_subnet_1.id,
    aws_subnet.public_subnet_2.id,
  ]
  security_groups = [aws_security_group.elb_allowed_ports.id]
}

resource "aws_lb_target_group" "frontend" {
  vpc_id      = aws_vpc.sock_shop.id
  target_type = "ip"
  protocol    = "HTTP"
  port        = 8079
}
