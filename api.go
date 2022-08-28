package ts_backapi

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type UnixTimestampString struct{ time.Time }

type ResponseProcessOrder struct {
	OrderId              string               `json:"order_id"`
	OrderDescription     string               `json:"order_description"`
	OrderStatus          string               `json:"order_status"`
	LastUpdatedTimestamp *UnixTimestampString `json:"last_updated_timestamp"`
	SpecialOrder         bool                 `json:"special_order"`
}

type RequestProcessOrder struct {
	OrderId string `json:"order_id"`
	// request param
}

func (t *UnixTimestampString) MarshalJSON() ([]byte, error) {
	ts := t.Time.UnixMilli()
	stamp := fmt.Sprintf("\"%d\"", ts)
	return []byte(stamp), nil
}
func (t *UnixTimestampString) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	ts, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	t.Time = time.UnixMilli(int64(ts))
	return nil
}
