package plugin

import (
	"os"
	"testing"

	cln4go "github.com/vincenzopalazzo/cln4go/client"
)

func TestCallFistMethod(t *testing.T) {
	path := os.Getenv("CLN_UNIX_SOCKET")
	client, err := cln4go.NewUnix(path)
	if err != nil {
		panic(err)
	}
	response, err := cln4go.Call[map[string]any, map[string]any](client, "hello", map[string]any{})
	if err != nil {
		panic(err)
	}

	message, found := response["message"]
	if !found {
		t.Error("The message is not found")
	}

	if message != "hello from cln4go.template" {
		t.Errorf("message received %s different from expected %s", message, "hello from cln4go.template")
	}
}
