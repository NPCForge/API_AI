const adresse = "http://0.0.0.0:3000/"

export const connect = (identifier, password) => {
    fetch(adresse + "Connect", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            identifier: identifier,
            password: password
        }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`); // Vérifie si la réponse est correcte
        }
        return response.json();
    })
    .then(data => {
        console.log(data);
        if (data.status == 200 && data.token != "") {
            return true
        }
    })
    .catch(error => {
        console.error('Erreur:', error);
        return false
    });
    return false
}
