package main

import(
"github.com/stianeikeland/go-rpio"
"time"
"fmt"
)
func Initiate(){
        pin1 = rpio.Pin(17)
        pin2 = rpio.Pin(18)
        pin3 = rpio.Pin(27)
        pin4 = rpio.Pin(22)
        pin5 = rpio.Pin(6)
        pin6 = rpio.Pin(12)
        pin7 = rpio.Pin(23)
        pin8 = rpio.Pin(24)
        pin1.Output()
        pin2.Output()
        pin3.Output()
        pin4.Output()
        pin5.Output()
        pin6.Output()
        pin7.Output()
        pin8.Output()
}

func Lightson(){
		lit = true
		pin5.Low()
		pin6.High()
		pin7.High()
		pin8.Low()
}

func Lightsoff(){
		lit = false
		pin5.Low()
		pin6.Low()
		pin7.Low()
		pin8.Low()
}

func Rotate_left(wait int){
		if  lit == false{
			pin5.Low()
			pin6.High()
		}
      fmt.Println("Rotating left")
			pin1.Low()
			pin2.High()
			pin3.High()
			pin4.Low()
			time.Sleep(wait*time.Millisecond)
			reset()
		if lit == false{
		pin5.Low()
		pin6.Low()
		}
}

func Rotate_right(wait int){
		if lit == false{
		pin7.High()
		pin8.Low()
		}
        fmt.Println("Rotating right")
			pin1.High()
			pin2.Low()
			pin3.Low()
			pin4.High()
			time.Sleep(wait*time.Millisecond)
			reset();
		if lit == false{
		pin7.Low()
		pin8.Low()
		}
}

func Forward(wait int){
      fmt.Println("Going forward")
			pin1.High()
			pin2.Low()
			pin3.High()
			pin4.Low()
			time.Sleep(wait*time.Millisecond)
			reset()
}

func Backwards(wait int){
      fmt.Println("Going backwards")
			pin1.Low()
			pin2.High()
			pin3.Low()
			pin4.High()
			time.Sleep(wait*time.Millisecond)
			reset()
		}
}

func reset(){
        pin1.Low()
        pin2.Low()
        pin3.Low()
        pin4.Low()
}
