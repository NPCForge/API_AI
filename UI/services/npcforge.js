// Exemple de code pour se connecter et stocker le token
export const connect = async (identifier, password) => {
    try {
        const response = await fetch('http://0.0.0.0:3000/Connect', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                identifier: identifier,
                password: password,
            }),
        })

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`)
        }

        const data = await response.json()

        // Si le token est présent, le stocker dans localStorage
        if (data.token) {
            localStorage.setItem('token', data.token)  // Stocker le token dans localStorage
            return true
        } else {
            return false
        }
    } catch (error) {
        console.error('Erreur de connexion:', error)
        return false
    }
}

export const disconnect = async () => {
    try {
        if (!data.token)
            return true
        const response = await fetch('http://0.0.0.0:3000/Disconnect', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                token: localStorage.getItem("token")
            }),
        })

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`)
        }

        localStorage.removeItem('token')
        return true

    } catch (error) {
        console.error('Erreur de connexion:', error)
        return false
    }
}

export const getPrompts = async () => {
    try {
        const response = await fetch('/api/Prompts/getAll');
        const data = await response.json();
        console.log('Prompts récupérés:', data);
        return data;
    } catch (error) {
        console.error('Erreur lors de la récupération des prompts:', error);
    }
};

export const editPrompt = async (name, content) => {
    try {
        const response = await fetch(`/api/Prompts/${name}/edit`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json', // Assurez-vous que le body est en JSON
            },
            body: JSON.stringify({
                content: content,
            }),
        });

        // Vérifie si la réponse est OK (code HTTP 2xx)
        if (!response.ok) {
            throw new Error(`Erreur du serveur: ${response.statusText}`);
        }

        // Récupérer les données de la réponse
        const data = await response.json();

        // Retourner les données
        return data;
    } catch (error) {
        console.error('Erreur lors de la mise à jour du prompt:', error);
        throw error;
    }
};

export const removePrompt = async (name) => {
    try {
        const response = await fetch(`/api/Prompts/${name}/remove`, { method: 'POST' });

        // Vérifie si la réponse est OK (code HTTP 2xx)
        if (!response.ok) {
            throw new Error(`Erreur du serveur: ${response.statusText}`);
        }

        // Retourner les données
        return true
    } catch (error) {
        console.error('Erreur lors de la mise à jour du prompt:', error);
        throw error;
    }
};

export const createPrompt = async (name) => {
    try {
        // Changer la méthode HTTP en POST pour envoyer des données
        const response = await fetch(`/api/Prompts/${name}/create`, {
            method: 'POST',  // Utilisation de POST pour envoyer les données
            headers: {
                'Content-Type': 'application/json',  // Le serveur attend du JSON
            },
            body: JSON.stringify({ name })  // Envoi du nom du fichier dans le corps
        });

        // Vérification de la réponse du serveur
        if (!response.ok) {
            throw new Error(`Erreur du serveur: ${response.statusText}`);
        }

        // Retourner un message de succès
        return true;
    } catch (error) {
        console.error('Erreur lors de la mise à jour du prompt:', error);
        throw error;
    }
};


