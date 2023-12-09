#!/usr/bin/env bash

# Copyright (c) Meta Platforms, Inc. and affiliates.
# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

# Create certs dir if it does not exists
mkdir -p ../certs

# Install mkcert trusty certificate in the OS
mkcert -ecdsa -install

# Generate public - private key 
mkcert -ecdsa -cert-file "../certs/certificate.pem" -key-file "../certs/certificate.key" localhost 127.0.0.1 ::1
