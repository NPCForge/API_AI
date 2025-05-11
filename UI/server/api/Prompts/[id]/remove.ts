import fs from 'fs';
import path from 'path';
import { readBody } from 'h3';

interface Prompt {
    message: string;
}

export default defineEventHandler(async (event): Promise<Prompt | { error: string }> => {
    // Résoudre le chemin absolu vers le dossier 'prompts'
    const promptsDirectory = "../prompts";

    // Récupérer l'ID depuis l'URL
    const { id } = event.context.params;

    const filePath = path.join(promptsDirectory, `${id}`);

    try {
        // Vérifier si le fichier existe
        if (fs.existsSync(filePath)) {
            await fs.promises.unlink(filePath);

            return {
                message: "success"
            };
        } else {
            return { error: 'File not found' };
        }
    } catch (error) {
        return { error: `Error writing file: ${error instanceof Error ? error.message : 'Unknown error'}` };
    }
});
