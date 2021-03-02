package main

type cmdId int

const (
	CMD_NAME cmdId = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     cmdId
	client *client
	args   []string
}
