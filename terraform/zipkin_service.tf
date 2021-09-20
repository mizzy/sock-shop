resource "aws_lb_target_group" "zipkin_service" {
  vpc_id      = aws_vpc.sock_shop.id
  target_type = "ip"
  protocol    = "HTTP"
  port        = 9411

  health_check {
    path                = "/health"
    healthy_threshold   = 5
    unhealthy_threshold = 2
  }
}
