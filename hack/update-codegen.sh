#!/usr/bin/env bash

# Copyright 2024 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# 	http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail
set -x

SCRIPT_ROOT=$(realpath $(dirname "${BASH_SOURCE[0]}"))
PROJECT_ROOT=$(realpath "${SCRIPT_ROOT}/..")

GENERATED_PKG="pkg/generated"

# Ensure tools built by kube_codegen go in our own GOBIN rather than something
# shared under GOPATH. This guards against these tools being rebuilt by some
# other concurrent invocation, potentially with a different version.
export GOBIN="${PROJECT_ROOT}/bin"

cd "$PROJECT_ROOT"

# For this to work, the current working directory must be under a Go module which
# lists k8s.io/code-generator
CODEGEN_PKG=$(go list -f '{{ .Dir }}' k8s.io/code-generator)

source "${CODEGEN_PKG}/kube_codegen.sh"

# Deep-copies and what-not are generated by controller-gen, so we don't need to use kube::codegen::gen_helpers

declare -a gen_openapi_args=(
    --report-filename "${PROJECT_ROOT}/api_violations.report"
    --output-dir "${PROJECT_ROOT}/cmd/models-schema"
    --output-pkg main
    --boilerplate "${SCRIPT_ROOT}/boilerplate.go.txt"

    # We need to include all referenced types in our generated openapi schema
    # or applyconfiguration-gen won't be able to use it. Helpfully it will
    # generate an error including the missing type.
    --extra-pkgs sigs.k8s.io/cluster-api/api/v1beta1
    --extra-pkgs k8s.io/api/core/v1
)

# It is an error to make a change which updates the api violations. When doing
# this intentionally, for example when fixing violations, run with
# UPDATE_API_KNOWN_VIOLATIONS=true to update the api violations report.
if [ ! -z "${UPDATE_API_KNOWN_VIOLATIONS:-}" ]; then
    gen_openapi_args+=(--update-report)
fi

kube::codegen::gen_openapi "${gen_openapi_args[@]}" "${PROJECT_ROOT}/api"

openapi="${PROJECT_ROOT}/openapi.json"
go run "${PROJECT_ROOT}/cmd/models-schema" | jq > "$openapi"

kube::codegen::gen_client \
    --with-applyconfig \
    --with-watch \
    --applyconfig-openapi-schema "$openapi" \
    --output-dir "${PROJECT_ROOT}/${GENERATED_PKG}" \
    --output-pkg sigs.k8s.io/cluster-api-provider-openstack/${GENERATED_PKG} \
    --versioned-name clientset \
    --boilerplate "${SCRIPT_ROOT}/boilerplate.go.txt" \
    --one-input-api "api" \
    "${PROJECT_ROOT}"
