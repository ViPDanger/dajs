package models

type Event struct {
	User_id string `json:"user_id"`
	Date    string `json:"date"`
}

type Result struct {
	Event  Event  `json:"result"`
	Action string `json:"action"`
}

type Error struct {
	Err error `json:"error"`
}
