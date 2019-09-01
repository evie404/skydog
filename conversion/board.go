package conversion

import (
	proto "github.com/rickypai/skydog/protobufs/datadog"
	json "github.com/zorkian/go-datadog-api"
)

func BoardJSON(protoBoard *proto.Board) (jsonBoard *json.Board) {
	return nil
}

func BoardProto(jsonBoard *json.Board) (protoBoard *proto.Board) {
	return nil
}
