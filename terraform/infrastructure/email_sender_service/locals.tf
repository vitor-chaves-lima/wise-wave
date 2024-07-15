locals {
  domain_name                  = "wisewave.tech"
  sender_email                 = "noreply@${local.domain_name}"
  email_sender_lambda_zip_path = "../../../backend/cmd/email_sender_lambda/build/email_sender_lambda.zip"
}
