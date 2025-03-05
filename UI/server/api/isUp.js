import { ofetch } from 'ofetch';

export default defineEventHandler(async (event) => {
    try {
        const data = await ofetch("http://localhost:8000/Health", {
            method: "GET"
        });

        return { success: true, data, status_code: 200 };
    } catch (error) {
        console.error("Erreur lors de la requête à l'API externe :", error);

        return { success: false, error: error.message, status_code: 444 };
    }
});
