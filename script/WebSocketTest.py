import websocket
import json
import threading
import signal
import sys

token = None
action = ""

def on_message(ws, message):
    print(f"Message reçu : {message}")
    try:
        message_dict = json.loads(message)
        if "token" in message_dict and (action in ["register", "disconnect", "connect", "remove"]):
            global token
            token = message_dict["token"]
            print(f"Token mis à jour : {token}")
    except json.JSONDecodeError:
        print("Erreur de décodage JSON du message reçu.")

def on_error(ws, error):
    print(f"Erreur : {error}")

def on_close(ws, code, msg):
    print(f"Connexion fermée (code={code}, msg={msg})")

def on_open(ws):
    print("Connexion ouverte")
    print("Entrez une commande (Register, Connect, CreateEntity, NewMessage, MakeDecision, night, Disconnect, Remove, exit).")

def listen_to_stdin(ws):
    global token, action
    while True:
        try:
            user_input = input("Commande : ").strip()
        except (EOFError, KeyboardInterrupt):
            print("\nFermeture…")
            ws.close()
            break

        if user_input.lower() == "exit":
            print("Fermeture de la connexion…")
            ws.close()
            break

        if user_input.lower() == "register":
            message = {
                "action": "Register",
                "API_KEY": "VDCAjPZ8jhDmXfsSufW2oZyU8SFZi48dRhA8zyKUjSRU3T1aBZ7E8FFIjdEM2X1d",
                "identifier": "tom_user",
                "password": "password",
                "game_prompt": "A werewolf game but every messages should rime"
            }

        elif user_input.lower() == "connect":
            message = {
                "action": "Connect",
                "identifier": "tom_user",
                "password": "password",
            }

        elif user_input.lower() == "createentity":
            if token:
                message = {
                    "action": "CreateEntity",
                    "name": "tom_entity",
                    "prompt": "a fisherman named tom_entity",
                    "checksum": "TomChecksum",
                    "role": "Werewolf",
                    "token": token,
                }
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")
                continue

        elif user_input.lower() == "newmessage":
            if token:
                message = {
                    "action":    "NewMessage",
                    "sender":    "TomChecksum",
                    "receivers": ["TomChecksum"],
                    "message":   "hello tom",
                    "token":     token,
                }
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")
                continue

        elif user_input.lower() == "makedecision":
            if token:
                message = {
                    "action": "MakeDecision",
                    "token": token,
                    "message": '{"phase": "Discussion"}',
                    "checksum": "TomChecksum",
                }
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")
                continue

        elif user_input.lower() == "night":
            if token:
                message = {
                    "action": "MakeDecision",
                    "token": token,
                    "message": '{"phase": "Night"}',
                    "checksum": "TomChecksum",
                }
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")
                continue

        else:
            print("Commande inconnue. Essayez 'Register', 'Connect', 'CreateEntity', 'NewMessage', 'MakeDecision', 'night', 'Disconnect', 'Remove' ou 'exit'.")
            continue

        action = user_input.lower()
        ws.send(json.dumps(message))

def main():
    ws = websocket.WebSocketApp(
        "ws://localhost:3000/ws",
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close
    )

    def handle_sigint(sig, frame):
        print("\nSIGINT reçu, fermeture…")
        ws.close()
        sys.exit(0)

    signal.signal(signal.SIGINT, handle_sigint)

    thread = threading.Thread(target=listen_to_stdin, args=(ws,), daemon=True)
    thread.start()

    ws.run_forever(ping_interval=20, ping_timeout=10)

if __name__ == "__main__":
    main()
