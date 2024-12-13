import websocket
import json

def on_message(ws, message):
    print(f"Message reçu : {message}")

def on_error(ws, error):
    print(f"Erreur : {error}")

def on_close(ws, close_status_code, close_msg):
    print("Connexion fermée")

def on_open(ws):
    print("Connexion ouverte")
    message = json.dumps({
        "action": "TakeDecision",
        "message": "Votre message ici"
    })
    ws.send(message)

ws = websocket.WebSocketApp("ws://localhost:3000/ws",
                            on_open=on_open,
                            on_message=on_message,
                            on_error=on_error,
                            on_close=on_close)

ws.run_forever()