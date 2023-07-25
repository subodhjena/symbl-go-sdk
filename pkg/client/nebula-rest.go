// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package symbl

import (
	"context"
	"time"

	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
	rest "github.com/dvonthenen/symbl-go-sdk/pkg/client/rest"
)

// NewNebulaClientWithToken creates a new Nebula client.
// The client authenticates reusing an already valid Symbl Platform auth token
func NewNebulaClientWithToken(ctx context.Context, nebulaToken string) (*NebulaClient, error) {
	klog.V(6).Infof("NewRestClientWithToken ENTER\n")

	creds := interfaces.Credentials{
		Type: defaultAuthType,
	}
	resp := interfaces.AuthResp{
		NebulaToken: nebulaToken,
	}

	// if len(creds.AuthURI) > 0 {
	// 	klog.V(3).Infof("[OVERRIDE] AuthURI: %s\n", creds.AuthURI)
	// } else {
	// 	creds.AuthURI = defaultAuthURI
	// }

	// checks
	if ctx == nil {
		klog.V(3).Infof("Empty Context... Creating new one!\n")
		ctx = context.Background()
	}

	// validate input
	if resp.NebulaToken == "" {
		klog.V(1).Infof("Symbl Nebula Token is empty\n")
		klog.V(6).Infof("NewRestClientWithToken LEAVE\n")
		return nil, ErrInvalidInput
	}

	restClient := rest.New()
	restClient.SetAuthorization(&rest.AccessToken{
		NebulaToken: resp.NebulaToken,
		ExpiresOn:   time.Now().Add(time.Hour * 24),
	})

	c := &NebulaClient{
		Client: restClient,
		creds:  &creds,
		auth:   &resp,
	}

	klog.V(3).Infof("NewRestClientWithToken Succeeded\n")
	klog.V(6).Infof("NewRestClientWithToken LEAVE\n")
	return c, nil
}