package adapter

import "fmt"

type MusicPlayer interface {
	play(fileType string, fileName string) string
}

type QQMusicPlayer struct {
	
}

func (*QQMusicPlayer) playMp3(fileName string) string {
	msg := "play mp3:" + fileName
	fmt.Println(msg)
	return msg
}

func (*QQMusicPlayer) playWma(fileName string) string {
	msg := "play wma:" + fileName
	fmt.Println(msg)
	return msg
}

type PlayerAdaptor struct {
	qqMusicPlayer QQMusicPlayer
}

func (player *PlayerAdaptor) play(fileType string, fileName string) string {
	switch fileType {
	case "mp3":
		return player.qqMusicPlayer.playMp3(fileName)
	case "wma":
		return player.qqMusicPlayer.playWma(fileName)
	default:
		msg := "not support type"
		fmt.Println(msg)
		return msg
	}
}