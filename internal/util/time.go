package util

import (
    "google.golang.org/protobuf/types/known/timestamppb"
    "time"
)

func TimeToTimestamp(t *time.Time) *timestamppb.Timestamp {
    if t == nil {
        return nil
    }
    var timestamp = timestamppb.Timestamp{
        Seconds: t.Unix(),
        Nanos:   int32(t.Nanosecond()),
    }
    return &timestamp
}
