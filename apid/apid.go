// chain project main.go
package main

import (
	"bufio"
	"chain/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	//	"strconv"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var bcServer chan []*model.Block
var blockchain *model.Blockchain

// 服务器端
func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandle)
	http.HandleFunc("/blockchain/write", blockchainWriteHandle)
	http.ListenAndServe("localhost:8866", nil)
}

func blockchainGetHandle(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchainWriteHandle(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	mutex.Lock()
	blockchain.SendData(blockData)
	mutex.Unlock()
	blockchainGetHandle(w, r)
}

// 命令行模拟p2p连接
func main() {
	mutex.Lock()
	blockchain = model.NewBlockchain()
	mutex.Unlock()
	// 运行服务器端
	//	run()

	bcServer = make(chan []*model.Block)
	server, err := net.Listen("tcp", "localhost:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "Enter a new GenesisData:\n")

	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			data := scanner.Text()

			mutex.Lock()
			blockchain.SendData(data)
			mutex.Unlock()

			bcServer <- blockchain.Blocks
			io.WriteString(conn, "\n Enter a new Data:")

		}
	}()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			output, err := json.Marshal(blockchain)

			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn, string(output))
			io.WriteString(conn, "\n")
		}
	}()

	for _ = range bcServer {
		fmt.Println(blockchain)
	}
}
