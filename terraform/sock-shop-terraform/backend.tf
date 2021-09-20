terraform {
  backend "s3" {
    region         = "ap-northeast-1"
    profile        = "management"
    dynamodb_table = "terraform"
    bucket         = "terraform.mizzy.org"
    key            = "sock-shop-terraform.tfstate"
    session_name   = "sock-shop-terraform"
    encrypt        = true
  }
}
