resource "aws_ssm_parameter" "sender_email_identity" {
  name  = "/WiseWave/SenderEmailIdentity"
  type  = "SecureString"
  value = "Wise Wave <${aws_sesv2_email_identity.sender_email_identity.email_identity}>"
}
