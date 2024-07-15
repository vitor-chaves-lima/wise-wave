provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      "project"     = "WiseWave"
      "environment" = "Development"
    }
  }
}
