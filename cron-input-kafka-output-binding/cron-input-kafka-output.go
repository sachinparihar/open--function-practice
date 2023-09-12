package bindings

import (
	"fmt"
	"log"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
)

func HandleCronInput(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	var greeting []byte
	if in != nil {
		log.Printf("binding - Data: %s", in)
		greeting = in
	} else {
		log.Print("binding - Data: counting 1 to n")
		n := 10
		for i := 1; i <= n; i++ {
			greeting = append(greeting, []byte(fmt.Sprintf("%d\n", i))...)
		}
	}

	_, err := ctx.Send("kafka-server", greeting)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return ctx.ReturnOnInternalError(), err
	}

	return ctx.ReturnOnSuccess(), nil
}
