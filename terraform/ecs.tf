resource "aws_service_discovery_private_dns_namespace" "local" {
  name = "local"
  vpc  = aws_vpc.sock_shop.id
}
