package main

import "net"

var users = make(map[net.Conn]*UserConn)

type UserConn struct {
	conn net.Conn
}

//users := make([]UserConn, 10)
//userCnt := 0

func NewUser(conn net.Conn) *UserConn {
	return &UserConn{
		conn: conn,
	}
}

func addUser(user *UserConn) {
	conn := user.conn
	users[conn] = user
}

func (u *UserConn) equal(user *UserConn) bool {
	return u == user
}
