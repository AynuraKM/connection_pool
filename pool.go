package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	con := ConnectionPool{}
	cc := Connection{}
	con.init(5)

	userid1 := "1"
	userid2 := "2"
	userid3 := "3"
	userid4 := "4"
	userid5 := "5"

	con.getConnection(userid1)
	con.getConnection(userid2)
	con.getConnection(userid3)
	con.getConnection(userid4)
	con.getConnection(userid5)

	fmt.Println(con)

	time.Sleep(time.Second + time.Duration(cc.TimeOut))

	con.returnConnection(userid1)
	con.returnConnection(userid3)
	con.returnConnection(userid2)

	fmt.Println(con)
}

type ConnectionPool struct {
	connections map[Connection]string
}

type Connection struct {
	Id      string
	TimeOut int
}

func (c *ConnectionPool) getConnection(userId string) Connection {

	con := Connection{}

	for k, v := range c.connections {
		if v == "_" {
			c.connections[k] = userId
			break
		}
	}
	return con
}

func (c *ConnectionPool) returnConnection(userId string) Connection {
	con := Connection{}
	for k, v := range c.connections {
		if v == userId {
			c.connections[k] = "_"
			con = k
			break
		}
	}
	return con
}

func (c *ConnectionPool) init(size int) {
	conn := make(map[Connection]string)
	for i := 0; i < size; i++ {
		val := Connection{strconv.Itoa(i), 50}
		conn[val] = "_"
	}
	c.connections = conn
}

