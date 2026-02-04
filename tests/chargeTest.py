import requests
import time
import statistics
import json
from concurrent.futures import ThreadPoolExecutor, as_completed
from typing import Tuple
import os
from dotenv import load_dotenv
from websocket import create_connection

# Configuration
load_dotenv("../.env")
server = "localhost"
port = 3000
BASE_URL = f"http://{server}:{port}"
WS_URL = f"ws://{server}:{port}/ws"
API_KEY = os.getenv("API_KEY_REGISTER", "test_api_key")

MAX_USERS = 10
ENTITIES_PER_USER = 10

class LoadTestResults:
    def __init__(self, protocol: str):
        self.protocol = protocol
        self.results = []
    
    def add_result(self, num_users: int, latency: float, success: bool):
        self.results.append({
            'num_users': num_users,
            'latency': latency,
            'success': success
        })
    
    def print_summary(self):
        print("\n" + "="*80)
        print(f"RÉSUMÉ DES TESTS DE CHARGE - {self.protocol}")
        print("="*80)
        
        for num_users in range(1, MAX_USERS + 1):
            user_results = [r for r in self.results if r['num_users'] == num_users]
            if not user_results:
                continue
            
            latencies = [r['latency'] for r in user_results if r['success']]
            successes = sum(1 for r in user_results if r['success'])
            failures = len(user_results) - successes
            
            if latencies:
                avg_latency = statistics.mean(latencies)
                min_latency = min(latencies)
                max_latency = max(latencies)
                
                print(f"\n{num_users} utilisateur(s):")
                print(f"  ├─ Latence moyenne : {avg_latency:.3f}s")
                print(f"  ├─ Latence min     : {min_latency:.3f}s")
                print(f"  ├─ Latence max     : {max_latency:.3f}s")
                print(f"  ├─ Succès          : {successes}")
                print(f"  └─ Échecs          : {failures}")
            else:
                print(f"\n{num_users} utilisateur(s): Tous les tests ont échoué")
        
        print("\n" + "="*80)

# ==================== HTTP FUNCTIONS ====================

def register_user(user_id: str) -> Tuple[bool, str, float]:
    """Enregistre un utilisateur via HTTP et retourne (succès, token, temps)"""
    start_time = time.time()
    
    try:
        payload = {
            "API_KEY": API_KEY,
            "Identifier": user_id,
            "Password": "TestPassword123",
            "game_prompt": "game_prompt"
        }
        
        response = requests.post(
            f"{BASE_URL}/Register",
            json=payload,
            timeout=30
        )
        
        elapsed = time.time() - start_time
        
        if response.status_code == 200:
            data = response.json()
            token = data.get("token", "")
            return True, token, elapsed
        else:
            print(f"❌ Erreur d'enregistrement pour {user_id}: {response.status_code}")
            return False, "", elapsed
            
    except Exception as e:
        elapsed = time.time() - start_time
        print(f"❌ Exception lors de l'enregistrement de {user_id}: {e}")
        return False, "", elapsed

def create_entity(user_id: str, token: str, entity_num: int) -> Tuple[bool, float]:
    """Crée une entité via HTTP pour un utilisateur et retourne (succès, temps)"""
    start_time = time.time()
    
    try:
        payload = {
            "name": f"Entity_{user_id}_{entity_num}",
            "prompt": f"Prompt for entity {entity_num}",
            "role": "Villager" if entity_num % 2 == 0 else "Werewolf",
            "checksum": f"checksum_{entity_num}"
        }
        
        headers = {
            "Content-Type": "application/json",
            "Authorization": token
        }
        
        response = requests.post(
            f"{BASE_URL}/CreateEntity",
            json=payload,
            headers=headers,
            timeout=30
        )
        
        elapsed = time.time() - start_time
        
        if response.status_code == 200:
            return True, elapsed
        else:
            print(f"❌ Erreur création entité pour {user_id}: {response.status_code}")
            return False, elapsed
            
    except Exception as e:
        elapsed = time.time() - start_time
        print(f"❌ Exception création entité pour {user_id}: {e}")
        return False, elapsed

