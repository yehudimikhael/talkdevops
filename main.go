package main

import (
	"time"
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)


func main(){

	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 5*time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL: "https://google.com.br",
	})

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics 

	for res := range attacker.Attack(targeter, rate, duration, "Attack!"){
		metrics.Add(res)
	}

	metrics.Close()

	fmt.Println("Resultado: ", metrics.Latencies.P99)
}