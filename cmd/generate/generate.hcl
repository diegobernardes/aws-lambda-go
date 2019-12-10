handler {
  struct = "ALBTargetGroupRequest"
  output = ["events.ALBTargetGroupResponse", "error"]
}

handler {
  struct = "APIGatewayCustomAuthorizerRequest"
  output = ["events.APIGatewayCustomAuthorizerResponse", "error"]
}

handler {
  struct = "APIGatewayProxyRequest"
  output = ["events.APIGatewayProxyResponse", "error"]
}

handler {
  struct = "AppSyncResolverTemplate"
  output = ["error"]
}

handler {
  struct = "AutoScalingEvent"
}

handler {
  struct = "ChimeBotEvent"
  output = ["error"]
}

handler {
  struct = "CodeBuildEvent"
}

handler {
  struct = "CodeCommitEvent"
}

handler {
  struct = "CodeDeployEvent"
}

handler {
  struct = "CognitoEventUserPoolsVerifyAuthChallenge"
  output = ["events.CognitoEventUserPoolsVerifyAuthChallenge", "error"]
}

handler {
  struct = "CognitoEventUserPoolsPostConfirmation"
  output = ["events.CognitoEventUserPoolsPostConfirmation", "error"]
}

handler {
  struct = "CognitoEventUserPoolsPreAuthentication"
  output = ["events.CognitoEventUserPoolsPreAuthentication", "error"]
}

handler {
  struct = "CognitoEventUserPoolsPreSignup"
  output = ["events.CognitoEventUserPoolsPreSignup", "error"]
}

handler {
  struct = "CognitoEventUserPoolsPreTokenGen"
  output = ["events.CognitoEventUserPoolsPreTokenGen", "error"]
}

handler {
  struct = "CognitoEvent"
  output = ["error"]
}

handler {
  struct = "ConfigEvent"
}

handler {
  struct = "ConnectEvent"
  output = ["events.ConnectResponse", "error"]
}

handler {
  struct = "DynamoDBEvent"
}

handler {
  struct = "KinesisEvent"
  output = ["error"]
}

handler {
  struct = "KinesisAnalyticsOutputDeliveryEvent"
  output = ["events.KinesisAnalyticsOutputDeliveryResponse", "error"]
}

handler {
  struct = "KinesisFirehoseEvent"
  output = ["events.KinesisFirehoseResponse", "error"]
}

handler {
  struct = "LexEvent"
  output = ["events.LexResponse", "error"]
}

handler {
  struct = "S3BatchJobEvent"
  output = ["events.S3BatchJobResponse", "error"]
}

handler {
  struct = "S3Event"
}

handler {
  struct = "SimpleEmailEvent"
  output = ["error"]
}

handler {
  struct = "SNSEvent"
}

handler {
  struct = "SQSEvent"
  output = ["error"]
}
