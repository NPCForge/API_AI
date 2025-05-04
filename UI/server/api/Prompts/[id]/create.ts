import fs from 'fs';
import path from 'path';
import { readBody } from 'h3';  // Utilisation de h3 pour lire le corps de la requête

interface Prompt {
    message: string;
}

export default defineEventHandler(async (event): Promise<Prompt | { error: string }> => {
    // Résoudre le chemin absolu vers le dossier 'prompts'
    const promptsDirectory = "../prompts";

    // Récupérer le nom du fichier à partir du corps de la requête
    const body = await readBody(event);
    const { name } = body;

    // Vérifier si le nom du fichier est fourni
    if (!name) {
        return { error: 'Le nom du fichier est requis.' };
    }

    const filePath = path.join(promptsDirectory, name);  // Résolution du chemin complet du fichier
    console.log('Chemin du fichier:', filePath);

    try {
        // Vérifier si le fichier existe
        if (!fs.existsSync(filePath)) {
            // Si le fichier n'existe pas, le créer (le fichier est vide au départ)
            await fs.promises.appendFile(filePath, "Hello world!");  // Création d'un fichier vide
            console.log('Fichier créé avec succès:', filePath);

            // Retourner un message de succès
            return { message: "Fichier créé avec succès." };
        } else {
            return { error: 'Le fichier existe déjà.' };  // Si le fichier existe déjà, ne pas le recréer
        }
    } catch (error) {
        return { error: `Erreur lors de la création du fichier: ${error instanceof Error ? error.message : 'Unknown error'}` };
    }
});
