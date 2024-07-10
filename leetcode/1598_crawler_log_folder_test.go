package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrawlerLogFolder(t *testing.T) {
    r1 := CrawlerLogFolder([]string{"d1/","d2/","../","d21/","./"})
	assert.Equal(t, 2, r1)
}
