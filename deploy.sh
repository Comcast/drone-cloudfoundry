#!/bin/bash
#
################################################################
# Copyright 2016 Comcast Cable Communications Management, LLC *
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at *
# http://www.apache.org/licenses/LICENSE-2.0 *
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################

set -e

# check required vars and secrets
required=(PLUGIN_API PLUGIN_USER PLUGIN_PASSWORD PLUGIN_ORG PLUGIN_SPACE)

for arg in ${required[@]}
do
    eval [[ -n \$$arg ]] && continue
    eval [[ -z \$${arg/PLUGIN/CF} ]] && echo "${arg:7} is not defined!" && exit 1
    eval $arg=\$${arg/PLUGIN/CF}
done

# build api command
api_args=""
[[ -n $PLUGIN_SKIP_SSL ]] && [[ $PLUGIN_SKIP_SSL = true ]] && api_args="--skip-ssl-validation"

# set cloud foundry API
cf api $api_args $PLUGIN_API

# login to API
cf auth $PLUGIN_USER $PLUGIN_PASSWORD

# target org and space
cf target -o "$PLUGIN_ORG" -s "$PLUGIN_SPACE"

# build push command
push="cf push"

[[ -n $PLUGIN_NAME ]] && push="$push $PLUGIN_NAME"
[[ -n $PLUGIN_BUILDPACK ]] && push="$push -b $PLUGIN_BUILDPACK"
[[ -n $PLUGIN_COMMAND ]] && push="$push -c $PLUGIN_COMMAND"
[[ -n $PLUGIN_DOMAIN ]] && push="$push -d $PLUGIN_DOMAIN"
[[ -n $PLUGIN_MANIFEST ]] && push="$push -f $PLUGIN_MANIFEST"
[[ -n $PLUGIN_DOCKER_IMAGE ]] && push="$push -o $PLUGIN_DOCKER_IMAGE"
[[ -n $PLUGIN_INSTANCES ]] && push="$push -i $PLUGIN_INSTANCES"
[[ -n $PLUGIN_DISK ]] && push="$push -k $PLUGIN_DISK"
[[ -n $PLUGIN_MEMORY ]] && push="$push -m $PLUGIN_MEMORY"
[[ -n $PLUGIN_HOSTNAME ]] && push="$push -n $PLUGIN_HOSTNAME"
[[ -n $PLUGIN_PATH ]] && push="$push -p $PLUGIN_PATH"
[[ -n $PLUGIN_STACK ]] && push="$push -s $PLUGIN_STACK"
[[ -n $PLUGIN_TIMEOUT ]] && push="$push -t $PLUGIN_TIMEOUT"
[[ -n $PLUGIN_HEALTH_CHECK_TYPE ]] && push="$push -u $PLUGIN_HEALTH_CHECK_TYPE"
[[ -n $PLUGIN_ROUTE_PATH ]] && push="$push --route-path $PLUGIN_ROUTE_PATH"
[[ -n $PLUGIN_NO_HOSTNAME ]] && [[ $PLUGIN_NO_HOSTNAME = true ]] && push="$push --no-hostname"
[[ -n $PLUGIN_NO_MANIFEST ]] && [[ $PLUGIN_NO_MANIFEST = true ]] && push="$push --no-manifest"
[[ -n $PLUGIN_NO_ROUTE ]] && [[ $PLUGIN_NO_ROUTE = true ]] && push="$push --no-route"
[[ -n $PLUGIN_NO_START ]] && [[ $PLUGIN_NO_START = true ]] && push="$push --no-start"
[[ -n $PLUGIN_RANDOM_ROUTE ]] && [[ $PLUGIN_RANDOM_ROUTE = true ]] && push="$push --random-route"

# cf push
echo -e "\nExecuting:\n$push\n"
eval $push