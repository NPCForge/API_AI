import fs from 'fs';
import path from 'path';
import { readBody } from 'h3';  // Assure-toi que 'h3' est bien installé, cela permet de lire le body de la requête

interface Prompt {
    fileName: string;
    content: string;
}

export default defineEventHandler(async (event): Promise<Prompt | { error: string }> => {
    // Résoudre le chemin absolu vers le dossier 'prompts'
    const promptsDirectory = "../prompts";

    console.log(await fs.promises.readdir(promptsDirectory))

    // Récupérer l'ID depuis l'URL
    const { id } = event.context.params;  // Récupère l'ID de l'URL (par exemple: api/Prompt/[id]/edit)
    console.log(id)
    // Lire le corps de la requête (nouveau contenu)
    const body = await readBody(event);  // Récupère le body envoyé dans la requête
    const newContent = body.content;  // Le nouveau contenu est passé avec la clé 'content'

    console.log(newContent)

    if (!newContent) {
        return { error: 'No content provided in the request body' };
    }

    const filePath = path.join(promptsDirectory, `${id}`); // Le fichier sera nommé avec l'ID (par exemple: '1.txt')

    try {
        // Vérifier si le fichier existe
        if (fs.existsSync(filePath)) {
            await fs.promises.writeFile(filePath, newContent, 'utf8');

            // Retourner un message de succès
            return {
                fileName: `${id}`,
                content: newContent
            };
        } else {
            return { error: 'File not found' };
        }
    } catch (error) {
        return { error: `Error writing file: ${error instanceof Error ? error.message : 'Unknown error'}` };
    }
});
