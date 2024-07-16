provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      "project"     = "WiseWave"
      "environment" = "Development"
    }
  }
}

module "infrastructure" {
  source = "../../infrastructure"

  aws_region = var.aws_region
}
