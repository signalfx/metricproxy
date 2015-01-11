package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestProxyPidWrite(t *testing.T) {
	fileObj, _ := ioutil.TempFile("", "gotest")
	filename := fileObj.Name()
	defer os.Remove(filename)

	assert.Nil(t, writePidFile(filename))
}

func TestProxyPidWriteError(t *testing.T) {
	assert.Error(t, writePidFile("/root"))
}

func TestProxyInvalidConfig(t *testing.T) {
	fileObj, _ := ioutil.TempFile("", "gotest")
	filename := fileObj.Name()
	defer os.Remove(filename)

	ioutil.WriteFile(filename, []byte{}, os.FileMode(0666))
	proxyCommandLineConfiguration = proxyCommandLineConfigurationT{
		configFileName: filename,
		logDir:         os.TempDir(),
		logMaxSize:     1,
		logMaxBackups:  0,
		stopChannel:    make(chan bool),
	}
	go func() {
		proxyCommandLineConfiguration.stopChannel <- true
	}()
	proxyCommandLineConfiguration.main()
}

func TestProxyEmptyConfig(t *testing.T) {
	fileObj, _ := ioutil.TempFile("", "gotest")
	filename := fileObj.Name()
	defer os.Remove(filename)

	ioutil.WriteFile(filename, []byte(`{}`), os.FileMode(0666))
	proxyCommandLineConfiguration.configFileName = filename
	proxyCommandLineConfiguration.pprofaddr = "0.0.0.0:0"
	go func() {
		proxyCommandLineConfiguration.stopChannel <- true
	}()
	main()
}

func TestProxyOkLoading(t *testing.T) {
	fileObj, _ := ioutil.TempFile("", "gotest")
	filename := fileObj.Name()
	defer os.Remove(filename)
	ioutil.WriteFile(filename, []byte(`{"StatsDelay": "1m", "ForwardTo":[{"Type":"csv", "Filename":"/tmp/acsvfile"}], "ListenFrom":[{"Type":"carbon", "Port":"11616"}]}`), os.FileMode(0666))
	myProxyCommandLineConfiguration := proxyCommandLineConfigurationT{
		configFileName: filename,
		logDir:         os.TempDir(),
		logMaxSize:     1,
		logMaxBackups:  0,
		stopChannel:    make(chan bool),
	}
	go func() {
		myProxyCommandLineConfiguration.stopChannel <- true
	}()
	myProxyCommandLineConfiguration.main()
}

func TestProxyListenerError(t *testing.T) {
	fileObj, _ := ioutil.TempFile("", "gotest")
	filename := fileObj.Name()
	defer os.Remove(filename)

	ioutil.WriteFile(filename, []byte(`{"StatsDelay": "1m", "ListenFrom":[{"Type":"carbon"}, {"Type":"carbon"}]}`), os.FileMode(0666))
	myProxyCommandLineConfiguration := proxyCommandLineConfigurationT{
		configFileName: filename,
		logDir:         os.TempDir(),
		logMaxSize:     1,
		logMaxBackups:  0,
		stopChannel:    make(chan bool),
	}
	go func() {
		myProxyCommandLineConfiguration.stopChannel <- true
	}()
	myProxyCommandLineConfiguration.main()
}

func TestProxyForwardError(t *testing.T) {
	fileObj, _ := ioutil.TempFile("", "gotest")
	filename := fileObj.Name()
	defer os.Remove(filename)

	ioutil.WriteFile(filename, []byte(`{"StatsDelay": "1m", "ForwardTo":[{"Type":"carbon", "Host":"192.168.100.108", "Timeout": "1s"}]}`), os.FileMode(0666))
	myProxyCommandLineConfiguration := proxyCommandLineConfigurationT{
		configFileName: filename,
		logDir:         os.TempDir(),
		logMaxSize:     1,
		logMaxBackups:  0,
		stopChannel:    make(chan bool),
	}
	go func() {
		myProxyCommandLineConfiguration.stopChannel <- true
	}()
	myProxyCommandLineConfiguration.main()
}

func TestGetLogrusFormatter(t *testing.T) {
	myProxyCommandLineConfiguration := proxyCommandLineConfigurationT{
		logJSON: true,
	}
	_, ok := myProxyCommandLineConfiguration.getLogrusFormatter().(*log.JSONFormatter)
	assert.True(t, ok)
}

func TestGetLogrusOutput(t *testing.T) {
	myProxyCommandLineConfiguration := proxyCommandLineConfigurationT{
		logDir: "-",
	}
	assert.Equal(t, os.Stdout, myProxyCommandLineConfiguration.getLogrusOutput())
}

func TestProxyUnknownForwarder(t *testing.T) {
	fileObj, _ := ioutil.TempFile("", "gotest")
	filename := fileObj.Name()
	defer os.Remove(filename)

	ioutil.WriteFile(filename, []byte(`{"StatsDelay": "1m", "ForwardTo":[{"Type":"unknown"}]}`), os.FileMode(0666))
	myProxyCommandLineConfiguration := proxyCommandLineConfigurationT{
		configFileName: filename,
		logDir:         os.TempDir(),
		logMaxSize:     1,
		logMaxBackups:  0,
		stopChannel:    make(chan bool),
	}
	go func() {
		myProxyCommandLineConfiguration.stopChannel <- true
	}()
	myProxyCommandLineConfiguration.main()
}

func TestProxyUnknownListener(t *testing.T) {
	fileObj, _ := ioutil.TempFile("", "gotest")
	filename := fileObj.Name()
	defer os.Remove(filename)

	ioutil.WriteFile(filename, []byte(`{"StatsDelay": "1m", "ListenFrom":[{"Type":"unknown"}]}`), os.FileMode(0666))
	myProxyCommandLineConfiguration := proxyCommandLineConfigurationT{
		configFileName: filename,
		logDir:         os.TempDir(),
		logMaxSize:     1,
		logMaxBackups:  0,
		stopChannel:    make(chan bool),
	}
	go func() {
		myProxyCommandLineConfiguration.stopChannel <- true
	}()
	myProxyCommandLineConfiguration.main()
}
