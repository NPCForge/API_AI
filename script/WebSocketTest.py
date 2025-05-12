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
        if "token" in message_dict and (action in ["register", "disconnect", "connect", "remove"]):
            token = message_dict["token"]
            print(f"Token mis à jour : {token}")
    except json.JSONDecodeError:
        print("Erreur de décodage JSON du message reçu.")

def on_error(ws, error):
    print(f"Erreur : {error}")

def on_close(ws, close_status_code, close_msg):
    print("Connexion fermée")

def on_open(ws):
    print("Connexion ouverte")
    print("Entrez une commande (Register, Connect, MakeDecision, Disconnect, Remove, NewMessage, CreateEntity) ou 'exit' pour quitter.")

def listen_to_stdin(ws):
    global token, action
    while True:
        user_input = input("Commande : ").strip()

        if user_input.lower() == "exit":
            print("Fermeture de la connexion...")
            ws.close()
            break

        if user_input.lower() == "register":
            message = json.dumps({
                "action": "Register",
                "API_KEY": "VDCAjPZ8jhDmXfsSufW2oZyU8SFZi48dRhA8zyKUjSRU3T1aBZ7E8FFIjdEM2X1d",
                "identifier": "tom_user",
                "password": "password",
                "game_prompt": "A werewolf game but every messages should rime"
            })
        elif user_input.lower() == "connect":
            message = json.dumps({
                "action": "Connect",
                "identifier": "tom_user",
                "password": "password",
            })
        elif user_input.lower() == "createentity":
            if token:
                message = json.dumps({
                    "action": "CreateEntity",
                    "name": "tom_entity",
                    "prompt": "a fisherman named tom_entity",
                    "checksum": "TomChecksum",
                    "token": token,
                })
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")
                continue
        elif user_input.lower() == "makedecision":
            if token:
                message = json.dumps({
                    "action": "MakeDecision",
                    "token": token,
                    "message": "",
                    "checksum": "TomChecksum",
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
        elif user_input.lower() == "newmessage":
            if token:
                message = json.dumps({
                    "action":    "NewMessage",
                    "sender":    "TomChecksum",
                    "receivers": ["TomChecksum"],
                    "message":   "hello tom",
                    "token":     token,
                })
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")
                continue
        else:
            print("Commande inconnue. Essayez 'Register', 'Connect', 'MakeDecision', 'Disconnect', 'Remove', 'NewMessage', 'CreateEntity' 'exit'.")
            continue

        action = user_input.lower()
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