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

export const getPrompts = async () => {
    try {
        const response = await fetch('/api/getPrompts');
        const data = await response.json();
        console.log('Prompts récupérés:', data);
        return data;
    } catch (error) {
        console.error('Erreur lors de la récupération des prompts:', error);
    }
};