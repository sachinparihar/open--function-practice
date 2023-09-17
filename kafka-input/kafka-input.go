// kafka-input.go
package bindings

import (
	"encoding/json"
	"fmt"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
)

func HandleKafkaInput(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	var msg Message
	err := json.Unmarshal(in, &msg)
	if err != nil {
		fmt.Println("error reading message from Kafka binding", err)
		return ctx.ReturnOnInternalError(), err
	}
	for _, number := range msg.Numbers {
		fmt.Printf("%d\n", number)
	}
	return ctx.ReturnOnSuccess(), nil
}

type Message struct {
	Numbers []int `json:"numbers"`
}
