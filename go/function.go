package hello

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)


func init() {
	functions.HTTP("hello", HelloWorld)
	functions.CloudEvent("ce", CloudEvent)
}

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello, World!\n")
	panic("Intended panic")
}

func CloudEvent(ctx context.Context, ce cloudevents.Event) error {
	// Do something with event.Context and event.Data (via event.DataAs(foo)).
	//panic("Intended panic")
	e, err := json.Marshal(ce)
	if err != nil {
		return fmt.Errorf("marshalling cloud event: %v", err)
	}
	log.Printf("Got Cloud Event: %v", string(e))


	return errors.New("intended err")
}