def process_user(user_num: int) -> Tuple[bool, float]:
    """Traite un utilisateur complet: enregistrement + création de toutes ses entités (HTTP)"""
    user_id = f"LoadTestUser_HTTP_{user_num:03d}"
    
    # Enregistrement
    success, token, reg_time = register_user(user_id)
    if not success:
        return False, reg_time
    
    total_time = reg_time
    
    # Création des entités
    for entity_num in range(1, ENTITIES_PER_USER + 1):
        success, entity_time = create_entity(user_id, token, entity_num)
        total_time += entity_time
        if not success:
            return False, total_time
    
    return True, total_time

# ==================== WEBSOCKET FUNCTIONS ====================

def register_user_ws(user_id: str) -> Tuple[bool, str, float]:
    """Enregistre un utilisateur via WebSocket et retourne (succès, token, temps)"""
    start_time = time.time()
    
    try:
        ws = create_connection(WS_URL, timeout=30)
        
        message = {
            "action": "Register",
            "API_KEY": API_KEY,
            "identifier": user_id,
            "password": "TestPassword123",
            "game_prompt": "game_prompt"
        }
        
        ws.send(json.dumps(message))
        response = ws.recv()
        ws.close()
        
        elapsed = time.time() - start_time
        
        data = json.loads(response)
        token = data.get("token", "")
        
        if token:
            return True, token, elapsed
        else:
            print(f"❌ Erreur d'enregistrement WS pour {user_id}: pas de token")
            return False, "", elapsed
            
    except Exception as e:
        elapsed = time.time() - start_time
        print(f"❌ Exception lors de l'enregistrement WS de {user_id}: {e}")
        return False, "", elapsed

def create_entity_ws(user_id: str, token: str, entity_num: int) -> Tuple[bool, float]:
    """Crée une entité via WebSocket pour un utilisateur et retourne (succès, temps)"""
    start_time = time.time()
    
    try:
        ws = create_connection(WS_URL, timeout=30)
        
        message = {
            "action": "CreateEntity",
            "name": f"Entity_{user_id}_{entity_num}",
            "prompt": f"Prompt for entity {entity_num}",
            "role": "Villager" if entity_num % 2 == 0 else "Werewolf",
            "checksum": f"checksum_{entity_num}",
            "token": token
        }
        
        ws.send(json.dumps(message))
        response = ws.recv()
        ws.close()
        
        elapsed = time.time() - start_time
        
        data = json.loads(response)
        entity_id = data.get("id", "")
        
        if entity_id:
            return True, elapsed
        else:
            print(f"❌ Erreur création entité WS pour {user_id}: pas d'id")
            return False, elapsed
            
    except Exception as e:
        elapsed = time.time() - start_time
        print(f"❌ Exception création entité WS pour {user_id}: {e}")
        return False, elapsed

def process_user_ws(user_num: int) -> Tuple[bool, float]:
    """Traite un utilisateur complet: enregistrement + création de toutes ses entités (WebSocket)"""
    user_id = f"LoadTestUser_WS_{user_num:03d}"
    
    # Enregistrement
    success, token, reg_time = register_user_ws(user_id)
    if not success:
        return False, reg_time
    
    total_time = reg_time
    
    # Création des entités
    for entity_num in range(1, ENTITIES_PER_USER + 1):
        success, entity_time = create_entity_ws(user_id, token, entity_num)
        total_time += entity_time
        if not success:
            return False, total_time
    
    return True, total_time

# ==================== TEST FUNCTIONS ====================

def test_with_n_users(n_users: int, results: LoadTestResults, protocol: str = "HTTP"):
    """Teste avec n utilisateurs en parallèle"""
    print(f"\n🔄 Test {protocol} avec {n_users} utilisateur(s)...")
    
    start_time = time.time()
    process_func = process_user_ws if protocol == "WebSocket" else process_user
    
    with ThreadPoolExecutor(max_workers=n_users) as executor:
        futures = {
            executor.submit(process_func, i): i 
            for i in range(1, n_users + 1)
        }
        
        for future in as_completed(futures):
            user_num = futures[future]
            try:
                success, user_time = future.result()
                results.add_result(n_users, user_time, success)
                status = "✅" if success else "❌"
                print(f"  {status} Utilisateur {user_num}: {user_time:.3f}s")
            except Exception as e:
                print(f"  ❌ Utilisateur {user_num} a échoué: {e}")
                results.add_result(n_users, 0, False)
    
    total_elapsed = time.time() - start_time
    print(f"⏱️  Temps total pour {n_users} utilisateur(s): {total_elapsed:.3f}s")

