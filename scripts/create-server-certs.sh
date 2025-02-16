#!/usr/bin/env bash

# Copyright (c) Meta Platforms, Inc. and affiliates.
# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

# Create certs dir in the project's folder if it does not exists
mkdir -p certs

# Install mkcert trusty certificate in the OS
mkcert -ecdsa -install

# Generate public - private key 
go run filippo.io/mkcert -ecdsa -days 10 -cert-file "../certs/certificate.pem" -key-file "../certs/certificate.key" localhost 127.0.0.1 ::1
