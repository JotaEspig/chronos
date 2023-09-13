// package scheduling provides support for operations with Scheduling model
package scheduling

type Scheduling struct {
	ID             uint
	Start, End     string
	UserID, TimeID uint
}

func (s *Scheduling) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = s.ID
	m["start"] = s.Start
	m["end"] = s.End
	m["user_id"] = s.UserID
	m["time_id"] = s.TimeID
	return m
}
