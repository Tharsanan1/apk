/*
 *  Copyright (c) 2024, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  This file contains code derived from Envoy Gateway,
 *  https://github.com/envoyproxy/gateway
 *  and is provided here subject to the following:
 *  Copyright Project Envoy Gateway Authors
 *
 */

package gatewayapi

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/wso2/apk/adapter/internal/operator/gateway-api/crypto"
	"github.com/wso2/apk/adapter/internal/operator/utils"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ErrSecretExists = errors.New("skipped creating secret since it already exists")
)

// caCertificateKey is the key name for accessing TLS CA certificate bundles
// in Kubernetes Secrets.
const (
	caCertificateKey = "ca.crt"
	hmacSecretKey    = "hmac-secret"
)

func newSecret(secretType corev1.SecretType, name string, namespace string, data map[string][]byte) corev1.Secret {
	return corev1.Secret{
		Type: secretType,
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				"control-plane": "envoy-gateway",
			},
		},
		Data: data,
	}
}

// CertsToSecret creates secrets in the provided namespace, in compact form, from the provided certs.
func CertsToSecret(namespace string, certs *crypto.Certificates) []corev1.Secret {
	return []corev1.Secret{
		newSecret(
			corev1.SecretTypeTLS,
			"envoy-gateway",
			namespace,
			map[string][]byte{
				caCertificateKey:        certs.CACertificate,
				corev1.TLSCertKey:       certs.EnvoyGatewayCertificate,
				corev1.TLSPrivateKeyKey: certs.EnvoyGatewayPrivateKey,
			}),
		newSecret(
			corev1.SecretTypeTLS,
			"envoy",
			namespace,
			map[string][]byte{
				caCertificateKey:        certs.CACertificate,
				corev1.TLSCertKey:       certs.EnvoyCertificate,
				corev1.TLSPrivateKeyKey: certs.EnvoyPrivateKey,
			}),
		newSecret(
			corev1.SecretTypeTLS,
			"envoy-rate-limit",
			namespace,
			map[string][]byte{
				caCertificateKey:        certs.CACertificate,
				corev1.TLSCertKey:       certs.EnvoyRateLimitCertificate,
				corev1.TLSPrivateKeyKey: certs.EnvoyRateLimitPrivateKey,
			}),
		newSecret(
			corev1.SecretTypeOpaque,
			"envoy-oidc-hmac",
			namespace,
			map[string][]byte{
				hmacSecretKey: certs.OIDCHMACSecret,
			}),
	}
}

// CreateOrUpdateSecrets creates the provided secrets if they don't exist or updates
// them if they do.
func CreateOrUpdateSecrets(ctx context.Context, client client.Client, secrets []corev1.Secret, update bool) ([]corev1.Secret, error) {
	var (
		tidySecrets     []corev1.Secret
		existingSecrets []string
	)

	for secret := range secrets {
		secret := secrets[secret]
		current := new(corev1.Secret)
		key := utils.NamespacedName(&secret)
		if err := client.Get(ctx, key, current); err != nil {
			// Create if not found.
			if kerrors.IsNotFound(err) {
				if err := client.Create(ctx, &secret); err != nil {
					return nil, fmt.Errorf("failed to create secret %s/%s: %w", secret.Namespace, secret.Name, err)
				}
			} else {
				return nil, fmt.Errorf("failed to get secret %s/%s: %w", secret.Namespace, secret.Name, err)
			}
			// Update if current value is different and update arg is set.
		} else {
			if !update {
				existingSecrets = append(existingSecrets, fmt.Sprintf("%s/%s", secret.Namespace, secret.Name))
				continue
			}
			fmt.Println()

			if !reflect.DeepEqual(secret.Data, current.Data) {
				if err := client.Update(ctx, &secret); err != nil {
					return nil, fmt.Errorf("failed to update secret %s/%s: %w", secret.Namespace, secret.Name, err)
				}
			}
		}
		tidySecrets = append(tidySecrets, secret)
	}

	if len(existingSecrets) > 0 {
		return tidySecrets, fmt.Errorf("%v: %w;"+
			"Either update the secrets manually or set overwriteControlPlaneCerts "+
			"in the EnvoyGateway config", existingSecrets, ErrSecretExists)
	}

	return tidySecrets, nil
}
