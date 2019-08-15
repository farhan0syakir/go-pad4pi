package pad4pi

import (
	"log"
	"os"
	"os/signal"
	"reflect"
	"time"

	"github.com/warthog618/gpio"
)

const (
	J8p27 = iota
	J8p28
	J8p3
	J8p5
	J8p7
	J8p29
	J8p31
	J8p26
	J8p24
	J8p21
	J8p19
	J8p23
	J8p32
	J8p33
	J8p8
	J8p10
	J8p36
	J8p11
	J8p12
	J8p35
	J8p38
	J8p40
	J8p15
	J8p16
	J8p18
	J8p22
	J8p37
	J8p13
)

func GetPin(pin int) uint8 {
	pinAlias := []int{
		0,
		0,
		0,
		J8p3,
		0,
		J8p5,
		0,
		J8p7,
		J8p8,
		0,
		J8p10,
		J8p11,
		J8p12,
		J8p13,
		0,
		J8p15,
		J8p16,
		0,
		J8p18,
		J8p19,
		0,
		J8p21,
		J8p22,
		J8p23,
		J8p24,
		0,
		J8p26,
		J8p27,
		J8p28,
		J8p29,
		0,
		J8p31,
		J8p32,
		J8p33,
		0,
		J8p35,
		J8p36,
		J8p37,
		J8p38,
		0,
		J8p40,
	}
	return uint8(pinAlias[pin])
}

type Keypad struct {
	RowPins           []*gpio.Pin
	ColPins           []*gpio.Pin
	MyKeypad          [][]string
	ticker            *time.Ticker
	lastInterruptTime int64
}

var defaultKeypad *Keypad

func getField(pin *gpio.Pin, field string) int {
	r := reflect.ValueOf(pin)
	f := reflect.Indirect(r).FieldByName(field)
	// return f.Interface().(uint8)
	return int(f.Uint())
}

func getKey(pin *gpio.Pin) string {
	keypad := defaultKeypad
	rowVal := 0
	for i, s := range keypad.RowPins {
		if getField(pin, "pin") == getField(s, "pin") {
			rowVal = i
			break
		}
	}
	colVal := 0
	for i, s := range keypad.ColPins {
		s.High()
		if pin.Read() {
			colVal = i
			s.Low()
			break
		}
		s.Low()

	}
	return keypad.MyKeypad[rowVal][colVal]

}

func getMilis() int64 {
	now := time.Now()

	unixNano := now.UnixNano()
	umillisec := unixNano / 1000000

	// log.Println(unixNano)
	// log.Println(umillisec)
	return umillisec
}

func onKeyResponse(pin *gpio.Pin) {
	interruptTime := getMilis()
	if interruptTime-defaultKeypad.lastInterruptTime < 200 {
		return
	}
	log.Println(getKey(pin))
	defaultKeypad.lastInterruptTime = interruptTime

}

func initRow(RowPins []*gpio.Pin) {
	for _, p := range RowPins {

		err := p.Watch(gpio.EdgeFalling, onKeyResponse)
		if err != nil {
			panic(err)
		}

	}
	log.Println("suukses")
}

func (keypad *Keypad) Stop() {
	gpio.Close()

	for _, p := range keypad.RowPins {
		p.Unwatch()
	}
}
func NewKeypad(rowPinsInt []int, ColPinsInt []int, MyKeypad [][]string) *Keypad {
	err := gpio.Open()
	if err != nil {
		panic(err)
	}
	RowPins := make([]*gpio.Pin, len(rowPinsInt))
	for i, p := range rowPinsInt {
		pinR := gpio.NewPin(GetPin(p))
		pinR.Input()
		pinR.PullUp()
		RowPins[i] = pinR
	}
	ColPins := make([]*gpio.Pin, len(ColPinsInt))

	for i, p := range ColPinsInt {
		pinC := gpio.NewPin(GetPin(p))
		pinC.Output()
		pinC.Low()
		ColPins[i] = pinC

	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	// defer func() {
	defaultKeypad = &Keypad{
		RowPins:           RowPins,
		ColPins:           ColPins,
		MyKeypad:          MyKeypad,
		lastInterruptTime: 0,
	}
	go initRow(RowPins)
	return defaultKeypad
	// }()

}
