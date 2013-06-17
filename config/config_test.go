package config

import (
	"testing"
)

func TestSqrt1(t *testing.T) {
	v := Sqrt(16)
	if v != 4 {
		t.Errorf("Sqrt(16) failed. Got %v, expected 4.", v)
	}
}

func TestInit(t *testing.T) {
	result := Init()
	if !result {
		t.Errorf("Init() failed. Got %v, expected true.", result)
	}
}

func TestGet(t *testing.T) {
	Init()
	conf := Get("test")
	if conf != "test config" {
		t.Errorf("Get() failed. Got %v, expected \"test config\".", conf)
	}
}

func TestSet(t *testing.T) {
	Init()
	Set("test", "test set config")
	conf = Get("test")
	if conf != "test set config" {
		t.Errorf("Set() failed. Got %v, expected \"test set config\".", conf)
	}
}