def test_charge_http():
    """Test de charge HTTP: de 1 à 10 utilisateurs"""
    print("="*80)
    print("DÉBUT DU TEST DE CHARGE HTTP")
    print(f"Configuration: {MAX_USERS} utilisateurs, {ENTITIES_PER_USER} entités par utilisateur")
    print("="*80)
    
    results = LoadTestResults("HTTP")
    
    # Tester progressivement de 1 à MAX_USERS utilisateurs
    for n_users in range(1, MAX_USERS + 1):
        test_with_n_users(n_users, results, protocol="HTTP")
        # Petite pause entre chaque niveau pour éviter de surcharger le serveur
        time.sleep(1)
    
    # Afficher le résumé
    results.print_summary()
    
    # Sauvegarder les résultats dans un fichier JSON
    with open('load_test_results_http.json', 'w') as f:
        json.dump(results.results, f, indent=2)
    
    print("\n📊 Résultats HTTP sauvegardés dans: load_test_results_http.json")
    return results

def test_charge_websocket():
    """Test de charge WebSocket: de 1 à 10 utilisateurs"""
    print("\n" + "="*80)
    print("DÉBUT DU TEST DE CHARGE WEBSOCKET")
    print(f"Configuration: {MAX_USERS} utilisateurs, {ENTITIES_PER_USER} entités par utilisateur")
    print("="*80)
    
    results = LoadTestResults("WebSocket")
    
    # Tester progressivement de 1 à MAX_USERS utilisateurs
    for n_users in range(1, MAX_USERS + 1):
        test_with_n_users(n_users, results, protocol="WebSocket")
        # Petite pause entre chaque niveau pour éviter de surcharger le serveur
        time.sleep(1)
    
    # Afficher le résumé
    results.print_summary()
    
    # Sauvegarder les résultats dans un fichier JSON
    with open('load_test_results_websocket.json', 'w') as f:
        json.dump(results.results, f, indent=2)
    
    print("\n📊 Résultats WebSocket sauvegardés dans: load_test_results_websocket.json")
    return results

def compare_results(http_results: LoadTestResults, ws_results: LoadTestResults):
    """Compare les résultats HTTP et WebSocket"""
    print("\n" + "="*80)
    print("COMPARAISON HTTP vs WebSocket")
    print("="*80)
    
    for n_users in range(1, MAX_USERS + 1):
        http_latencies = [r['latency'] for r in http_results.results 
                         if r['num_users'] == n_users and r['success']]
        ws_latencies = [r['latency'] for r in ws_results.results 
                       if r['num_users'] == n_users and r['success']]
        
        if http_latencies and ws_latencies:
            http_avg = statistics.mean(http_latencies)
            ws_avg = statistics.mean(ws_latencies)
            diff = ((http_avg - ws_avg) / http_avg) * 100
            
            faster = "WebSocket" if ws_avg < http_avg else "HTTP"
            symbol = "🚀" if ws_avg < http_avg else "🐢"
            
            print(f"\n{n_users} utilisateur(s):")
            print(f"  ├─ HTTP       : {http_avg:.3f}s")
            print(f"  ├─ WebSocket  : {ws_avg:.3f}s")
            print(f"  └─ {symbol} {faster} est {abs(diff):.1f}% plus rapide")
    
    print("\n" + "="*80)

def test_charge_complete():
    """Lance les tests HTTP et WebSocket puis compare les résultats"""
    # Test HTTP
    http_results = test_charge_http()
    
    # Pause entre les deux protocoles
    print("\n⏸️  Pause de 3 secondes avant le test WebSocket...")
    time.sleep(3)
    
    # Test WebSocket
    ws_results = test_charge_websocket()
    
    # Comparaison
    compare_results(http_results, ws_results)

if __name__ == "__main__":
    test_charge_complete()
