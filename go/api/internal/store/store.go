// Package store reaches the Postgres instance ../../compose.yaml publishes.
//
// The check is a TCP dial rather than a real driver: this package exists to give
// the api component a service dependency, and a driver would add a dependency
// tree that has nothing to do with what is being exercised.
package store

import (
	"net"
	"time"
)

// DefaultAddr is the address compose.yaml publishes on the host.
const DefaultAddr = "127.0.0.1:5432"

// Reachable reports whether the database is accepting connections.
func Reachable(addr string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
