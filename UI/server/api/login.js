import { getUserByName } from '../DataBase';
import bcrypt from 'bcryptjs';
import jwt from 'jsonwebtoken';
import { setToken } from '../TokenStore';
const config = useRuntimeConfig().public

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const { username, password } = body;

    const manage_token = (user) => {
        const token = jwt.sign(
            {
                id: user.id,
                name: user.name,
                access: user.access
            },
            config.JWT_SECRET_KEY,
            { expiresIn: '1h' }
        );
        console.log("token", token)
        return token
    }

    if (!username || !password) {
        throw createError({ statusCode: 400, message: 'Nom d\'utilisateur et mot de passe requis.' });
    }

    try {
        const user = await getUserByName(username);

        if (!user) {
            throw createError({ statusCode: 401, message: 'Identifiants invalides.' });
        }

        const isPasswordValid = await bcrypt.compare(password, user.password);

        if (isPasswordValid) {
            console.log('Passwords match! User authenticated.');

            const token = manage_token(user);

            await setToken(token, user.id);

            return { success: true, message: 'Connexion réussie !', token: token };
        } else {
            console.log('Passwords do not match! Authentication failed.');
            return { success: false, message: 'La Connexion a échoué, mauvais mot de passe !' };
        }
    } catch (error) {
        throw createError({ statusCode: 500, message: error.message });
    }
});
