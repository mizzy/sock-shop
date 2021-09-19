migration "state" "mv" {
  actions = [
    "mv aws_ecs_task_definition.carts_db module.ecs_carts_db.aws_ecs_task_definition.main",
    "mv aws_ecs_service.carts_db module.ecs_carts_db.aws_ecs_service.main",
    "mv aws_service_discovery_service.carts_db module.ecs_carts_db.aws_service_discovery_service.main",
  ]
}
