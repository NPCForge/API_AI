import fs from 'fs';
import path from 'path';

export default defineEventHandler(async (event) => {
    // Résoudre le chemin absolu vers le dossier 'prompts'
    const promptsDirectory = "../prompts"

    try {
        const files = await fs.promises.readdir(promptsDirectory);

        if (files.length > 0) {
            let arr = [];
            for (let val of files) { // Corriger 'for of' avec une déclaration 'let' pour 'val'
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
        return { error: `Error reading file: ${error.message}` };
    }
});
