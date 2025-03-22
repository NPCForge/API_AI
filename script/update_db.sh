#!/bin/bash

# Configuration
DB_NAME="api_db"
DB_USER="API"
DB_HOST="localhost"   # Modifie si nécessaire
DB_PORT="5432"        # Modifie si nécessaire
MIGRATION_FILE="../migrations/migration.sql" # Fichier SQL à exécuter

# Vérifier si le fichier de migration existe
if [[ ! -f "$MIGRATION_FILE" ]]; then
    echo "Erreur : Le fichier $MIGRATION_FILE n'existe pas."
    exit 1
fi

# Exécuter la migration
echo "Mise à jour de la base de données $DB_NAME avec $MIGRATION_FILE..."
psql -U "$DB_USER" -d "$DB_NAME" -h "$DB_HOST" -p "$DB_PORT" -f "$MIGRATION_FILE"

# Vérifier si la migration s'est bien déroulée
if [[ $? -eq 0 ]]; then
    echo "Migration appliquée avec succès !"
else
    echo "Erreur lors de l'application de la migration."
    exit 1
fi
