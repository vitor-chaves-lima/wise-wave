locals {
  sender_email = "noreply@${local.domain_name}"
}

resource "aws_sesv2_configuration_set" "sender_configuration_set" {
  configuration_set_name = "WiseWaveEmailSender"

  reputation_options {
    reputation_metrics_enabled = true
  }
}

resource "aws_sesv2_email_identity" "sender_domain_identity" {
  configuration_set_name = aws_sesv2_configuration_set.sender_configuration_set.configuration_set_name
  email_identity         = local.domain_name
}

resource "aws_sesv2_email_identity" "sender_email_identity" {
  configuration_set_name = aws_sesv2_configuration_set.sender_configuration_set.configuration_set_name
  email_identity         = local.sender_email
}

resource "aws_ses_template" "first_access_magic_link" {
  name    = "FirstAccessMagicLink"
  subject = "Seja bem-vindo!"
  html    = <<-HTML
    <!DOCTYPE html>
    <html>
    <body>
        <p>Olá,</p>
        <p>Seu acesso à experiência WiseWave está pronto!</p>
        <p>Para começar, clique no link abaixo:</p>
        <p><a href="{{link}}">Magic Link</a></p>
    </body>
    </html>
  HTML
}

resource "aws_ses_template" "magic_link" {
  name    = "MagicLink"
  subject = "Seu link de acesso!"
  html    = <<-HTML
    <!DOCTYPE html>
    <html>
    <body>
        <p>Olá,</p>
        <p>Para acessar a experiência WiseWave clique no link abaixo:</p>
        <p><a href="{{link}}">Magic Link</a></p>
    </body>
    </html>
  HTML
}
