package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var quantProcessos int

	fmt.Print("Insira o número de processos desejados: ")
	fmt.Scan(&quantProcessos)
	
	if (quantProcessos >= 2) {
	    
	    pN := criaMatriz(quantProcessos, 10)
    	contadorFrequenciaProcessos(pN, quantProcessos)
    
    	eventos := criaMatriz(4, 1)
    
    	var acao string
    	for {
    
    		mostrarMatrizProcessos(pN, quantProcessos)
    
    		fmt.Println()
    		fmt.Println(" * Novo evento *")
    		var processoOrigem, tempoOrigem, processoDestino, tempoDestino int
    
    		fmt.Print("Processo de origem: ")
    		fmt.Scan(&processoOrigem)
    
    		if processoOrigem >= 0 && processoOrigem <= quantProcessos {
    
    			fmt.Print("Tempo de origem: ")
    			fmt.Scan(&tempoOrigem)
    
    			if tempoOrigem >= 0 && tempoOrigem <= (tempoOrigem*9) && tempoOrigem % (pN[1][processoOrigem] - pN[0][processoOrigem]) == 0 {
    
    				fmt.Print("Processo de destino: ")
    				fmt.Scan(&processoDestino)
    
    				if processoDestino >= 0 && processoDestino <= quantProcessos {
    
    					fmt.Print("Tempo de destino: ")
    					fmt.Scan(&tempoDestino)
    
    					if tempoDestino >= 0 && tempoDestino <= (tempoDestino*9) && tempoDestino % (pN[1][processoDestino] - pN[0][processoDestino]) == 0 {
    
    						novoTempoDestino := 0
    
    						if tempoOrigem < tempoDestino {
    							novoTempoDestino = 0
    						} else {
    							novoTempoDestino = ajustaContadorProcesso(pN, tempoOrigem, processoDestino, tempoDestino)
    						}
    
    						evento := []int{processoOrigem, tempoOrigem, processoDestino, tempoDestino, novoTempoDestino}
    						eventos = append(eventos, evento)
    						mostrarUltimoEvento(eventos)
    
    					} else {
    				        fmt.Println("Tempo de destino ", tempoDestino, " é inválido para o processo ", processoDestino, "!")
    					}
    
    				} else {
    			        fmt.Println("O ID ", processoDestino, " é inválido!")
    				}
    
    			} else {
    				fmt.Println("Tempo de origem ", tempoOrigem, " é inválido para o processo ", processoOrigem, "!")
    			}
    
    		} else {
    			    fmt.Println("O ID ", processoOrigem, " é inválido!")
    		}
    
    		fmt.Print("Deseja sair? (S/N): ")
    		fmt.Scan(&acao)
    
    		if acao == "S" || acao == "s" {
    			break
    		}
    
    	}
	    
	} else {
	    fmt.Println("Erro na entrada da quantidade de processos! Programa encerrado.")
	}

}

func criaMatriz(colunas int, linhas int) [][]int {

	m := make([][]int, linhas)

	for i := 0; i < linhas; i++ {

		m[i] = make([]int, colunas)
		for j := 0; j < colunas; j++ {
			m[i][j] = 0
		}

	}

	return m

}

func contadorFrequenciaProcessos(pN [][]int, quantProcessos int) {

	for j := 0; j < quantProcessos; j++ {

		frequencia := rand.Intn(10) + 1
		for i := 0; i < 10; i++ {
			pN[i][j] = i * frequencia
		}

	}

}

func mostrarMatrizProcessos(pN [][]int, quantProcessos int) {

	fmt.Println()

	for i := 0; i < quantProcessos; i++ {
		fmt.Print("p", i, "\t")
	}
	fmt.Println()

	for i := 0; i < 10; i++ {

		for j := 0; j < quantProcessos; j++ {
			fmt.Print(pN[i][j], "\t")
		}

		fmt.Println()

	}

}

func ajustaContadorProcesso(pN [][]int, tempoOrigem, processoDestino, tempoDestino int) int {

	frequencia := pN[1][processoDestino] - pN[0][processoDestino]
	var valorTrocado int

	for i := 0; i < 10; i++ {

		if tempoDestino == pN[i][processoDestino] {

			pN[i][processoDestino] = tempoOrigem + frequencia
			valorTrocado = pN[i][processoDestino]

		} else if tempoDestino < pN[i][processoDestino] {
			pN[i][processoDestino] = pN[i-1][processoDestino] + frequencia
		}

	}

	return valorTrocado

}

// Mostra 
func mostrarUltimoEvento(eventos [][]int) {

	ultimoEvento := len(eventos) - 1

	fmt.Println()
	fmt.Print("O processo ", eventos[ultimoEvento][0], " enviou no tempo ", eventos[ultimoEvento][1])
	fmt.Print(" para o processo ", eventos[ultimoEvento][2], " que recebeu no tempo ", eventos[ultimoEvento][3])
	fmt.Println()

	if eventos[ultimoEvento][4] != 0 {

		frequencia := eventos[ultimoEvento][4] - eventos[ultimoEvento][1]
		fmt.Print("Foi necessário trocar o tempo ", eventos[ultimoEvento][3], " pelo tempo ", eventos[ultimoEvento][4])
		fmt.Print(" (tempo de envio ", eventos[ultimoEvento][1], " mais a frequência do destino ", frequencia, ")")
		fmt.Println()

	}

	fmt.Println()

}