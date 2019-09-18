// Copyright © 2019 Banzai Cloud
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

package main

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"testing"
	"time"
)

var c = context.Background() // context.DeadlineExceeded()//  Background()

func TestVaultSecretsMutator(t *testing.T) {
	k8sClient, err := newK8SClient()
	if err != nil {
		t.Fatal(err.Error())
	}

	mutatingWebhook := mutatingWebhook{k8sClient: k8sClient}
	fubar := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:            name + "-vault-agent-config",
			OwnerReferences: ownerReferences,
		},
		Data: map[string]string{
			"config.hcl": fmt.Sprintf(vaultAgentConfig, vaultConfig.path, vaultConfig.role),
		},
	}

	fu, err := mutatingWebhook.vaultSecretsMutator(c)
}

func TestVaultTokenStore(t *testing.T) {
	os.Setenv("VAULT_ADDR", "http://localhost:8200")

	startTime := time.Now()
	ts := NewVaultTokenStore("")
	userID := "12"
	token := NewToken("id12", "my token")

	err := ts.Store(userID, token)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = ts.Store(userID, token)
	if err == nil {
		t.Fatal("Second Store for the same ids shouldn't work")
	}

	tokens, err := ts.List(userID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(tokens) != 1 {
		t.Fatal("There should be only one Token listed, but there are", len(tokens))
	}

	listedToken := tokens[0]

	if listedToken.ID != token.ID || listedToken.Name != token.Name || listedToken.CreatedAt.Before(startTime) {
		t.Fatalf("The listed token doesn't match: %#v", listedToken)
	}

	err = ts.Revoke(userID, token.ID)
	if err != nil {
		t.Fatal(err.Error())
	}

	expiredToken := NewToken("id13", "my expired token")
	now := time.Now().Add(-time.Minute)
	expiredToken.ExpiresAt = &now

	err = ts.Store(userID, expiredToken)
	if err != nil {
		t.Fatal(err.Error())
	}

	expiredToken, err = ts.Lookup(userID, "id13")
	if err != nil {
		t.Fatal(err.Error())
	}

	if expiredToken != nil {
		t.Fatalf("the expired token shouldn't be visible: %+v", expiredToken)
	}

	err = ts.Store(userID, token)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = ts.GC()
	if err != nil {
		t.Fatal(err.Error())
	}

	tokens, err = ts.List(userID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(tokens) != 1 {
		t.Fatal("There should be no expired Tokens listed, but there are", len(tokens))
	}
}
