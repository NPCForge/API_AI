import websocket
import json
import threading

# Variable globale pour stocker le token reçu
token = None
action = ""

def on_message(ws, message):
    global token
    print(f"Message reçu : {message}")
    try:
        message_dict = json.loads(message)
        if "data" in message_dict and (action == "register" or action == "disconnect" or action == "connection"):
            token = message_dict["data"]
            print(f"Token mis à jour : {token}")
    except json.JSONDecodeError:
        print("Erreur de décodage JSON du message reçu.")

def on_error(ws, error):
    print(f"Erreur : {error}")

def on_close(ws, close_status_code, close_msg):
    print("Connexion fermée")

def on_open(ws):
    print("Connexion ouverte")
    print("Entrez une commande (Register, Connection, TakeDecision, Disconnect) ou 'exit' pour quitter.")

def listen_to_stdin(ws):
    global token, action  # Ajout de 'action' à la liste des variables globales
    while True:
        user_input = input("Commande : ").strip()

        if user_input.lower() == "exit":
            print("Fermeture de la connexion...")
            ws.close()
            break

        # Préparation des messages en fonction de la commande
        if user_input.lower() == "register":
            message = json.dumps({
                "action": "Register",
                "API_KEY": "VDCAjPZ8jhDmXfsSufW2oZyU8SFZi48dRhA8zyKUjSRU3T1aBZ7E8FFIjdEM2X1d",
                "checksum": "azerty",
                "name": "tom",
                "prompt": "juste le boss",
            })
        elif user_input.lower() == "connection":
            message = json.dumps({
                "action": "Connection",
                "checksum": "azerty"
            })
        elif user_input.lower() == "takedecision":
            if token:
                message = json.dumps({
                    "action": "TakeDecision",
                    "token": token,
                    "message": "Hello World"
                })
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")
                continue
        elif user_input.lower() == "disconnect":
            if token:
                message = json.dumps({
                    "action": "Disconnect",
                    "token": token
                })
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")
                continue
        else:
            print("Commande inconnue. Essayez 'Register', 'Connection', 'TakeDecision', 'Disconnect', ou 'exit'.")
            continue

        action = user_input.lower()  # Mise à jour correcte de la variable globale 'action'
        ws.send(message)

# Initialisation du WebSocket
ws = websocket.WebSocketApp("ws://localhost:3000/ws",
                            on_open=on_open,
                            on_message=on_message,
                            on_error=on_error,
                            on_close=on_close)

# Lancer le WebSocket dans un thread séparé pour écouter stdin
thread = threading.Thread(target=listen_to_stdin, args=(ws,))
thread.daemon = True
thread.start()

ws.run_forever()
