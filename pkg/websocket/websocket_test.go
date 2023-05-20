package websocket

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebSocket_Connect(t *testing.T) {
	// Cria um servidor mock
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Error upgrading connection to websocket: %v", err)
		}
		defer ws.Close()

		// Envia uma mensagem para o cliente assim que a conexão é estabelecida
		err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, world!"))
		if err != nil {
			t.Fatalf("Error writing message to websocket: %v", err)
		}
	}))
	defer server.Close()

	// Cria uma instância de WebSocket com o URL do servidor mock
	ws := WebSocket{URL: "ws" + strings.TrimPrefix(server.URL, "http")}

	// Estabelece uma conexão com o servidor
	conn := ws.Connect()
	if conn == nil {
		t.Fatalf("Expected a websocket connection, but got nil")
	}
	defer conn.Close()

	// Lê a mensagem enviada pelo servidor
	_, message, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Error reading message from websocket: %v", err)
	}

	// Verifica se a mensagem recebida é a esperada
	if string(message) != "Hello, world!" {
		t.Fatalf("Expected message \"Hello, world!\", but got \"%s\"", string(message))
	}
}

func TestWebSocket_WriteMessage(t *testing.T) {
	// Cria um servidor mock
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Error accepting websocket connection: %v", err)
		}
		defer ws.Close()

		// Lê a mensagem enviada pelo cliente
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("Error reading message from websocket: %v", err)
		}

		// Verifica se a mensagem recebida tem o tamanho esperado
		expectedSize := len("Hello, world!")
		if len(p) != expectedSize {
			t.Fatalf("Expected message size of %d, but got %d", expectedSize, len(p))
		}

		// Envia uma mensagem de resposta para o cliente
		err = ws.WriteMessage(messageType, p)
		if err != nil {
			t.Fatalf("Error writing message to websocket: %v", err)
		}
	}))
	defer server.Close()

	// Cria uma instância de WebSocket com o URL do servidor mock
	ws := WebSocket{URL: "ws" + strings.TrimPrefix(server.URL, "http")}

	// Estabelece uma conexão com o servidor
	conn := ws.Connect()
	if conn == nil {
		t.Fatalf("Expected a websocket connection, but got nil")
	}
	defer conn.Close()

	// Envia uma mensagem para o servidor
	ws.WriteMessage(conn, []byte("Hello, world!"))

	// Lê a mensagem de resposta enviada pelo servidor
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Error reading message from websocket: %v", err)
	}

	// Verifica se a mensagem recebida é a esperada
	expectedMessage := "Hello, world!"
	if string(p) != expectedMessage {
		t.Fatalf("Expected message \"%s\", but got \"%s\"", expectedMessage, string(p))
	}

	// Verifica se o tipo da mensagem recebida é o mesmo tipo da mensagem enviada
	expectedType := websocket.TextMessage
	if messageType != expectedType {
		t.Fatalf("Expected message type %d, but got %d", expectedType, messageType)
	}
}
