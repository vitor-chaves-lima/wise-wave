resource "aws_ssm_parameter" "sender_identity" {
  name  = "/WiseWave/SenderIdentity"
  type  = "SecureString"
  value = <<EOF
{
  "arn": "${aws_sesv2_email_identity.sender_email_identity.arn}",
  "email": "${aws_sesv2_email_identity.sender_email_identity.email_identity}"
}
EOF
}
