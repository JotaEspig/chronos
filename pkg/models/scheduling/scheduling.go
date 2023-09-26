// package scheduling provides support for operations with Scheduling model
package scheduling

import (
	"chronos/pkg/types"

	"github.com/microcosm-cc/bluemonday"
)

type Scheduling struct {
	ID     uint   `json:"id"`
	Start  string `json:"start"`
	End    string `json:"end"`
	UserID uint   `json:"user_id"`
	TimeID uint   `json:"time_id"`
}

func (s *Scheduling) Sanitize(policy *bluemonday.Policy) {
	s.Start = policy.Sanitize(s.Start)
	s.End = policy.Sanitize(s.End)
}

func (s *Scheduling) ToMap() types.JsonMap {
	m := make(types.JsonMap)
	m["id"] = s.ID
	m["start"] = s.Start
	m["end"] = s.End
	m["user_id"] = s.UserID
	m["time_id"] = s.TimeID
	return m
}
