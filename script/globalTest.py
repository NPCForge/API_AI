import sys
import time
import threading
import asyncio
import websockets
import json

uri = "ws://localhost:3000/ws"

stop_event = threading.Event()  # Déclarer globalement

async def run():
    async with websockets.connect(uri) as websocket:
        # Construire un message JSON
        message = {
            "action": "Register",
            "API_KEY": "VDCAjPZ8jhDmXfsSufW2oZyU8SFZi48dRhA8zyKUjSRU3T1aBZ7E8FFIjdEM2X1d",
            "checksum": "azerty",
            "name": "tom",
            "prompt": "I am a pirate. I always finish my sentences by 'gomu gomu'",
        }
        json_message = json.dumps(message)  # Convertir en JSON

        # Envoyer le message
        await websocket.send(json_message)
        print(f"\nMessage envoyé: {json_message}")

        # Recevoir la réponse
        response = await websocket.recv()
        print(f"\nRéponse du serveur: {response}")

def spinner():
    """Affiche une animation de chargement en arrière-plan."""
    spinner_symbols = ['-', '\\', '|', '/']
    i = 0
    while not stop_event.is_set():
        sys.stdout.write(f"\rChargement {spinner_symbols[i % len(spinner_symbols)]}")
        sys.stdout.flush()
        time.sleep(0.1)
        i += 1
    sys.stdout.write("\rChargement terminé !\n")

def main():
    # Lancer le thread du spinner
    spinner_thread = threading.Thread(target=spinner)
    spinner_thread.start()

    # Exécuter la tâche asyncio dans l'event loop
    asyncio.run(run())

    # Arrêter l'animation
    stop_event.set()

    # Attendre la fin du thread
    spinner_thread.join()

    print("Toutes les opérations sont terminées !")

if __name__ == "__main__":
    main()
