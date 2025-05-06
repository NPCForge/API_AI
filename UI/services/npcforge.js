// services/npcforge.js

export const connect = async (identifier, password, simulate = false) => {
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
            const errorData = await response.json();
            throw new Error(`HTTP error! Status: ${response.status} - Message: ${errorData.message || 'Unknown error'}`);
        }

        const data = await response.json();

        if (data.token) {
            if (!simulate) {
                localStorage.setItem('token', data.token);
            }
            return {
                Status: "Success",
                Data: data
            };
        } else {
            return {
                Status: "Failed",
                Data: data
            };
        }
    } catch (error) {
        console.error('Erreur de connexion:', error.message);
        return {
            Status: "Failed",
            Data: null,
            Message: error.message
        };
    }
}

export const register = async (identifier, password, API_TOKEN) => {
    console.log(identifier, password, API_TOKEN)
    try {
        const response = await fetch('http://0.0.0.0:3000/Register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                identifier: identifier,
                password: password,
                API_KEY: API_TOKEN
            }),
        })

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(`HTTP error! Status: ${response.status} - Message: ${errorData.message || 'Unknown error'}`);
        }

        const data = await response.json();

        if (data.token) {
            return {
                Status: "Success",
                Data: data
            };
        } else {
            return {
                Status: "Failed",
                Data: data
            };
        }
    } catch (error) {
        console.error('Erreur de connexion:', error.message);
        return {
            Status: "Failed",
            Data: null,
            Message: error.message
        };
    }
}

export const removeUser = async (identifier=null, API_TOKEN) => {
    try {
        const response = await fetch('http://0.0.0.0:3000/RemoveUser', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': API_TOKEN
            },
            body: JSON.stringify({
                username: identifier,
            }),
        })

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(`HTTP error! Status: ${response.status} - Message: ${errorData.message || 'Unknown error'}`);
        }

        const data = await response.json();

        if (data.token) {
            return {
                Status: "Success",
                Data: data
            };
        } else {
            return {
                Status: "Failed",
                Data: data
            };
        }
    } catch (error) {
        console.error('Erreur de connexion:', error.message);
        return {
            Status: "Failed",
            Data: null,
            Message: error.message
        };
    }
}

export const status = async (token = null) => {
    try {
        if (!token) {
            token = localStorage.getItem("token");
        }
        if (!token) {
            return {
                Status: "Failed",
                Message: "Token is missing"
            };
        }

        const response = await fetch('http://0.0.0.0:3000/Status', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${token}`
            }
        });

        if (!response.ok) {
            const errorData = await response.json();
            localStorage.removeItem('token');
            throw new Error(`HTTP error! Status: ${response.status} - Message: ${errorData.message || 'Unknown error'}`);
        }

        return {
            Status: "Success",
            Message: "Status request successful"
        };
    } catch (error) {
        console.error('Erreur de connexion:', error.message);
        return {
            Status: "Failed",
            Message: error.message || "An error occurred during the status check"
        };
    }
};

export const disconnect = async (tokenSimulation = null) => {
    try {
        let token = null;

        // Utiliser le token soit de localStorage, soit fourni par le paramètre
        if (!tokenSimulation) {
            token = localStorage.getItem("token");
            if (!token) {
                return { Status: "Failed", Message: "Token is missing" };
            }
        } else {
            token = tokenSimulation;
        }

        const response = await fetch('http://0.0.0.0:3000/Disconnect', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${token}`
            }
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(`HTTP error! Status: ${response.status} - Message: ${errorData.message || 'Unknown error'}`);
        }
        if (!tokenSimulation)
            localStorage.removeItem('token');

        return {
            Status: "Success",
            Message: "Disconnected successfully"
        };

    } catch (error) {
        console.error('Error during disconnection:', error);
        return {
            Status: "Failed",
            Message: error.message || "An error occurred during disconnection"
        };
    }
};


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


