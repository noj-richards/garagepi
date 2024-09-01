/*
For Raspberry PI Version B (32bit)
GOOS=linux GOARCH=arm GOARM=6 go build main.go
*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"garagepi/appconfig"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/kmmndr/goPi/piface"
	"github.com/kmmndr/goPi/spi"
	"google.golang.org/protobuf/encoding/prototext"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var videoInProgress atomic.Bool

func (p *realPi) init() error {
	p.pfd = piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)
	return p.pfd.InitBoard()
}

func (p *realPi) toggleDoor(relayNumber int) error {
	if relayNumber < 0 || relayNumber >= len(p.pfd.Relays) {
		return fmt.Errorf("Invalid relay number: %v", relayNumber)
	}
	p.pfd.Relays[relayNumber].SetValue(1)
	time.Sleep(1 * time.Second)
	p.pfd.Relays[relayNumber].SetValue(0)
	return nil
}

func (p *realPi) isOpen(doorInputNumber int) (bool, error) {
	if doorInputNumber < 0 || doorInputNumber >= len(p.pfd.InputPins) {
		return false, fmt.Errorf("Invalid input number: %v", doorInputNumber)
	}
	return p.pfd.InputPins[doorInputNumber].Value() == 1, nil
}

func (p *realPi) takePic() error {
	cmd := exec.Command("raspistill", "-o", "snap.jpg", "-vf", "-hf", "-t", "1")
	return cmd.Run()
}

type pi interface {
	init() error
	toggleDoor(relayNumber int) error
	isOpen(inputNumber int) (bool, error)
	takePic() error
}

type fakepi struct {
}

func (p *fakepi) init() error {
	fmt.Println("Fake Pi initialized")
	return nil
}

func (p *fakepi) toggleDoor(relayNumber int) error {
	fmt.Println("Toggling door")
	return nil
}

func (p *fakepi) isOpen(inputNumber int) (bool, error) {
	return true, nil
}

func (p *fakepi) takePic() error {
	return nil
}

type realPi struct {
	pfd *piface.PiFaceDigital
}

func processMessage(config *appconfig.AppConfig, message string, id int64, p pi, bot *tgbotapi.BotAPI) error {
	if strings.ToLower(message) == "status" {
		status(config, p, id, bot)
	}
	if strings.ToLower(message) == "open" {
		return openDoor(config, p, id, bot)
	}
	if strings.ToLower(message) == "toggle" {
		return toggleDoor(config, p, id, bot)
	}
	if strings.ToLower(message) == "close" {
		return closeDoor(config, p, id, bot)
	}
	if strings.ToLower(message) == "pic" {
		return takeAndSendPic(id, p, bot)
	}
	if strings.ToLower(message) == "video" {
		go recordAndSendVideo(id, bot)
	}
	return nil
}

func takeAndSendPic(id int64, p pi, bot *tgbotapi.BotAPI) error {
	if err := p.takePic(); err != nil {
		return err
	}
	return sendPic(id, bot)
}

func sendPic(id int64, bot *tgbotapi.BotAPI) error {
	filePath := "snap.jpg"
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new photo upload request
	photo := tgbotapi.NewPhoto(id, tgbotapi.FileReader{
		Name:   filePath,
		Reader: file,
	})
	// Send the photo
	_, err = bot.Send(photo)
	if err != nil {
		return err
	}
	return nil
}

func sendVideo(id int64, bot *tgbotapi.BotAPI) error {
	filePath := "video2.mp4"
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new photo upload request
	video := tgbotapi.NewInputMediaVideo(tgbotapi.FileReader{
		Name:   filePath,
		Reader: file,
	})

	mediaGroup := tgbotapi.NewMediaGroup(id, []interface{}{video})

	// Send the video
	// Send the media group using the bot.Request method
	response, err := bot.Request(mediaGroup)
	if err != nil {
		return err
	}

	// Handle the response
	var messages []tgbotapi.Message
	if err := json.Unmarshal(response.Result, &messages); err != nil {
		return err
	}
	for _, msg := range messages {
		log.Printf("Message ID: %d\n", msg.MessageID)
	}
	return nil
}

func recordAndSendVideo(id int64, bot *tgbotapi.BotAPI) {
	if !videoInProgress.CompareAndSwap(false, true) {
		bot.Send(tgbotapi.NewMessage(id, "Video capture already in progress, can't start another"))
		return
	}

	defer videoInProgress.Store(false)

	bot.Send(tgbotapi.NewMessage(id, "Video capture started..."))
	if err := os.Remove("video2.h264"); err != nil && !errors.Is(err, os.ErrNotExist) {
		logToGroup(id, bot, fmt.Sprintf("Error removing source file %v", err.Error()))
	}
	if err := os.Remove("video2.mp4"); err != nil && !errors.Is(err, os.ErrNotExist) {
		logToGroup(id, bot, fmt.Sprintf("Error removing converted file %v", err.Error()))
	}
	if err := exec.Command("raspivid", "-hf", "-vf", "-w", "640", "-h", "480", "-o", "video2.h264", "-t", "30000").Run(); err != nil {
		logToGroup(id, bot, fmt.Sprintf("Error calling raspivid %v", err.Error()))
		return
	}
	if err := exec.Command("ffmpeg", "-i", "video2.h264", "-vcodec", "copy", "video2.mp4").Run(); err != nil {
		logToGroup(id, bot, fmt.Sprintf("Error calling ffmpeg %v", err.Error()))
		return
	}
	if err := sendVideo(id, bot); err != nil {
		logToGroup(id, bot, fmt.Sprintf("Error sending video %v", err.Error()))
	}
}

func status(config *appconfig.AppConfig, p pi, id int64, bot *tgbotapi.BotAPI) {
	open, err := p.isOpen(int(config.DoorInputNumber))
	if err != nil {
		bot.Send(tgbotapi.NewMessage(id, fmt.Sprintf("Failed to check door status: %v", err.Error())))
		return
	}
	if open {
		bot.Send(tgbotapi.NewMessage(id, "The door is open"))
	}
	msg := tgbotapi.NewMessage(id, "The door is closed")
	bot.Send(msg)
}

func openDoor(config *appconfig.AppConfig, p pi, id int64, bot *tgbotapi.BotAPI) error {
	open, err := p.isOpen(int(config.DoorInputNumber))
	if err != nil {
		bot.Send(tgbotapi.NewMessage(id, fmt.Sprintf("Failed to check door status: %v", err.Error())))
		return nil
	}
	if open {
		bot.Send(tgbotapi.NewMessage(id, "The door is already open"))
		return nil
	}
	if err := p.toggleDoor(int(config.DoorRelayNumber)); err != nil {
		return err
	}
	bot.Send(tgbotapi.NewMessage(id, "The door is opening"))
	if id != config.GroupChannelId && config.GroupChannelId != 0 {
		bot.Send(tgbotapi.NewMessage(config.GroupChannelId, "The door is opening"))
	}
	takeDelayedPic(id, config.GroupChannelId, p, bot)
	return nil
}

func toggleDoor(config *appconfig.AppConfig, p pi, id int64, bot *tgbotapi.BotAPI) error {
	err := p.toggleDoor(int(config.DoorRelayNumber))
	if err != nil {
		return err
	}
	bot.Send(tgbotapi.NewMessage(id, "The door is opening or closing"))
	// also send to channel if available
	if id != config.GroupChannelId && config.GroupChannelId != 0 {
		bot.Send(tgbotapi.NewMessage(config.GroupChannelId, "The door is opening or closing"))
	}
	go recordAndSendVideo(id, bot)
	return nil
}

func closeDoor(config *appconfig.AppConfig, p pi, id int64, bot *tgbotapi.BotAPI) error {
	open, err := p.isOpen(int(config.DoorInputNumber))
	if err != nil {
		return err
	}
	closed := !open
	if closed {
		bot.Send(tgbotapi.NewMessage(id, "The door is already closed"))
		return nil
	}
	if err := p.toggleDoor(int(config.DoorRelayNumber)); err != nil {
		return err
	}
	bot.Send(tgbotapi.NewMessage(id, "The door is closing"))
	logToGroup(config.GroupChannelId, bot, "The door is closing")
	takeDelayedPic(id, config.GroupChannelId, p, bot)
	return nil
}

func takeDelayedPic(id int64, groupChannelD int64, p pi, bot *tgbotapi.BotAPI) {
	go func() {
		time.Sleep(30 * time.Second)
		err := p.takePic()
		if err != nil {
			logToGroup(groupChannelD, bot, err.Error())
			return
		}
		if err := sendPic(id, bot); err != nil {
			logToGroup(groupChannelD, bot, err.Error())
		}
		// Also send to group channel if available
		if id != groupChannelD && groupChannelD != 0 {
			if err := sendPic(groupChannelD, bot); err != nil {
				logToGroup(groupChannelD, bot, err.Error())
			}
		}
	}()
}

func loadConfig(filename string) (*appconfig.AppConfig, error) {
	textData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := &appconfig.AppConfig{}

	err = prototext.Unmarshal(textData, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func authorised(userIds []int64, userId int64) bool {
	for _, v := range userIds {
		if v == userId {
			return true
		}
	}
	return false
}

func logToGroup(groupChannelId int64, bot *tgbotapi.BotAPI, logmsg string) {
	log.Printf("Failed to process message: %v", logmsg)
	if groupChannelId == 0 {
		return
	}
	msg := tgbotapi.NewMessage(groupChannelId, fmt.Sprintf("Failed to process message: %v", logmsg))
	bot.Send(msg)
}

type serverConfig struct {
	config *appconfig.AppConfig
	bot    *tgbotapi.BotAPI
	p      pi
}

func (s *serverConfig) webhookHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("auth")
	if query != s.config.WebhookAuthToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	query = r.URL.Query().Get("cmd")
	if query == "toggle" {
		err := toggleDoor(s.config, s.p, s.config.GroupChannelId, s.bot)
		if err != nil {
			logToGroup(s.config.GroupChannelId, s.bot, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	if query == "pic" {
		err := takeAndSendPic(s.config.GroupChannelId, s.p, s.bot)
		if err != nil {
			logToGroup(s.config.GroupChannelId, s.bot, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

}

func (s *serverConfig) startServer() {
	log.Println("Server listening on port 8080")
	server := &http.Server{
		Addr:              ":8080",
		Handler:           http.DefaultServeMux, // Or your custom handler
		ReadTimeout:       10 * time.Second,     // Maximum duration for reading the entire request, including the body
		WriteTimeout:      10 * time.Second,     // Maximum duration before timing out writes of the response
		IdleTimeout:       60 * time.Second,     // Maximum amount of time to wait for the next request when keep-alives are enabled
		ReadHeaderTimeout: 5 * time.Second,      // Maximum time to read request headers
	}
	http.HandleFunc("/webhook", s.webhookHandler)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Failed to start server: %v", err)
		}
	}()
}

func main() {
	config, err := loadConfig("config.txt")
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
	}
	var p pi
	if runtime.GOOS == "darwin" {
		p = &fakepi{}
	} else {
		p = &realPi{}
	}
	if err := p.init(); err != nil {
		log.Panicf("Failed to initialize Pi: %v", err)
	}

	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	// Print bot info
	log.Printf("Authorized on account %s", bot.Self.UserName)

	sConfig := &serverConfig{
		config: config,
		bot:    bot,
		p:      p,
	}
	sConfig.startServer()

	// Set up a channel to receive updates from Telegram
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Process each update (message) received from Telegram
	for update := range updates {
		if update.ChannelPost != nil {
			if update.ChannelPost.Chat.ID != config.GroupChannelId {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, you are not authorized to use this bot")
				bot.Send(msg)
				continue
			}
			err := processMessage(config, update.ChannelPost.Text, update.ChannelPost.Chat.ID, p, bot)
			if err != nil {
				logToGroup(config.GroupChannelId, bot, err.Error())
			}
			continue
		}
		if update.Message == nil {
			log.Printf("chat id: %v, msg %v\n", update.ChannelPost.Chat.ID, update.ChannelPost.Text)
			continue
		}

		if !authorised(config.AllowListIds, update.Message.From.ID) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, you are not authorized to use this bot")
			bot.Send(msg)
			logToGroup(config.GroupChannelId, bot, fmt.Sprintf("Unauthorized user: %v", update.Message.From.ID))
			continue
		}
		if err := processMessage(config, update.Message.Text, update.Message.Chat.ID, p, bot); err != nil {
			logToGroup(config.GroupChannelId, bot, err.Error())
		}
	}
}
