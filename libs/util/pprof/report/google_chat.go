package report

import (
	"github.com/curtis0505/bridge/libs/util"
)

type GoogleChatReporter struct {
	AppName string
	Zone    string
}

// GoogleChatReporterOption is the option for the GoogleChat reporter.
type GoogleChatReporterOption struct {
	AppName string
	Zone    string
}

func NewGoogleChatReporter(opts GoogleChatReporterOption) *GoogleChatReporter {
	return &GoogleChatReporter{
		AppName: opts.AppName,
		Zone:    opts.Zone,
	}
}

func (g *GoogleChatReporter) ReportCPUProfile(msg, bucketName, fileName string) {
	message := util.NewMessage().
		SetZone(g.Zone).
		SetMessageType(util.MessageTypeAlert).
		SetTitle("Auto pprof Alert(CPU)").
		AddKeyValueWidget("App", g.AppName).
		AddKeyValueWidget("Zone", g.Zone).
		AddKeyValueWidget("Msg", msg).
		AddKeyValueWidget("S3 bucketName", bucketName).
		AddKeyValueWidget("FileName", fileName).
		EndSection()

	message.SendMessage()
}

func (g *GoogleChatReporter) ReportHeapProfile(msg, bucketName, fileName string) {
	message := util.NewMessage().
		SetZone(g.Zone).
		SetMessageType(util.MessageTypeAlert).
		SetTitle("Auto pprof Report(Heap)").
		AddKeyValueWidget("App", g.AppName).
		AddKeyValueWidget("Zone", g.Zone).
		AddKeyValueWidget("Msg", msg).
		AddKeyValueWidget("S3 bucketName", bucketName).
		AddKeyValueWidget("FileName", fileName).
		EndSection()

	message.SendMessage()
}
