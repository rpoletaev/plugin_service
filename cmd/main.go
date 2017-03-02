package main

import (
	"github.com/Sirupsen/logrus"
	log "github.com/Sirupsen/logrus"
	ps "github.com/rpoletaev/plugin_service"
	"gopkg.in/mgo.v2"
	"gopkg.in/polds/logrus-papertrail-hook.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"plugin"
)

const (
	pluginFunc = "GetExportInterface"
)

type LogHook struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	From string `yaml:"from"`
}

type Config struct {
	LogHook      *LogHook `yaml:"log_hook"`
	Mongo        string   `yaml:"mongo"`
	MongoBase    string   `yaml:"mongo_base"`
	MySQL        string   `yaml:"mysql"`
	Port         string   `yaml:"port"`
	RedisAddress string   `yaml:"redis_address"`
}

var (
	loger  *log.Logger
	config *Config
	mongo  *mgo.Session
)

func main() {
	readConfig()
	setupLoger()

	http.HandleFunc("/load", getExportHandler)
	logrus.Error(http.ListenAndServe(":3131", nil))
}

func getExportHandler(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		loger.Error(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	ei, err := ps.GetExportInfo(string(body))
	if err != nil {
		loger.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	exportFunc, err := getPluginFunc()
	if err != nil {
		loger.Errorf("Не удалось загрузить плагин: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	obj := exportFunc(body, ei.Title)
	err = mongoExec(ei.Title, func(c *mgo.Collection) error {
		return c.Insert(obj)
	})

	if err != nil {
		loger.Error()
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func getPluginFunc() (func([]byte, string) interface{}, error) {
	p, err := plugin.Open("export.so")
	if err != nil {
		return nil, err
	}

	f, err := p.Lookup(pluginFunc)
	if err != nil {
		return nil, err
	}

	exportFunc := f.(func([]byte, string) interface{})
	return exportFunc, nil
}

func readConfig() {
	cnfBts, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		loger.Fatalf("Ошбка при чтении конфига: %v", err)
	}

	err = yaml.Unmarshal(cnfBts, config)
	if err != nil {
		loger.Fatalf("Ошбка конфигурации %v", err)
	}
}

func setupLoger() {
	logHook, err := logrus_papertrail.NewPapertrailHook(&logrus_papertrail.Hook{
		Host:     config.LogHook.Host,
		Port:     config.LogHook.Port,
		Appname:  "save_export",
		Hostname: config.LogHook.From,
	})
	if err != nil {
		loger.Fatalf("Ошибка при инициализации логера: %v", err)
	}
	loger.Hooks.Add(logHook)
}

func setupMongo() {
	mongo, err := mgo.Dial(config.Mongo)
	if err != nil {
		loger.Fatalf("Не удалось подключиться к MongoDB: %v", err)
	}
	mongo.SetMode(mgo.Monotonic, true)
}

func mongoExec(colectionName string, execFunc func(*mgo.Collection) error) error {
	session := mongo.Clone()
	defer session.Close()

	db := session.DB(config.MongoBase)
	collection := db.C(colectionName)
	return execFunc(collection)
}
