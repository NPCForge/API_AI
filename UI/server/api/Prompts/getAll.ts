import fs from 'fs';
import path from 'path';

interface Prompt {
    fileName: string;
    content: string;
}

export default defineEventHandler(async (event): Promise<Prompt[] | { error: string }> => {
    // Résoudre le chemin absolu vers le dossier 'prompts'
    const promptsDirectory = "../prompts";

    try {
        const files = await fs.promises.readdir(promptsDirectory);

        if (files.length > 0) {
            const arr: Prompt[] = []; // Déclarer le tableau avec le type approprié
            for (let val of files) {
                const filePath = path.join(promptsDirectory, val); // Résolution du chemin du fichier
                const fileContent = await fs.promises.readFile(filePath, 'utf8'); // Lire le contenu du fichier
                arr.push({
                    fileName: val,  // Utiliser 'val' ici pour le nom du fichier courant
                    content: fileContent
                });
            }
            return arr; // Retourner tous les fichiers et leur contenu
        } else {
            return { error: 'No files found in prompts directory' };
        }
    } catch (error) {
        return { error: `Error reading file: ${error instanceof Error ? error.message : 'Unknown error'}` };
    }
});
