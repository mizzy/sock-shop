resource "aws_lb_target_group" "zipkin" {
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

resource "aws_lb_listener" "zipkin" {
  load_balancer_arn = aws_lb.lb.arn
  port              = 9411
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.zipkin.arn
  }
}

resource "aws_lb_listener_rule" "zipkin" {
  listener_arn = aws_lb_listener.zipkin.arn
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.zipkin.arn
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

module "ecs_zipkin" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-ZipkinService-1eV4WLC8fHCo"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
    load_balancer = {
      container_name   = "zipkin"
      container_port   = 9411
      target_group_arn = aws_lb_target_group.zipkin.arn
    }
  }

  task = {
    name               = "zipkin"
    image              = "openzipkin/zipkin"
    family             = "sock-shop-ZipkinTask-xR5HsG5Dz7Qr"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    portMappings = [
      {
        containerPort = 9411,
        hostPort      = 9411,
      }
    ]
    environment = [
      {
        name  = "STORAGE_TYPE",
        value = "mysql",
      },
      {
        name  = "MYSQL_HOST",
        value = "zipkin-mysql",
      },
      {
        name  = "MYSQL_USER",
        value = "zipkin",
      },
      {
        name  = "MYSQL_PASS",
        value = "zipkin",
      }
    ]
    cpu    = 512
    memory = 1024
  }

  registry = {
    name         = "zipkin"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
