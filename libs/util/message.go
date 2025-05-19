package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	alertLive = "https://chat.googleapis.com/v1/spaces/AAAAWtFROGI/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=6D5_AyVRf5TQ8J-QbsMq7YUzqp6DellSgbnGzyb1qLA%3D"
	alertDQ   = "https://chat.googleapis.com/v1/spaces/AAAAzYK8RK4/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=yry5PuM8fH1kkLkqfsYavD-_tAXwKB3DEY8rTanu7aE%3D"
	alertDEV  = "https://chat.googleapis.com/v1/spaces/AAAAmzSuwiM/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=WBeAJmPVYLCpdTo9ovhgQ-VWmj17qeYWt4eaaQbUdrY%3D"

	reportLive = "https://chat.googleapis.com/v1/spaces/AAAAHy-efxI/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=Id6mNfaUmI2hETUQsRP3K-ZaZVDJIrev7pf7d1oV5MA%3D"
	reportDQ   = alertDQ
	reportDEV  = alertDEV

	backend   = "https://chat.googleapis.com/v1/spaces/AAAAdZp4aQY/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=_SfiSbORYLcziahLxzIOeOaCoaoyElfxE8Aq3oznY-A%3D"
	kycStatus = "https://chat.googleapis.com/v1/spaces/AAAAhTNBv5k/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=FUJ9AUixPponWxzcTqAbhktg_Fb769K6qCMPRJBgTDI"

	panicLive = "https://chat.googleapis.com/v1/spaces/AAAAUMjfbzA/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=eLKKRLeIVF0erAHW08us054yMNWnqxTr3Zh4yi7lRis%3D"
	panicDQ   = "https://chat.googleapis.com/v1/spaces/AAAAwnlYCa4/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=NOsJ6UZnYzeaUfpeDkSX8ffK2NggYerjRed9KjEWaik%3D"
	panicDEV  = "https://chat.googleapis.com/v1/spaces/AAAAq8FFmbc/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=NMqwxWtROy2SUNXvewseuv47N-TG8pby54DL_H5mpfM%3D"
)

const (
	AppName    = "[Batch]"
	TitleError = AppName + " Error Report"
	TitlePanic = AppName + " Panic Report"

	// Deprecated: TitleScheduler
	TitleScheduler = AppName + " Scheduler Report"
	TitleApp       = AppName + " Report"

	MessageTypePanic  = "panic"
	MessageTypeAlert  = "alert"
	MessageTypeReport = "report"

	MessageTypeBackendAlert = "backend"

	// Deprecated: MessageTypeScheduler
	MessageTypeScheduler = "scheduler"
)

type Message struct {
	zone        string
	messageType string
	backend     bool
	card        int
	section     int
	start       time.Time

	Cards []Card `json:"cards"`
}

type Card struct {
	Header   Header    `json:"header"`
	Sections []Section `json:"sections"`
}

type Header struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle,omitempty"`
}

type Section struct {
	Widgets []map[string]interface{} `json:"widgets"`
}

func NewMessage() *Message {
	return &Message{
		backend: false,
		card:    0,
		section: 0,
		Cards: []Card{{
			Sections: []Section{{
				Widgets: []map[string]interface{}{},
			}},
		}},
	}
}

func (m *Message) SetZone(zone string) *Message {
	m.zone = zone
	return m
}

func (m *Message) SetMessageType(messageType string) *Message {
	m.messageType = messageType
	return m
}

func (m *Message) SetTitle(title string) *Message {
	m.Cards[m.card].Header.Title = title
	return m
}

func (m *Message) SetSubtitle(subtitle string) *Message {
	m.Cards[m.card].Header.SubTitle = subtitle
	return m
}

func (m *Message) AlertBackend() *Message {
	m.backend = true
	return m
}

func (m *Message) EndSection() *Message {
	if len(m.Cards[m.card].Sections[m.section].Widgets) == 0 {
		return m
	}

	m.section++
	m.Cards[m.card].Sections = append(m.Cards[m.card].Sections, Section{
		Widgets: []map[string]interface{}{},
	})
	return m
}

