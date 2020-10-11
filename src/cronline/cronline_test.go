package cronline

import (
	"testing"
)

func TestGetCommand(t *testing.T) {
	if GetCommand("* * * * * date") != "date" {
		t.Error("GetCommand Error", GetCommand("* * * * * date"))
	}
}

func TestGetCommand2(t *testing.T) {
	if GetCommand("* * * * * date") != "date" {
		t.Error("GetCommand Error", GetCommand("* */4 * * * date"))
	}
}

func TestGetCommand3(t *testing.T) {
	if GetCommand("* * * * * date") != "date" {
		t.Error("GetCommand Error", GetCommand("* */4 10-50 * * date"))
	}
}

func TestGetCommand4(t *testing.T) {
	if GetCommand("* */4 10-50 * * date -d `1 day ago`") != "date -d `1 day ago`" {
		t.Error("GetCommand Error", GetCommand("* */4 10-50 * * date -d `1 day ago`"))
	}
}

func TestGetCommand5(t *testing.T) {
	if GetCommand("@daily date -d `1 day ago`") != "date -d `1 day ago`" {
		t.Error("GetCommand Error", GetCommand("@daily date -d `1 day ago`"))
	}
}

func TestGetCommand6(t *testing.T) {
	if GetCommand("@every 1h30 date -d `1 day ago`") != "date -d `1 day ago`" {
		t.Error("GetCommand Error", GetCommand("@every 1h30 date -d `1 day ago`"))
	}
}

func TestGetCron(t *testing.T) {
	if GetCron("* */4 10-50 * * date -d `1 day ago`") != "* */4 10-50 * * " {
		t.Error("GetCron Error", GetCron("* */4 10-50 * * date -d `1 day ago`"))
	}
}

func TestGetCron2(t *testing.T) {
	if GetCron("* * * * * date -d `1 day ago`") != "* * * * * " {
		t.Error("GetCron Error", GetCron("* * * * * date -d `1 day ago`"))
	}
}

func TestGetCron3(t *testing.T) {
	if GetCron("@daily date -d `1 day ago`") != "@daily " {
		t.Error("GetCron Error", GetCron("@daily date -d `1 day ago`"))
	}
}

func TestGetCron4(t *testing.T) {
	if GetCron("@every 1h30 date -d `1 day ago`") != "@every 1h30 " {
		t.Error("GetCron Error", GetCron("@every 1h30 date -d `1 day ago`"))
	}
}

func TestGetCron5(t *testing.T) {
	if GetCron("every 1h30 date -d `1 day ago`") != "" {
		t.Error("GetCron Error", GetCron("every 1h30 date -d `1 day ago`"))
	}
}

func TestGetCommandOfEvery1(t *testing.T) {
	if GetCommandOfEvery("@every 1h30 date -d `1 day ago`") != "date -d `1 day ago`" {
		t.Error("GetCommandOfEvery Error", GetCommandOfEvery("@every 1h30 date -d `1 day ago`"))
	}
}

func TestGetCommandOfEvery2(t *testing.T) {
	if GetCommandOfEvery("every 1h30 date -d `1 day ago`") != "" {
		t.Error("GetCommandOfEvery Error", GetCommandOfEvery("@every 1h30 date -d `1 day ago`"))
	}
}
