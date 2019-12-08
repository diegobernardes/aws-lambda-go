# rename from event to handler.

event {
  struct = "ALBTargetGroupRequest"
  output = ["events.ALBTargetGroupResponse", "error"]
}

event {
  struct = "APIGatewayCustomAuthorizerRequest"
  output = ["events.APIGatewayCustomAuthorizerResponse", "error"]
}

event {
  struct = "APIGatewayProxyRequest"
  output = ["events.APIGatewayProxyResponse", "error"]
}

event {
  struct = "AppSyncResolverTemplate"
  output = ["error"]
}

event {
  struct = "AutoScalingEvent"
}

event {
  struct = "ChimeBotEvent"
  output = ["error"]
}

event {
  struct = "CodeBuildEvent"
}

event {
  struct = "CodeCommitEvent"
}

event {
  struct = "CodeDeployEvent"
}

event {
  struct = "CognitoEventUserPoolsVerifyAuthChallenge"
  output = ["events.CognitoEventUserPoolsVerifyAuthChallenge", "error"]
}

event {
  struct = "CognitoEventUserPoolsPostConfirmation"
  output = ["events.CognitoEventUserPoolsPostConfirmation", "error"]
}

event {
  struct = "CognitoEventUserPoolsPreAuthentication"
  output = ["events.CognitoEventUserPoolsPreAuthentication", "error"]
}

event {
  struct = "CognitoEventUserPoolsPreSignup"
  output = ["events.CognitoEventUserPoolsPreSignup", "error"]
}

event {
  struct = "CognitoEventUserPoolsPreTokenGen"
  output = ["events.CognitoEventUserPoolsPreTokenGen", "error"]
}

event {
  struct = "CognitoEvent"
  output = ["error"]
}

event {
  struct = "ConfigEvent"
}

event {
  struct = "ConnectEvent"
  output = ["events.ConnectResponse", "error"]
}

event {
  struct = "DynamoDBEvent"
}

event {
  struct = "KinesisEvent"
  output = ["error"]
}

event {
  struct = "KinesisAnalyticsOutputDeliveryEvent"
  output = ["events.KinesisAnalyticsOutputDeliveryResponse", "error"]
}

event {
  struct = "KinesisFirehoseEvent"
  output = ["events.KinesisFirehoseResponse", "error"]
}

event {
  struct = "LexEvent"
  output = ["events.LexResponse", "error"]
}

event {
  struct = "S3BatchJobEvent"
  output = ["events.S3BatchJobResponse", "error"]
}

event {
  struct = "S3Event"
}

event {
  struct = "SimpleEmailEvent"
  output = ["error"]
}

event {
  struct = "SNSEvent"
}

event {
  struct = "SQSEvent"
  output = ["error"]
}