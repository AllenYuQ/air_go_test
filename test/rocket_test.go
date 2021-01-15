package test

import (
	"encoding/json"
	"fmt"
	"task5/models"
	"task5/pkg/util"
	"testing"
	"time"
)

func TestRocket(t *testing.T) {
	util.StartRocketMQ()
	mlmets := models.ListMlmets()
	fmt.Println(len(mlmets))
	for _, mlmet := range mlmets {
		fmt.Println(mlmet)
		data, _ := json.Marshal(&mlmet)
		fmt.Printf("%s", data)
		fmt.Println()
		util.SendMessageToRocketChan(data, 1, "test")
	}

	time.Sleep(10 * time.Second)
}
