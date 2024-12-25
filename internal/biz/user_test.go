package biz

import (
    "testing"

    "github.com/davecgh/go-spew/spew"
    "github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
    s := hashPassword("abc")
    spew.Dump(s)
}
func TestVerifyPassword(t *testing.T) {
    a := assert.New(t)
    a.True(verifyPassword("$2a$10$efRreXBQ/uwa8cNDcxMazORGIbIWrqYKjn4qG6YT2tSWg9Sy.xARu", "abc"))
    spew.Dump("abc")
}
