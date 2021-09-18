resource "aws_vpc" "sock_shop" {
  cidr_block = "172.31.0.0/16"
}

resource "aws_vpc_dhcp_options" "local" {
  domain_name         = "local"
  domain_name_servers = ["AmazonProvidedDNS"]
}

resource "aws_vpc_dhcp_options_association" "local" {
  vpc_id          = aws_vpc.sock_shop.id
  dhcp_options_id = aws_vpc_dhcp_options.local.id
}
