package main

import (
	"time"
	"fmt" //FMT -> O Pacato FMT implementa  asf unções de I/O(Entrada e Saída) análogas ao Printf e o Scanf de C em GO
	vegeta "github.com/tsenart/vegeta/v12/lib"
)


func main(){

	rate := vegeta.Rate{Freq: 100, Per: time.Second} //Fazer 100 request a cada segundo
	duration := 5*time.Second // E vamos fazer isso por 5s
	// Definir nosso alvo 
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL: "https://google.com.br",
	})

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics 
	
	// Vamos realizar as requisições
	for res := range attacker.Attack(targeter, rate, duration, "Attack!"){
		metrics.Add(res) //Atribuindo o resultado ao objeto
	}

	metrics.Close()
	
	fmt.Println("Resultado: ", metrics.Latencies.P99) 
}