func (m *Message) AddKeyValueWidget(key, value string) *Message {
	if value == "" {
		value = "Empty"
	}

	keyValue := struct {
		TopLabel string `json:"topLabel"`
		Content  string `json:"content"`
	}{
		TopLabel: key,
		Content:  value,
	}

	widget := make(map[string]interface{})
	widget["keyValue"] = keyValue

	m.Cards[m.card].Sections[m.section].Widgets = append(m.Cards[m.card].Sections[m.section].Widgets, widget)
	return m
}

func (m *Message) AddTextParagraphWidget(text string) *Message {
	if text == "" {
		text = "Empty"
	}

	textParagraph := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}

	widget := make(map[string]interface{})
	widget["textParagraph"] = textParagraph

	m.Cards[m.card].Sections[m.section].Widgets = append(m.Cards[m.card].Sections[m.section].Widgets, widget)
	return m
}

// ----------
// scheduler

func (m *Message) StartScheduler() *Message {
	m.start = time.Now()

	return m
}

func (m *Message) EndScheduler() *Message {
	//m.AddKeyValueWidget("Total Duration", time.Now().Sub(m.start).String())
	return m
}

func (m *Message) AddScheduler(name string, scheduler func() error) *Message {
	schedulerTime := time.Now()

	if err := scheduler(); err != nil {
		m.AddKeyValueWidget("Job", name)
		m.AddKeyValueWidget("Failed", err.Error())
	} else {
		m.AddKeyValueWidget("Job", fmt.Sprintf("%s: %s", name, time.Now().Sub(schedulerTime).String()))
	}

	m.EndSection()

	return m
}

func (m *Message) AddSchedulerFunc(name string, scheduler func()) *Message {
	schedulerTime := time.Now()

	scheduler()

	m.AddKeyValueWidget("Job", fmt.Sprintf("%s: %s", name, time.Now().Sub(schedulerTime).String()))
	m.EndSection()

	return m
}

func (m *Message) AddSchedulerString(name string, scheduler func(string), value string) *Message {
	schedulerTime := time.Now()

	scheduler(value)

	m.AddKeyValueWidget("Job", fmt.Sprintf("%s: %s", name, time.Now().Sub(schedulerTime).String()))
	m.EndSection()

	return m
}

func (m *Message) AddSchedulerAdmin(name string, scheduler func(bool, time.Time) error, isDaily bool, t time.Time) *Message {
	schedulerTime := time.Now()

	if err := scheduler(isDaily, t); err != nil {
		m.AddKeyValueWidget("Job", name)
		m.AddKeyValueWidget("Failed", err.Error())
	} else {
		m.AddKeyValueWidget("Job", fmt.Sprintf("%s: %s", name, time.Now().Sub(schedulerTime).String()))
	}

	m.EndSection()

	return m
}

func (m *Message) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(b)
}

func (m *Message) SendMessage() {
	go func() {
		defer func() {
			recover()
		}()

		if len(m.Cards[m.card].Sections[m.section].Widgets) == 0 {
			m.Cards[m.card].Sections = m.Cards[m.card].Sections[:m.section]
		}

		url := alertDEV

		switch m.messageType {
		case MessageTypeAlert:
			switch m.zone {
			case "prod":
				url = alertLive
			case "dq":
				url = alertDQ
			case "kyc":
				url = kycStatus
			default:
				url = alertDEV
			}

		case MessageTypeReport:
			switch m.zone {
			case "prod":
				url = reportLive
			case "dq":
				url = reportDQ
			default:
				url = reportDEV
			}

		case MessageTypeScheduler:
			// 스케쥴러는 앞을 삭제될 예정
			return

		case MessageTypePanic:
			switch m.zone {
			case "prod":
				url = panicLive
			case "dq":
				url = panicDQ
			default:
				url = panicDEV
			}
		}
		if m.backend && m.zone == "prod" {
			go http.Post(backend, "application/json", bytes.NewBuffer([]byte(m.String())))
		}
		_, _ = http.Post(url, "application/json", bytes.NewBuffer([]byte(m.String())))
	}()
}
