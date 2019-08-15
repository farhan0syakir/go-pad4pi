# go-pad4pi

Interrupt based golang library for reading matrix keypad on raspberry

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Prepare raspberry and keypad. In the example we use the 2x2 raspberry like this
https://robu.in/wp-content/uploads/2019/05/2-x-2-Matrix-4-Push-Button-Keyboard-Module-1.jpg

and we map the pin like this

```
L1 to 35
L2 to 36
R1 to 37
R2 to 38
```

### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
cd example
go get -d
```


## Running the tests

Explain how to run the automated tests for this system

```
go run app.go
```

### Example


```
var myKeypad = [][]string{
	{"1", "2"},
	{"3", "4"},
}

var rowPins = []int{35, 36}
var colPins = []int{37, 38}

func handler(key string) {
	log.Println(key)
}

func main() {
	mykeypad := pad4pi.NewKeypad(rowPins, colPins, myKeypad)
	mykeypad.RegisterKeyPressHandler(handler)
	defer mykeypad.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	defer signal.Stop(quit)

	// In a real application the main thread would do something useful here.
	// But we'll just run for a minute then exit.
	fmt.Println("Watching Pin")
	select {
	case <-time.After(time.Minute):
	case <-quit:
	}

}

```

## Authors

* **Farhan Syakir** - *Initial work* - [github](https://github.com/farhan0syakir)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments
* **Inspiration from pad4pi python** - [github](https://github.com/brettmclean/pad4pi)
* **Special thanks to Warung pintar** 
    * **Sofian HW** [github](https://github.com/sofianhw)
    * **Abda Barias Salam**
    * **Elvandry**



