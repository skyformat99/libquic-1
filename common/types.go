package common

import (
	"bytes"
	"crypto/rand"
	"io"
	"math/big"
)

// ConnectionID for large number manipulation
type ConnectionID struct {
	value big.Int
}

// SetByRand set connection by random
func (c *ConnectionID) SetByRand() error {
	buf := make([]byte, 18)
	if n, err := rand.Read(buf[:]); err != nil {
		return err
	} else {
		c.value.SetBytes(buf[:n])
	}
	return nil
}

// SetByReader read from a buffer and set value
func (c *ConnectionID) SetByReader(r *bytes.Buffer, len int) (err error) {
	if len < 4 || len > 18 {
		return ErrInvalidConnectionID
	}

	buf := make([]byte, len)
	if _, err = io.ReadFull(r, buf[:len]); err != nil {
		return err
	}
	c.value.SetBytes(buf)
	return nil
}

// SetByBytes read from bytes and set value
func (c *ConnectionID) SetByBytes(data []byte) error {
	length := len(data)
	if length < 4 || length > 18 {
		return ErrInvalidConnectionID
	}

	c.value.SetBytes(data)
	return nil
}

// SetByValue set connection id via multiple uint64 value
func (c *ConnectionID) SetByValue(value ...uint64) error {
	if len(value) > 3 || (len(value) == 3 && value[0] > (1<<(144-128)-1)) {
		return ErrVarLenIntTooLarge
	}

	return nil
}

// Equal compare connection id value equal
func (c *ConnectionID) Equal(other *ConnectionID) bool {
	return bytes.Equal(c.value.Bytes(), other.value.Bytes())
}

// Bytes of connection id
func (c *ConnectionID) Bytes() []byte {
	return c.value.Bytes()
}

// Len of connection id
func (c *ConnectionID) Len() int {
	return c.value.BitLen()
}

// String the connection id number string
func (c *ConnectionID) String() string {
	return c.value.String()
}
