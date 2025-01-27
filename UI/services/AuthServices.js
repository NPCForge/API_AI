import { useFetch } from 'nuxt/app';
import bcrypt from 'bcryptjs';

const isSettup = () => {
    const { data, error } = useFetch('/api/get-all-user');

    if (error.value) {
        console.error('Erreur lors de la récupération des utilisateurs :', error.value.message);
        return false;
    }
    // console.log("herree", data.value.length)
    return computed(() => data.value && data.value.length > 0);
}

const firstRegister = async (username, password) => {
    if (username !== "" && password !== "") {
        try {
            const saltRounds = 10;
            const hashedPassword = await bcrypt.hash(password, saltRounds);  // Hash du mot de passe

            const { data, error } = await useFetch('/api/add-user', {
                method: 'POST',
                body: {
                    name: username,
                    password: hashedPassword,
                    access: 7 // root
                }
            });

            if (error.value) {
                console.error('Erreur lors de la création de l\'utilisateur :', error.value.message);
                return false;
            }

            // console.log('Utilisateur créé avec succès :', data.value);
            return true;

        } catch (err) {
            console.error('Erreur :', err.message);
            return false;
        }
    } else {
        console.error('Le nom d\'utilisateur et le mot de passe sont requis.');
        return false;
    }
};

const login = async (username, password) => {
    if (username !== "" && password !== "") {
        try {
            // Requête de connexion avec $fetch
            const response = await $fetch('/api/login', {
                method: 'POST',
                body: { username, password }
            });

            // Vérification de la réponse
            if (response.success && response.token) {
                // Enregistrer le token dans le localStorage
                localStorage.setItem('authToken', response.token);

                // console.log('Connexion réussie, token enregistré dans le localStorage.');
                return true;
            } else {
                console.error('Connexion échouée :', response.message || 'Erreur inconnue.');
                return false;
            }
        } catch (err) {
            console.error('Erreur lors de la connexion :', err.message);
            return false;
        }
    } else {
        console.error('Nom d\'utilisateur et mot de passe requis.');
        return false;
    }
};

const getAuthToken = () => {
    return localStorage.getItem('authToken');
};

const logout = () => {
    localStorage.removeItem('authToken');
    // console.log('Déconnexion réussie, token supprimé du localStorage.');
};


export {
    isSettup,
    firstRegister,
    login,
    getAuthToken,
    logout
}