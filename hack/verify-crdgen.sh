#!/usr/bin/env bash
#
# Copyright 2021 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${SCRIPT_ROOT}/hack/lib/init.sh"

kube::golang::verify_go_version

CRD_OPTIONS="crd:generateEmbeddedObjectMeta=true"

# Download controller-gen locally
CONTROLLER_GEN="${GOPATH}/bin/controller-gen"
go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.14.0

# Generate CRD
# api_paths="./pkg/apis/alert/v1alpha1/...;./pkg/apis/rule/v1alpha1/...;./pkg/apis/template/v1alpha1/...;./pkg/apis/diagnose/v1alpha1/..."
api_paths="./pkg/apis/diagnosis/v1alpha1/..."
${CONTROLLER_GEN} ${CRD_OPTIONS} paths="${api_paths}" output:dir="./manifests/install"

if ! _out="$(git --no-pager diff -I"edited\smanually" --exit-code ./manifests)"; then
    echo "Generated output differs" >&2
    echo "${_out}" >&2
    echo "Verification for CRD generators failed."
    exit 1
fi

echo "Controllers Gen for CRD verified."
