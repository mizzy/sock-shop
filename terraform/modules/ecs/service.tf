resource "aws_ecs_service" "main" {
  name            = var.service.name
  desired_count   = 1
  launch_type     = "FARGATE"
  cluster         = var.service.cluster_id
  task_definition = "${aws_ecs_task_definition.main.id}:${aws_ecs_task_definition.main.revision}"

  network_configuration {
    assign_public_ip = true
    security_groups  = var.service.security_group_ids
    subnets          = var.service.subnet_ids
  }

  service_registries {
    registry_arn = var.service.registry_arn
  }
}
