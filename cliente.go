package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net"
  "os"
)

const SEND_BUFFER_SIZE = 2048

//Função simples para checar e reportar algum erro que venha a ocorrer
func checkErrorClient(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
    os.Exit(1)
  }
}

/* TODO: client()
 * Abrir socket e enviar mensagem de stdin.
*/
func client(server_ip string, server_port string) {

  //Leitura da mensagem (arquivo) do terminal
  //ioutil.ReadAll lê até chegar a um EOF (end of file)
  str, err := ioutil.ReadAll(os.Stdin)
  checkErrorClient(err)

  //Tradução das strings recebidas (server_ip e server_port) para um endereço TCP
  addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", server_ip, server_port))
  checkErrorClient(err)

  //Criação do socket e estabelecimento de conexão com o servidor
  conn, err := net.DialTCP("tcp", nil, addr)
  checkErrorClient(err)

  //Fecha a conexão após a função finalizar sua execução
  defer conn.Close()

  //Enviando a mensagem do cliente para o servidor através da conexão estabelecida
  conn.Write(str)

  os.Exit(0)
}

// Main obtém argumentos da linha de comando e chama função client
func main() {
  if len(os.Args) != 3 {
    log.Fatal("Uso: h./cliente [IP servidor] [porta servidor] < [arquivo mensagem]")
  }
  server_ip := os.Args[1]
  server_port := os.Args[2]
  client(server_ip, server_port)
}
