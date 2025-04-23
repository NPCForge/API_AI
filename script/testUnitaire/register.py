import asyncio
import json
import websockets
import nest_asyncio
import requests
import sys

nest_asyncio.apply()

uri = "ws://localhost:3000/ws"
token_file = "token.json"

# Données communes
API_KEY = "VDCAjPZ8jhDmXfsSufW2oZyU8SFZi48dRhA8zyKUjSRU3T1aBZ7E8FFIjdEM2X1d"
PASSWORD = "Password"

# Identifiants séparés pour différencier les requêtes
WS_IDENTIFIER = "User_01_test_ws"
HTTP_IDENTIFIER = "User_01_test_http"

# Réinitialise le fichier token.json
with open(token_file, "w") as f:
    json.dump({}, f)

async def websocket():
    try:
        async with websockets.connect(uri) as websocket:
            message = {
                "action": "Register",
                "API_KEY": API_KEY,
                "identifier": WS_IDENTIFIER,
                "password": PASSWORD
            }

            await websocket.send(json.dumps(message))
            print(f"\n✅ WS - Message envoyé : {json.dumps(message)}")

            response = await websocket.recv()
            print(f"✅ WS - Réponse : {response}")
            response_data = json.loads(response)

            token = response_data.get("token")
            if token:
                # Lire les tokens existants
                try:
                    with open(token_file, "r") as f:
                        tokens = json.load(f)
                except (FileNotFoundError, json.JSONDecodeError):
                    tokens = {}

                tokens["ws"] = token

                with open(token_file, "w") as f:
                    json.dump(tokens, f)
                print("✅ WS - Token sauvegardé.")
                return True

    except Exception as e:
        print(f"❌ WS - Erreur : {e}")
        return False

def http():
    try:
        url = "http://localhost:3000/Register"
        payload = json.dumps({
            "API_KEY": API_KEY,
            "Identifier": HTTP_IDENTIFIER,
            "Password": PASSWORD
        })
        headers = {'Content-Type': 'application/json'}

        response = requests.post(url, headers=headers, data=payload)

        print(f"\n✅ HTTP - Code : {response.status_code}")
        print(f"✅ HTTP - Réponse : {response.text}")

        response_data = json.loads(response.text)

        token = response_data.get("token")

        if token:
            # Lire les tokens existants
            try:
                with open(token_file, "r") as f:
                    tokens = json.load(f)
            except (FileNotFoundError, json.JSONDecodeError):
                tokens = {}

            tokens["http"] = token

            with open(token_file, "w") as f:
                json.dump(tokens, f)
            print("✅ HTTP - Token sauvegardé.")
            return True
    except Exception as e:
        print(f"❌ HTTP - Erreur : {e}")
        return False

async def main():
    ws_success = await websocket()
    http_success = http()

    if not ws_success:
        print("\n❌ WS.")
    if not http_success:
        print("\n❌ HTTP.")

    return ws_success and http_success

if __name__ == "__main__":
    result = asyncio.run(main())
    if result:
        print("\n✅ Les deux connexions (WS + HTTP) ont réussi.")
        sys.exit(0)
    else:
        print("\n❌ Échec de l'une ou des deux connexions.")
        sys.exit(1)