output "application_url" {
  value = "http://${aws_lb.lb.dns_name}"
}

output "zipkin_url" {
  value = "http://${aws_lb.lb.dns_name}:9411"
}
