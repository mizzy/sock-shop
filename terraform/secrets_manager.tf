resource "aws_secretsmanager_secret" "ssh_key" {
  name = "ssh-key"
}
