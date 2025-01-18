import { getUserByName } from '../DataBase';
import bcrypt from 'bcryptjs';

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const { username, password } = body;

    if (!username || !password) {
        throw createError({ statusCode: 400, message: 'Nom d\'utilisateur et mot de passe requis.' });
    }

    try {
        const user = await getUserByName(username);
        console.log("user login", user)

        if (!user) {
            throw createError({ statusCode: 401, message: 'Identifiants invalides.' });
        }

        const isPasswordValid = await bcrypt.compare(password, user.password);

        if (isPasswordValid) {
            console.log('Passwords match! User authenticated.');
            return { success: true, message: 'Connexion réussie !' };
        } else {
            console.log('Passwords do not match! Authentication failed.');
            return { success: false, message: 'La Connexion a échoué, mauvais mot de passe !' };
        }
    } catch (error) {
        throw createError({ statusCode: 500, message: error.message });
    }
});
