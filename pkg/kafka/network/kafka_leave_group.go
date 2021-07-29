package network

import (
	"github.com/paashzj/kafka_go/pkg/kafka/codec"
	"github.com/paashzj/kafka_go/pkg/kafka/log"
	"github.com/paashzj/kafka_go/pkg/kafka/network/context"
	"github.com/paashzj/kafka_go/pkg/kafka/service/low"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) LeaveGroup(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 4 {
		return s.ReactLeaveGroupVersion(ctx, frame, version)
	}
	klog.Error("unknown leave group version ", version)
	return nil, gnet.Close
}

func (s *Server) ReactLeaveGroupVersion(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	req, err := codec.DecodeLeaveGroupReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("leave group req ", req)
	lowReq := &low.LeaveGroupReq{}
	lowReq.GroupId = req.GroupId
	lowReq.Members = make([]*low.LeaveGroupMember, len(req.Members))
	for i, member := range req.Members {
		m := &low.LeaveGroupMember{}
		m.MemberId = member.MemberId
		m.GroupInstanceId = member.GroupInstanceId
		lowReq.Members[i] = m
	}
	resp := codec.NewLeaveGroupResp(req.CorrelationId)
	lowResp, err := s.kafkaImpl.GroupLeave(ctx.Addr, lowReq)
	if err != nil {
		return nil, gnet.Close
	}
	resp.ErrorCode = int16(lowResp.ErrorCode)
	resp.Members = make([]*codec.LeaveGroupMember, len(lowResp.Members))
	for i, member := range resp.Members {
		m := &codec.LeaveGroupMember{}
		m.MemberId = member.MemberId
		m.GroupInstanceId = member.GroupInstanceId
		resp.Members[i] = m
	}
	resp.MemberErrorCode = int16(lowResp.MemberErrorCode)
	return resp.Bytes(), gnet.None
}
