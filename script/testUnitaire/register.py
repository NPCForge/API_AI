import asyncio
import json
import websockets
import nest_asyncio

nest_asyncio.apply()

uri = "ws://localhost:3000/ws"
token_file = "token.json"

async def register(websocket):
    message = {
        "action": "Register",
        "API_KEY": "VDCAjPZ8jhDmXfsSufW2oZyU8SFZi48dRhA8zyKUjSRU3T1aBZ7E8FFIjdEM2X1d",
        "identifier": "User_01_test",
        "password": "Password"
    }

    await websocket.send(json.dumps(message))
    print(f"\nMessage envoyé: {json.dumps(message)}")

    try:
        response = await websocket.recv()
        print(f"\nRéponse du serveur: {response}")
        response_data = json.loads(response)
        token = response_data.get("token")

        if token:
            with open(token_file, "w") as f:
                json.dump({"token": token}, f)
            print("✅ Token sauvegardé.")
        else:
            print("❌ Le serveur n'a pas renvoyé de token.")
    except json.JSONDecodeError:
        print("❌ Réponse du serveur invalide.")

async def run():
    try:
        async with websockets.connect(uri) as websocket:
            await register(websocket)
    except Exception as e:
        print(f"❌ Erreur de connexion : {e}")

asyncio.run(run())
