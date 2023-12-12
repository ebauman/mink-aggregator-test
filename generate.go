//go:generate go run github.com/acorn-io/baaah/cmd/deepcopy ./pkg/apis/badapple.com/v1
//go:generate go run github.com/acorn-io/baaah/cmd/deepcopy ./pkg/apis/example.com/v1
//go:generate go run k8s.io/kube-openapi/cmd/openapi-gen -i github.com/ebauman/mink-aggregator-test/pkg/apis/example.com/v1,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version,k8s.io/apimachinery/pkg/api/resource,k8s.io/api/core/v1,k8s.io/api/rbac/v1,k8s.io/apimachinery/pkg/util/intstr -o ./  -p /pkg/openapi/example_com -h hack/boilerplate.go.txt
//go:generate go run k8s.io/kube-openapi/cmd/openapi-gen -i github.com/ebauman/mink-aggregator-test/pkg/apis/badapple.com/v1,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version,k8s.io/apimachinery/pkg/api/resource,k8s.io/api/core/v1,k8s.io/api/rbac/v1,k8s.io/apimachinery/pkg/util/intstr -o ./  -p /pkg/openapi/badapple_com -h hack/boilerplate.go.txt
package main

import (
	_ "github.com/acorn-io/baaah/pkg/deepcopy"
	_ "k8s.io/kube-openapi/cmd/openapi-gen/args"
)
