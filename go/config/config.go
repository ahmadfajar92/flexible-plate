package config

import (
	"os"
	"strconv"

	"github.com/getsentry/raven-go"
)

// Config struct
type Config struct {
	debug bool
	port  int

	stream bool
	rpc    bool
	rest   bool
}

// NewConfig func
func NewConfig() *Config {
	return new(Config)
}

// Debug func
func (c *Config) Debug() bool {
	v := os.Getenv("DEBUG")
	c.debug, _ = strconv.ParseBool(v)

	return c.debug
}

// REST func
func (c *Config) REST() bool {
	v := os.Getenv("REST")
	c.rest, _ = strconv.ParseBool(v)

	return c.rest
}

// RPC func
func (c *Config) RPC() bool {
	v := os.Getenv("RPC")
	c.rpc, _ = strconv.ParseBool(v)

	return c.rpc
}

// STREAM func
func (c *Config) STREAM() bool {
	v := os.Getenv("STREAM")
	c.stream, _ = strconv.ParseBool(v)

	return c.stream
}

// Port func
func (c *Config) Port() int {
	v := os.Getenv("PORT")
	c.port, _ = strconv.Atoi(v)

	return c.port
}

// SentryEnable func
func (c *Config) SentryEnable() {
	sentry, _ := strconv.ParseBool(os.Getenv("SENTRY"))
	if sentry {
		dsn := os.Getenv("SENTRY_DSN")
		raven.SetDSN(dsn)
	}
}
