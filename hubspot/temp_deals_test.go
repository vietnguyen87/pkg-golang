package hubspot

import (
	"fmt"
	"testing"
)

func TestDeal(t *testing.T) {
    c := NewClient(NewClientConfig("", ""))
    r, _ := c.Deals().Get("9361327381")
    
    fmt.Println(r)
}