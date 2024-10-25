package topic1

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPrintNumbers(t *testing.T) {

	originalStdout := os.Stdout //terminal a qayitmaq ucun evvelki veziyyeti
	r, w, _ := os.Pipe()        // boru yaradiriq (file yaradib yazmaq kimi )
	os.Stdout = w

	n := 5
	delay := 100 * time.Millisecond
	go PrintNumbers(n, delay)

	time.Sleep(time.Second)

	w.Close()                  //Boruya artiq data yazmamamq ucun baglagiriq
	os.Stdout = originalStdout // normal veziyyete qaydiriq
	var buf bytes.Buffer
	buf.ReadFrom(r) // w ile boruya yazdigmiz parametrlerin funksiyadki qiymetlerini  r ile oxuyuruq sonra string e cevririk
	actual := buf.String()
	var expected string
	for i := 1; i <= n; i++ {
		expected += fmt.Sprintf("%d\n", i) //excepted value -lari bir yere yigiib string kimi birlesdiriik
	}
	if actual != expected {
		t.Errorf("Excepted value is diffrent than actual ==> EXCEPTED %s,   ACTUAL %s", expected, actual)
	}
}
