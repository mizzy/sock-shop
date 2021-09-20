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

resource "aws_lb_listener" "zipkin_service" {
  load_balancer_arn = aws_lb.lb.arn
  port = 9411
  default_action {
    type = "forward"
    target_group_arn = aws_lb_target_group.zipkin_service.arn
  }
}

resource "aws_lb_listener_rule" "zipkin_service" {
  listener_arn = aws_lb_listener.zipkin_service.arn
  action {
    type = "forward"
    target_group_arn = aws_lb_target_group.zipkin_service.arn
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}
