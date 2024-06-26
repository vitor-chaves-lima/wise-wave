provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "example_bucket" {
  bucket = "example-bucket-name"

  tags = {
    Name = "Example Bucket"
  }
}
