// Copyright 2023 The Nephio Authors
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
//
// SPDX-License-Identifier: Apache-2.0

/*

 */

package controllers

var UPFCfg string = `
info:
  version: 1.0.0
  description: UPF configuration

configuration:
  ReportCaller: false
  debugLevel: info
  dnn_list:
  - cidr: $DNN_CIDR
    dnn: $DNN
    natifname: n6

  pfcp:
    - addr: $PFCP_IP

  gtpu:
    - addr: $GTPU_IP
    # [optional] gtpu.name
    # - name: upf.5gc.nctu.me
    # [optional] gtpu.ifname
    # - ifname: gtpif
`
