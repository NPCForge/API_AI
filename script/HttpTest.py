import requests
import json

# Base URL de l'API
BASE_URL = "http://localhost:3000/"

# Variable globale pour stocker le token reçu
token = None

def send_request(endpoint, payload=None):
    """Envoie une requête HTTP POST avec les données fournies et ajoute le token dans les en-têtes."""
    global token
    headers = {}

    # Ajoute le token dans l'en-tête Authorization si disponible
    if token:
        headers["Authorization"] = token

    try:
        response = requests.post(f"{BASE_URL}/{endpoint}", json=payload, headers=headers)
        response.raise_for_status()  # Vérifie si la réponse est une erreur HTTP
        print(response.json())
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f"Erreur lors de la requête : {e}")
        return None

def main():
    global token
    print("Entrez une commande (Register, Connection, TakeDecision, Disconnect) ou 'exit' pour quitter.")

    while True:
        user_input = input("Commande : ").strip().lower()

        if user_input == "exit":
            print("Arrêt du programme.")
            break

        # Préparation des données en fonction de la commande
        if user_input == "register":
            payload = {
                "API_KEY": "VDCAjPZ8jhDmXfsSufW2oZyU8SFZi48dRhA8zyKUjSRU3T1aBZ7E8FFIjdEM2X1d",
                "checksum": "azerty",
                "name": "tom",
                "prompt": "juste le boss"
            }
            response = send_request("Register", payload)
            if response and "token" in response:
                token = response["token"]
                print(f"Token mis à jour : {token}")
            else:
                print("Erreur : Échec de l'enregistrement.")

        elif user_input == "connection":
            payload = {
                "action": "Connection",
                "checksum": "azerty"
            }
            response = send_request("Connect", payload)
            if response and "token" in response:
                token = response["token"]
                print(f"Connexion réussie. Token : {token}")
            else:
                print("Erreur : Échec de la connexion.")

        elif user_input == "takedecision":
            if token:
                payload = {
                    "message": "Hello World"
                }
                response = send_request("MakeDecision", payload)
                if response:
                    print(f"Réponse : {response}")
                else:
                    print("Erreur : Échec de la prise de décision.")
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")

        elif user_input == "disconnect":
            if token:
                payload = {
                    "action": "Disconnect"
                }
                response = send_request("Disconnect", payload)
                if response:
                    print("Déconnexion réussie.")
                    token = None
                else:
                    print("Erreur : Échec de la déconnexion.")
            else:
                print("Erreur : Aucun token disponible. Veuillez vous connecter ou vous enregistrer d'abord.")

        else:
            print("Commande inconnue. Essayez 'Register', 'Connection', 'TakeDecision', 'Disconnect', ou 'exit'.")

if __name__ == "__main__":
    main()
