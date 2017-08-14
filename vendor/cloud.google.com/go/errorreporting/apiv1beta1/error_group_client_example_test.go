// Copyright 2017, Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package errorreporting_test

import (
	"cloud.google.com/go/errorreporting/apiv1beta1"
	"golang.org/x/net/context"
	clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"
)

func ExampleNewErrorGroupClient() {
	ctx := context.Background()
	c, err := errorreporting.NewErrorGroupClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleErrorGroupClient_GetGroup() {
	ctx := context.Background()
	c, err := errorreporting.NewErrorGroupClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &clouderrorreportingpb.GetGroupRequest{
	// TODO: Fill request struct fields.
	}
	resp, err := c.GetGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleErrorGroupClient_UpdateGroup() {
	ctx := context.Background()
	c, err := errorreporting.NewErrorGroupClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &clouderrorreportingpb.UpdateGroupRequest{
	// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
