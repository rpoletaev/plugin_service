package main

import (
	log "github.com/Sirupsen/logrus"
	ps "github.com/rpoletaev/plugin_service"
	"github.com/weekface/mgorus"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"plugin"
	"reflect"
)

const (
	pluginFunc = "GetExportInterface"
)

type MongoLogHookConfig struct {
	Host       string `yaml:"log_host"`
	System     string `yaml:"log_system"`
	DBName     string `yaml:"log_dbname"`
	Collection string `yaml:"log_collection"`
}

type Config struct {
	LogHook      *MongoLogHookConfig `yaml:"log_hook"`
	Mongo        string              `yaml:"mongo"`
	MongoBase    string              `yaml:"mongo_base"`
	MySQL        string              `yaml:"mysql"`
	RedisAddress string              `yaml:"redis_address"`
}

var (
	loger  *log.Logger
	config *Config
	mongo  *mgo.Session
)

func main() {
	config = &Config{}
	loger = log.New()
	readConfig()
	setupLoger()
	mongodb, err := mgo.Dial(config.Mongo)
	if err != nil {
		logEntry().Fatalf("Не удалось подключиться к MongoDB: %v", err)
	}
	mongo = mongodb
	mongo.SetMode(mgo.Monotonic, true)

	http.HandleFunc("/load", getExportHandler)
	logEntry().Error(http.ListenAndServe(":3131", nil))
}

func getExportHandler(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logEntry().Error(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logEntry().Info("Читаем данные")
	ei, err := ps.GetExportInfo(string(body))
	if err != nil {
		logEntry().Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	logEntry().Info("Получаем функцию")
	exportFunc, err := getPluginFunc()
	if err != nil {
		logEntry().Errorf("Не удалось загрузить плагин: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logEntry().Info("Получаем данные")
	obj := exportFunc(body, ei.Title)
	if obj == nil {
		logEntry().Errorf("Не удалось получить объект из плагина!")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logEntry().Info("Пишем в монгу")
	//Возвращаться может как одно значение, так и слайс, поэтому предварительно обрабатываем, проверяем и сохраняем каждое
	err = storeExportObject(ei.Title, obj)
	if err != nil {
		logEntry().Error(err)
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
		logEntry().Fatalf("Ошбка при чтении конфига: %v", err)
	}

	err = yaml.Unmarshal(cnfBts, &config)
	if err != nil {
		logEntry().Fatalf("Ошбка конфигурации %v", err)
	}
}

func setupLoger() {
	logHook, err := mgorus.NewHooker(config.LogHook.Host, config.LogHook.DBName, config.LogHook.System)
	if err != nil {
		logEntry().Fatalf("Ошибка при инициализации логера: %v", err)
	}
	loger.Hooks.Add(logHook)
}

func logEntry() *log.Entry {
	return loger.WithFields(log.Fields{
		"system": config.LogHook.System,
		"host":   config.LogHook.Host,
	})
}

func mongoExec(colectionName string, execFunc func(*mgo.Collection) error) error {
	session := mongo.Clone()
	defer session.Close()

	db := session.DB(config.MongoBase)
	collection := db.C(colectionName)
	return execFunc(collection)
}

func storeExportObject(colName string, obj interface{}) error {
	objRefVal := reflect.ValueOf(obj)
	if objRefVal.Kind() != reflect.Slice {
		return mongoExec(colName, func(col *mgo.Collection) error {
			return col.Insert(obj)
		})
	}

	objLen := objRefVal.Len()
	if objLen > 0 {
		for i := 0; i < objLen; i++ {
			if err := mongoExec(colName, func(col *mgo.Collection) error {
				return col.Insert(objRefVal.Index(i).Interface())
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
