// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeMetadataRespV9(t *testing.T) {
	protocolConfig := KafkaProtocolConfig{}
	protocolConfig.ClusterId = "shoothzj"
	protocolConfig.AdvertiseHost = "localhost"
	protocolConfig.AdvertisePort = 9092
	protocolConfig.NodeId = 0
	metadataResp := NewMetadataResp(2, &protocolConfig, "764edee3-007e-48e0-b9f9-df7f713ff707", 0)
	bytes := metadataResp.Bytes(9)
	expectBytes := testHex2Bytes(t, "00000002000000000002000000000a6c6f63616c686f73740000238400000973686f6f74687a6a000000000200002537363465646565332d303037652d343865302d623966392d6466376637313366663730370002000000000000000000000000000002000000000200000000010080000000008000000000000000")
	assert.Equal(t, expectBytes, bytes)
}

func TestCodeMetadataRespV1(t *testing.T) {
	protocolConfig := KafkaProtocolConfig{}
	protocolConfig.NodeId = 1
	protocolConfig.AdvertiseHost = "localhost"
	protocolConfig.AdvertisePort = 9092
	metadataResp := NewMetadataResp(2, &protocolConfig, "topic", 0)
	bytes := metadataResp.Bytes(1)
	expectBytes := testHex2Bytes(t, "00000002000000010000000100096c6f63616c686f737400002384ffff000000010000000100000005746f70696300000000010000000000000000000100000001000000010000000100000001")
	assert.Equal(t, expectBytes, bytes)
}
