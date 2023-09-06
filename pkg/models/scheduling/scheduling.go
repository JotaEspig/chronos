// package scheduling provides support for operations with Scheduling model
package scheduling

type Scheduling struct {
	ID             uint
	Start, End     string
	UserID, TimeID uint
}
