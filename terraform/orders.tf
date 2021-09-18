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
