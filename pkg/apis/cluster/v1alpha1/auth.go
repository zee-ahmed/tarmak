// Copyright Jetstack Ltd. See LICENSE for details.
// Copyright © 2017 The Kubicorn Authors
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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SSH struct {
	metav1.TypeMeta      `json:",inline"`
	metav1.ObjectMeta    `json:"metadata,omitempty"`
	User                 string `json:"user,omitempty"`
	Identifier           string `json:"identifier,omitempty"`
	PublicKeyPath        string `json:"publicKeyPath,omitempty"`
	PublicKeyData        []byte `json:"publicKeyData,omitempty"`
	PublicKeyFingerprint string `json:"publicKeyFingerprint,omitempty"`
	PrivateKeyPath       string `json:"privateKeyPath,omitempty"`
}
