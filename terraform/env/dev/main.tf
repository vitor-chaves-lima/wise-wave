provider "aws" {
  region = var.aws_region
}

module "base" {
  source = "../../base"

  aws_region = var.aws_region
}
