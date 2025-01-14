#!/bin/bash -e

# Copyright 2018- The Pixie Authors.
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
#
# SPDX-License-Identifier: Apache-2.0

workspace=$(bazel info workspace 2> /dev/null)
pushd "${workspace}" &> /dev/null || exit

function label_to_path() {
  path="${1#"//"}"
  echo "${path/://}"
}

function build() {
  # Exits with message if the bazel build command goes wrong.
  # Force bazel to download all targets since bazel will delete
  # targets with minimal download mode. This is a bug in bazel:
  #   https://github.com/bazelbuild/bazel/issues/12855
  if ! out=$(bazel build --remote_download_outputs=all "$@" 2>&1); then
    echo "${out}"
    exit 1
  fi
}

function copy() {
  for build_label in "$@"; do
    echo "Updating ${build_label} ..."

    srcPath=$(dirname "$(label_to_path "${build_label}")")
    # TODO(nick,PC-738): When moving the UI code to its own package (as opposed to the root), update this path.
    uiTypesDir="src/ui/src/types/generated"
    apiTypesDir="src/ui/packages/pixie-api/src/types/generated"

    # TODO(nick): pixie-api may not require all grpc_web protos; we should not copy over unused files.
    abs_paths=$(find "bazel-bin/${srcPath}" -iregex ".*/[^/]*_pb.*[tj]s")
    if [[ "${abs_paths}" == "" ]]; then
      echo "Failed to locate TypeScript and Javascript Proto files in bazel-bin/${srcPath}"
      return 1
    fi

    # VizierapiServiceClient.ts has a relative import; we're copying elsewhere. We fix this with perl string substitution.
    regexRelativeImport="s|import \* as ([^ ]+) from '([^ /]+/)+vizierapi_pb'\;|import * as \1 from './vizierapi_pb';|m"
    # vizierapi_pb.d.ts incorrectly includes an unused (and non-existent) relative import related to Go protos. Remove it.
    regexExtraneousImport="s|^import \* as github_com_gogo_protobuf_gogoproto_gogo_pb.*$||m"

    for abs_path in $abs_paths; do
      echo "Propagating ${abs_path} ..."
      fname=$(basename "$abs_path")
      cp -f "${abs_path}" "${uiTypesDir}"
      cp -f "${abs_path}" "${apiTypesDir}"

      # Using Perl instead of sed because BSD and GNU treat the -i flag differently: https://stackoverflow.com/a/22247781
      perl -pi -e "${regexRelativeImport}" "${uiTypesDir}/${fname}"
      perl -pi -e "${regexRelativeImport}" "${apiTypesDir}/${fname}"
      perl -pi -e "${regexExtraneousImport}" "${uiTypesDir}/${fname}"
      perl -pi -e "${regexExtraneousImport}" "${apiTypesDir}/${fname}"
    done
  done
}

if [[ $# == 0 ]]; then
  mapfile -t targets < <(bazel query --noshow_progress --noshow_loading_progress "kind('grpc_web_library rule', //...)")
else
  targets=("$@")
fi

build "${targets[@]}"
copy "${targets[@]}"
