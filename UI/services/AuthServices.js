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

            console.log('Utilisateur créé avec succès :', data.value);
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
            const saltRounds = 10;
            // Hash du mot de passe côté client
            const hashedPassword = await bcrypt.hash(password, saltRounds);

            // Requête de connexion
            const { data, error } = await useFetch('/api/login', {
                method: 'POST',
                body: { username, password: hashedPassword }
            });

            console.log(data)

            if (error.value) {
                console.error('Erreur lors de la connexion :', error.value.message);
                return false;
            }

            console.log('Connexion réussie :', data.value);
            return true;

        } catch (err) {
            console.error('Erreur de connexion :', err.message);
            return false;
        }
    } else {
        console.error('Nom d\'utilisateur et mot de passe requis.');
        return false;
    }
};

export {
    isSettup,
    firstRegister,
    login
}