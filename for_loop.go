package main

import "fmt"

func main() {
	hour := 24
	minute := 60
	second := 60

	for h := 0; h < 24; hour-- {
		fmt.Println(hour == 0)
		if hour == 0 {
			break
		}

		for m := 0; m < minute; minute-- {
			if minute == 0 {
				minute = 60
			}

			for s := 0; s < second; second-- {

				if second == 0 {
					second = 60
				}

				fmt.Printf("%d:%d:%d\n", hour, minute, second)
			}
		}
	}
}
