locals {
  domain_name                  = "wisewave.tech"
  sender_email                 = "noreply@${local.domain_name}"
  email_sender_lambda_zip_path = "../../../backend/email_sender_service/build/sqs_consumer_lambda.zip"
}
