// http 服务
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
}

type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

var clients = make(map[*Client]bool)
var broadcast = make(chan Message)

type Server struct {
	Host string
	Port string
	Data string
}

var srv http.Server

// 提供服务
func (s *Server) Serve() {
	mx := http.NewServeMux()
	mx.Handle("/file", http.FileServer(http.Dir(s.Data)))

	mx.HandleFunc("/socket", socketHandler)
	go sendMessage()

	mx.HandleFunc("/sayhelloName", sayhelloName) //设置访问的路由

	mx.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Printf("'%s','%s'\n", remoteAddr(r), path)
		// 静态资源
		if isStaticFile(path) {
			static := "static" + path
			http.ServeFile(w, r, static)
			return
		}
		if path == "/" {
			static := "static/index.html"
			http.ServeFile(w, r, static)
			return
		}
		// 404
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "error")
	})

	mx.HandleFunc("/upload", uplad)

	srv = http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: mx,
	}

	// strconv.Itoa(8080)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// shutdown server
func (s *Server) Shutdown() {
	log.Println("server is stopping")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		select {
		case <-ctx.Done():
			// 发生异常，超时未关闭
			log.Println("timeout of 5 seconds.")
		default:
			// 发生异常，未知原因未正常关闭
			log.Fatalf("Server shutdown failed: %v\n", err)
		}
		return
	}
	log.Println("Server shutdown gracefully")
}

var suffixArray = []string{".css", ".js", ".html"}

func isStaticFile(path string) bool {
	for _, suffix := range suffixArray {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// 接收 socket 请求
func socketHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("'%s','%s'\n", remoteAddr(r), r.URL.Path)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	client := &Client{conn: conn}
	clients[client] = true

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("messageType: %d, Received: %s", messageType, message)
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("Error during message reading:", err)
			delete(clients, client)
			break
		}
		broadcast <- msg
	}
}

// 转发 socket 消息
func sendMessage() {
	for {
		msg := <-broadcast
		for client := range clients {
			if client.conn != nil {
				err := client.conn.WriteJSON(msg)
				if err != nil {
					log.Println(err)
					client.conn.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	//这些信息是输出到服务器端的打印信息
	log.Printf("path: %s, scheme: %s, url_long: %s", r.URL.Path, r.URL.Scheme, r.Form["url_long"])
	for k, v := range r.Form {
		log.Printf("key: %s, val: %s", k, strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}

func uplad(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	log.Printf("'%s','%s', '%s'\n", remoteAddr(r), r.URL.Path, name)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "ok")
}

// proxy_set_header X-Real-IP $remote_addr;
// proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
func remoteAddr(r *http.Request) string {
	header := r.Header
	remoteAddr := header.Get("X-Forwarded-For")
	if remoteAddr != "" {
		return remoteAddr
	}
	return r.RemoteAddr
}
