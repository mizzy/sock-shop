data "aws_iam_policy_document" "ecs_tasks_assume_role_policy" {
  version = "2008-10-17"
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      identifiers = ["ecs-tasks.amazonaws.com"]
      type        = "Service"
    }
  }
}

resource "aws_iam_role" "dynamodb_task_role" {
  name               = "sock-shop-DynamoDbTaskRole-13YK50YRC8S9F"
  assume_role_policy = data.aws_iam_policy_document.ecs_tasks_assume_role_policy.json
}

resource "aws_iam_role_policy_attachment" "dynamodb_task_role-amazon_dynamodb_full_access" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
  role       = aws_iam_role.dynamodb_task_role.name
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name               = "sock-shop-EcsTaskExecutionRole-4VE06B3BDH02"
  assume_role_policy = data.aws_iam_policy_document.ecs_tasks_assume_role_policy.json
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_role-amazon_ecs_task_execution_role_policy" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
  role       = aws_iam_role.ecs_task_execution_role.name
}

resource "aws_service_discovery_private_dns_namespace" "local" {
  name = "local"
  vpc  = aws_vpc.sock_shop.id
}
