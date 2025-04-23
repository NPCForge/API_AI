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
PASSWORD = "Password"

WS_TOKEN = ""
HTTP_TOKEN = ""

try:
    with open(token_file, "r") as f:
        data = json.load(f)
        WS_TOKEN = data.get("ws", "")
        HTTP_TOKEN = data.get("http", "")

        if WS_TOKEN:
            print(f"✅ WS Token récupéré : {WS_TOKEN}")
        else:
            print("❌ Aucun token WS trouvé.")

        if HTTP_TOKEN:
            print(f"✅ HTTP Token récupéré : {HTTP_TOKEN}")
        else:
            print("❌ Aucun token HTTP trouvé.")

except FileNotFoundError:
    print("❌ Fichier token.json introuvable.")
except json.JSONDecodeError:
    print("❌ Erreur de lecture JSON.")



async def websocket():
    try:
        async with websockets.connect(uri) as websocket:
            message = {
                "action": "Disconnect",
                "token": WS_TOKEN,
            }

            await websocket.send(json.dumps(message))
            print(f"\n✅ WS - Message envoyé : {json.dumps(message)}")

            response = await websocket.recv()
            print(f"✅ WS - Réponse : {response}")

            response_data = json.loads(response)

            status = response_data.get("status")

            return status != "error"
    except Exception as e:
        print(f"❌ WS - Erreur : {e}")
        return False

def http():
    try:
        url = "http://localhost:3000/Disconnect"
        payload = json.dumps({
        })
        headers = {'Content-Type': 'application/json', "Authorization": HTTP_TOKEN}

        response = requests.post(url, headers=headers, data=payload)

        print(f"\n✅ HTTP - Code : {response.status_code}")
        print(f"✅ HTTP - Réponse : {response.text}")

        return response.status_code == 200
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