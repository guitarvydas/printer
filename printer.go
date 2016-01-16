package printer

import (
	"fmt"
	"github.com/guitarvydas/ip"
)

func Print(in <-chan ip.IP) {
	for {
		merged := <-in
		switch merged.Kind {
		case ip.Open:
			fmt.Println("{")
		case ip.Close:
			fmt.Println("}")
		case ip.EOF:
			fmt.Println("===> EOF")
			return
		default:
			fmt.Printf("%s", merged.Data)
		}
	}
}
