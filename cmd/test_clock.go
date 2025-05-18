package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AccessLogRequest struct {
	EmployeeID   int    `json:"employee_id"`
	AccessTime   string `json:"access_time"`
	Direction    string `json:"direction"`     // "in" or "out"
	GateType     string `json:"gate_type"`     // "entry" or "exit"
	GateName     string `json:"gate_name"`     // e.g., "AZ_door_1"
	AccessResult string `json:"access_result"` // e.g., "Admitted"
}

func testClockInsert() {
	total := 4000
	url := "http://localhost:8080/api/clock"

	for i := 0; i < total; i++ {
		request := AccessLogRequest{

			// -----更改成真實資料------ //

			EmployeeID: (i % 100) + 1, // 模擬 100 位員工，ID 從 1 開始

			// ----------------------- //

			AccessTime:   time.Now().Format(time.RFC3339),
			Direction:    "in", // or "out"
			GateType:     "entry",
			GateName:     fmt.Sprintf("AZ_door_%d", i%10+1), // 模擬 10 個門
			AccessResult: "Admitted",
		}

		payload, err := json.Marshal(request)
		if err != nil {
			fmt.Printf("JSON Marshal error at %d: %v\n", i, err)
			continue
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Printf("Request failed at %d: %v\n", i, err)
			continue
		}
		resp.Body.Close()

		if i%500 == 0 {
			fmt.Printf("✅ Sent %d requests\n", i)
		}
	}

	fmt.Println("🎉 Done sending access log test requests.")
}
