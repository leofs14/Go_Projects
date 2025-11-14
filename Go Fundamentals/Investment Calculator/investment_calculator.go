package main

import (
	"fmt"
	"math"
)

	// Criando variáveis globais
	const inflationRate = 2.5

func main() { //função obrigatória

	// Variáveis
	var investmentAmount, years float64 
	expectedReturnRate := 5.5

	//Entrada de Dados:
	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount) // & é um pointer.

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years) 

	//Contas:
	futureValue, inflationAdjustedValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// inflationAdjustedValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Outputs:
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedIAV := fmt.Sprintf("Future Value(Adjusted To Inflation): %.1f\n", inflationAdjustedValue)
	//fmt.Println("Future Value: ", futureValue)
	//fmt.Printf("Future Value: %.1f\nFuture Value(Adjusted To Inflation): %.1f", futureValue, inflationAdjustedValue)
	// fmt.Printf(`Future Value: %.1f

	// Future Value(Adjusted To Inflation): %.1f`, futureValue, inflationAdjustedValue) // a crase pula a linha do texto
	//fmt.Println("Future Value(Adjusted To Inflation): ", inflationAdjustedValue)
	fmt.Print(formattedFV, formattedIAV)
}

//Função de Texto para Output
func outputText(text string){
	fmt.Print(text)
}

//Função para executar os cálculos
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, iav float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	iav = fv / math.Pow(1+inflationRate/100, years)
	// return fv, iav
	return //outro jeito de retornar variáveis, mas quando tem mais variaveis é melhor usar o outro 
}