package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net"
  "os"
)

const RECV_BUFFER_SIZE = 2048

//Função simples para checar e reportar algum erro que venha a ocorrer
func checkErrorServer(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
    os.Exit(1)
  }
}

/* TODO: server()
 * Abra socket e espere o cliente conectar
 * Imprima a mensagem recebida em stdout
*/
func server(server_port string) {

  //Tradução das strings recebidas (server_ip e server_port) para um endereço TCP
  addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%s", server_port))
  checkErrorServer(err)

  //Inicia a escuta de pedidos de conexão pelo endereço TCP acima (addr)
  listener, err := net.ListenTCP("tcp", addr)
  checkErrorServer(err)

  //Fecha a escuta de pedidos após a função finalizar sua execução
  defer listener.Close()

  for {

    //Gera uma conexão para todos os clientes que desejam se conectar
    conn, err := listener.AcceptTCP()

    if err != nil {
      continue
    }

    //Lendo da conexão a resposta referente ao pedido efetuado pelo cliente e gravando numa variável
    str, err := ioutil.ReadAll(conn)
    checkErrorServer(err)

    fmt.Print(string(str))
  }
}

// Main obtém argumentos da linha de comando e chama a função servidor
func main() {
  if len(os.Args) != 2 {
    log.Fatal("Uso: ./servidor [porta servidor]")
  }
  server_port := os.Args[1]
  server(server_port)
}
