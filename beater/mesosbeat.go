package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/berfinsari/mesosbeat/config"
)

// Mesosbeat configuration.
type Mesosbeat struct {
	done     chan struct{}
	config   config.MesosbeatConfig
	client   beat.Client
	MbConfig config.ConfigSettings
	period   time.Duration
	port     string
	host     string
}

const selector = "mesosbeat"

// New creates an instance of mesosbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	mb := &Mesosbeat{
		done: make(chan struct{}),
	}
	err := cfgfile.Read(&mb.MbConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return nil, fmt.Errorf("Error reading configuration file: %v", err)
	}
	return mb, nil
}

// Run starts mesosbeat.
func (mb *Mesosbeat) Run(b *beat.Beat) error {
	logp.Info("mesosbeat is running! Hit CTRL-C to stop it.")
	mb.CheckConfig(b)

	var err error
	mb.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(mb.period)
	for {
		select {
		case <-mb.done:
			return nil
		case <-ticker.C:
		}
		masterstats, err := mb.GetMasterStats(b)
		if err != nil {
			logp.Debug(selector, "error reading master stats")
			return err
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":   b.Info.Name,
				"master": masterstats,
			},
		}
		mb.client.Publish(event)
		logp.Info("Event sent")
	}
}

func (mb *Mesosbeat) CheckConfig(b *beat.Beat) error {
	if mb.MbConfig.Mesosbeat.Period != nil {
		mb.period = time.Duration(*mb.MbConfig.Mesosbeat.Period) * time.Second
	} else {
		mb.period = 30 * time.Second
	}

	if mb.MbConfig.Mesosbeat.Port != nil {
		mb.port = *mb.MbConfig.Mesosbeat.Port
	} else {
		mb.port = "5050"
	}

	if mb.MbConfig.Mesosbeat.Host != nil {
		mb.host = *mb.MbConfig.Mesosbeat.Host
	} else {
		mb.port = "localhost"
	}

	logp.Debug(selector, "Init Mesosbeat")
	logp.Debug(selector, "Port %v", mb.port)
	logp.Debug(selector, "Period %v", mb.period)
	logp.Debug(selector, "Host %v", mb.host)

	return nil
}

// Stop stops mesosbeat.
func (mb *Mesosbeat) Stop() {
	mb.client.Close()
	close(mb.done)
}
