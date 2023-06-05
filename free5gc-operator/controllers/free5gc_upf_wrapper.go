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

var UPFWrapperScript string = `#!/bin/sh

### Implement networking rules
iptables -A FORWARD -j ACCEPT
iptables -t nat -A POSTROUTING -s $DNN_NETWORK -o $N6_INTERFACE_NAME -j MASQUERADE  # route traffic comming from the UE  SUBNET to the interface N6
echo "1200 n6if" >> /etc/iproute2/rt_tables # create a routing table for the interface N6
ip rule add from $DNN_NETWORK table n6if   # use the created ip table to route the traffic comming from  the UE SUBNET
ip route add default via $N6_GATEWAY dev $N6_INTERFACE_NAME table n6if  # add a default route in the created table so  that all UEs will use this gateway for external communications (target IP not in the Data Network attached  to the interface N6) and then the Data Network will manage to route the traffic

/free5gc/free5gc-upfd/free5gc-upfd -c /free5gc/config//upfcfg.yaml
`
