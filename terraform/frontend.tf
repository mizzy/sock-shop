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

resource "aws_lb_listener" "frontend" {
  load_balancer_arn = aws_lb.frontend.arn
  port              = 80
  protocol          = "HTTP"
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.frontend.arn
  }
}

resource "aws_lb_listener_rule" "frontend" {
  priority     = 1
  listener_arn = aws_lb_listener.frontend.arn
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.frontend.arn
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

module "ecs_frontend" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-FrontEndService-k3P0fqY2J0CD"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
    load_balancer = {
      container_name   = "front-end"
      container_port   = 8079
      target_group_arn = aws_lb_target_group.frontend.arn
    }
  }

  task = {
    name               = "front-end"
    image              = "weaveworksdemos/front-end"
    family             = "sock-shop-FrontEndTask-HW7tO3wK3fXn"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    port               = 8079
    memory             = 512
    command            = []
    environment = [
      {
        name  = "SESSION_REDIS"
        value = "true"
      },
    ]
    dockerLabels = {}
  }

  registry = {
    name         = "front-end"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
