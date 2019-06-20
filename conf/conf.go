package conf

import (
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/welcome112s/go-library/pkg/log"
	"github.com/welcome112s/go-library/pkg/net/http/blademaster"
	"github.com/welcome112s/go-library/pkg/net/trace"
)

var (
	confPath string
	// Conf is global config
	Conf = &Config{}
)

type Config struct {
	// Log
	Log *log.Config

	// tracer
	Tracer *trace.Config

	BM         *blademaster.ServerConfig
}
func init() {
	flag.StringVar(&confPath, "conf", "./cmd/spider.toml", "config path")
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

// Init conf.
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
}
//
func remote() (err error) {
	//if client, err = conf.New(); err != nil {
	//	return
	//}
	if err = load(); err != nil {
		return
	}
	//client.Watch("activity.toml")
	//go func() {
	//	for range client.Event() {
	//		log.Info("config reload")
	//		if load() != nil {
	//			log.Error("config reload error (%v)", err)
	//		}
	//	}
	//}()
	return
}

func load() (err error) {
	var (
		s       string
		tmpConf *Config
	)
	//if s, ok = client.Toml2(); !ok {
	//	return errors.New("load config center error")
	//}
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}
