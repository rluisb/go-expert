package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//Criando e escrevendo em arquivos com bytes ou strings
	f, err := os.Create("arquivo.txt")
	handleError(err)

	// Quando lidamos com objetos, o ideal seria que ustilizassemos o Write para escrever os bytes no arquivo
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))

	//Utilizamos WriteString quando temos dados simples como strings
	//tamanho, err := f.WriteString("Hello, World!")
	handleError(err)

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d Bytes\n", tamanho)

	//Por se tratar de uma operação IO, devemos fechar sempre
	f.Close()

	//Leitura de arquivos
	//Ao invés de abrir o arquivo para leitura, utilizamos o os.ReadFile para ler todo o conteúdo
	arquivo, err := os.ReadFile("arquivo.txt")
	handleError(err)

	// Imprimindo o conteudo do arquivo, lembrando que sempre que lermos algum arquivo, teremos o conteúdo em bytes, por isso convertemos para string
	fmt.Println("Conteudo do arquivo: " + string(arquivo))

	// Leitura de arquivo em chunks
	//Primeiramente usamos o os.Open para abrir o arquivo
	arquivoGigante, err := os.Open("arquivo.txt")
	//Tratamos se erro
	handleError(err)

	//Criamos um reader para ler o arquivo e passamos os bytes do arquivo aberto previamente
	reader := bufio.NewReader(arquivoGigante)

	//Criamos um buffer para indicar de quantos em quantos bytes queremos ler o arquivo, nesse caso 10 bytes
	buffer := make([]byte, 3)

	// Iteramos com loop infinito afim de ler todo o arquivo
	for {
		// Executamos a leitura chamando o metodo Read do reader passando o buffer para ele
		n, err := reader.Read(buffer)
		//Tratamos o erro
		if err != nil {
			break
		}

		// Imprimimos o valor de cada chunk de bytes lido
		// o n indica a posição onde está sendo feita a leitura
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	handleError(err)
}
