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

func TestCodeFetchRespV10(t *testing.T) {
	record := &Record{}
	record.RecordAttributes = 0
	record.RelativeTimestamp = 0
	record.RelativeOffset = 0
	record.Value = "msg"
	recordBatch := &RecordBatch{}
	recordBatch.Offset = 0
	recordBatch.MessageSize = 59
	recordBatch.LeaderEpoch = 0
	recordBatch.MagicByte = 2
	recordBatch.Flags = 0
	recordBatch.LastOffsetDelta = 0
	recordBatch.FirstTimestamp = 1640435896600
	recordBatch.LastTimestamp = 1640435896600
	recordBatch.ProducerId = -1
	recordBatch.ProducerEpoch = -1
	recordBatch.BaseSequence = -1
	recordBatch.Records = []*Record{record}
	fetchPartitionResp := &FetchPartitionResp{}
	fetchPartitionResp.PartitionIndex = 0
	fetchPartitionResp.ErrorCode = 0
	fetchPartitionResp.HighWatermark = 1
	fetchPartitionResp.LastStableOffset = 1
	fetchPartitionResp.LogStartOffset = 0
	fetchPartitionResp.AbortedTransactions = -1
	fetchPartitionResp.RecordBatch = recordBatch
	fetchTopicResp := &FetchTopicResp{}
	fetchTopicResp.Topic = "topic"
	fetchTopicResp.PartitionDataList = []*FetchPartitionResp{fetchPartitionResp}
	fetchResp := NewFetchResp(6)
	fetchResp.ErrorCode = 0
	fetchResp.SessionId = 0
	fetchResp.TopicResponses = []*FetchTopicResp{fetchTopicResp}
	bytes := fetchResp.Bytes(10)
	expectBytes := testHex2Bytes(t, "0000000600000000000000000000000000010005746f70696300000001000000000000000000000000000100000000000000010000000000000000ffffffff0000004700000000000000000000003b00000000022c30096c0000000000000000017df19951180000017df1995118ffffffffffffffffffffffffffff000000011200000001066d736700")
	assert.Equal(t, expectBytes, bytes)
}

func TestCodeFetchRespV11(t *testing.T) {
	record := &Record{}
	record.RecordAttributes = 0
	record.RelativeTimestamp = 0
	record.RelativeOffset = 0
	record.Value = "Hzj"
	recordBatch := &RecordBatch{}
	recordBatch.Offset = 0
	recordBatch.MessageSize = 59
	recordBatch.LeaderEpoch = 0
	recordBatch.MagicByte = 2
	recordBatch.Flags = 0
	recordBatch.LastOffsetDelta = 0
	recordBatch.FirstTimestamp = 1625965841631
	recordBatch.LastTimestamp = 1625965841631
	recordBatch.ProducerId = -1
	recordBatch.ProducerEpoch = -1
	recordBatch.BaseSequence = -1
	recordBatch.Records = []*Record{record}
	fetchPartitionResp := &FetchPartitionResp{}
	fetchPartitionResp.PartitionIndex = 0
	fetchPartitionResp.ErrorCode = 0
	fetchPartitionResp.HighWatermark = 1
	fetchPartitionResp.LastStableOffset = 1
	fetchPartitionResp.LogStartOffset = 0
	fetchPartitionResp.AbortedTransactions = -1
	fetchPartitionResp.ReplicaData = -1
	fetchPartitionResp.RecordBatch = recordBatch
	fetchTopicResp := &FetchTopicResp{}
	fetchTopicResp.Topic = "test-5"
	fetchTopicResp.PartitionDataList = []*FetchPartitionResp{fetchPartitionResp}
	fetchResp := NewFetchResp(10)
	fetchResp.ErrorCode = 0
	fetchResp.SessionId = 997895662
	fetchResp.TopicResponses = []*FetchTopicResp{fetchTopicResp}
	bytes := fetchResp.Bytes(11)
	expectBytes := testHex2Bytes(t, "0000000a0000000000003b7aadee000000010006746573742d3500000001000000000000000000000000000100000000000000010000000000000000ffffffffffffffff0000004700000000000000000000003b000000000206cbcc440000000000000000017a931dccdf0000017a931dccdfffffffffffffffffffffffffffff00000001120000000106487a6a00")
	assert.Equal(t, expectBytes, bytes)
}
