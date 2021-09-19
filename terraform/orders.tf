resource "aws_dynamodb_table" "orders" {
  name           = "orders"
  hash_key       = "id"
  range_key      = "customerId"
  read_capacity  = 5
  write_capacity = 5

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "customerId"
    type = "S"
  }
}

module "ecs_orders" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-OrdersService-vKEiumEDxDGY"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "orders"
    image              = "weaveworksdemos/orders-aws"
    family             = "sock-shop-OrdersTask-XTi9XoWDVfLy"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    port               = 80
    memory             = 1024
    environment = [
      {
        name  = "AWS_DYNAMODB_ENDPOINT"
        value = "dynamodb.ap-northeast-1.amazonaws.com"
      },

    ]
    dockerLabels  = {}
    task_role_arn = aws_iam_role.dynamodb_task_role.arn
  }

  registry = {
    name         = "orders"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
