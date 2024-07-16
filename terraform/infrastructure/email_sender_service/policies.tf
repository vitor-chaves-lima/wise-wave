data "aws_iam_policy_document" "send_email_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "ses:SendTemplatedEmail",
    ]

    resources = ["*"]

    condition {
      test     = "StringLike"
      variable = "ses:FromAddress"
      values   = [aws_sesv2_email_identity.sender_email_identity.email_identity]
    }

    effect = "Allow"
  }
}

resource "aws_iam_policy" "send_email_policy" {
  name_prefix = "WiseWaveSendEmail"
  policy      = data.aws_iam_policy_document.send_email_policy.json
}

data "aws_iam_policy_document" "send_logs_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]

    resources = ["*"]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "send_logs_policy" {
  name_prefix = "WiseWaveSendLogs"
  policy      = data.aws_iam_policy_document.send_logs_policy.json
}

data "aws_iam_policy_document" "email_sender_queue_receive_message_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "sqs:ReceiveMessage",
      "sqs:DeleteMessage",
      "sqs:GetQueueAttributes"
    ]

    resources = [aws_sqs_queue.email_sender_queue.arn]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "email_sender_queue_receive_message_policy" {
  name_prefix = "WiseWaveEmailSenderQueueReceiveMessage"
  policy      = data.aws_iam_policy_document.email_sender_queue_receive_message_policy.json
}

data "aws_iam_policy_document" "email_sender_parameter_read_policy" {
  statement {
    actions = [
      "ssm:GetParameter"
    ]

    resources = [
      aws_ssm_parameter.sender_identity.arn
    ]
  }
}

resource "aws_iam_policy" "email_sender_parameter_read_policy" {
  name_prefix = "WiseWaveEmailSenderParameterRead"
  policy      = data.aws_iam_policy_document.email_sender_parameter_read_policy.json
}
