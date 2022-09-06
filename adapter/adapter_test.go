package adapter

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAdaptor(t *testing.T)  {
	player := PlayerAdaptor{}
	fileName := "十年"
	msg := player.play("mp3", fileName)
	assert.Equal(t, "play mp3:" + fileName, msg)

	msg = player.play("wma", fileName)
	assert.Equal(t, "play wma:" + fileName, msg)

	msg = player.play("mp4", fileName)
	assert.Equal(t, "not support type", msg)
}
