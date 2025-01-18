import { verifyCredentials } from '../DataBase';

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const { username, password } = body;

    if (!username || !password) {
        throw createError({ statusCode: 400, message: 'Nom d\'utilisateur et mot de passe requis.' });
    }

    try {
        // Vérifie les identifiants (le mot de passe est déjà hashé côté client)
        const user = await verifyCredentials(username, password);
        console.log(user)

        if (!user) {
            throw createError({ statusCode: 401, message: 'Identifiants invalides.' });
        }

        return { success: true, message: 'Connexion réussie !' };
    } catch (error) {
        throw createError({ statusCode: 500, message: error.message });
    }
});
