package types


type JoinedUserCount struct {
	Id          string `json:"id"`
	JoinedUsers uint16 `json:"joinedUsers"`
	RoomId      string `json:"roomId"`
}