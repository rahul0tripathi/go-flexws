package types

import "fmt"

const (
	UpdateJoinedUserCount = "room.joined.update"
)

var (
	JoinedUsersCount = func(contestId string) string {
		return fmt.Sprintf("joinedUsersCount_%s", contestId)
	}
)
