// Code generated by go generate; DO NOT EDIT.
// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package lambda

import (
	"context"
  "encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda/handlertrace"
)

// NewHandlerConnectEvent is used to process ConnectEvent types.
func NewHandlerConnectEvent(handler func(context.Context, events.ConnectEvent) (events.ConnectResponse, error)) Handler {
	return lambdaHandler(func(ctx context.Context, payload []byte) (interface{}, error) {
		trace := handlertrace.FromContext(ctx)

		var e events.ConnectEvent
		if err := json.Unmarshal(payload, &e); err != nil {
			panic(err)
		}

		if trace.RequestEvent != nil {
			trace.RequestEvent(ctx, e)
		}
		
    response, err := handler(ctx, e)
    if trace.ResponseEvent != nil {
      trace.ResponseEvent(ctx, response)
    }

    return response, err
	})
}