// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func Test_extractManagedEnvVars(t *testing.T) {
	type args struct {
		container *corev1.Container
	}
	tests := []struct {
		name string
		args args
		want []corev1.EnvVar
	}{
		{"When there is a managed key",
			args{container: &corev1.Container{Env: []corev1.EnvVar{
				{Name: infinispanEnvKeyUsername, Value: "username"},
				{Name: "key1", Value: "value1"},
			}}},
			[]corev1.EnvVar{{Name: infinispanEnvKeyUsername, Value: "username"}}},
		{"When there is no managed key",
			args{container: &corev1.Container{Env: []corev1.EnvVar{
				{Name: "key1", Value: "value1"},
			}}},
			[]corev1.EnvVar{}},
		{"When there is no key",
			args{container: &corev1.Container{Env: []corev1.EnvVar{}}},
			[]corev1.EnvVar{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractManagedEnvVars(tt.args.container); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractManagedEnvVars() = %v, want %v", got, tt.want)
			}
		})
	}
}
