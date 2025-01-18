import pkg from 'pg';

const { Pool } = pkg;
const config = useRuntimeConfig().public

// Configuration de la connexion PostgreSQL
const pool = new Pool({
    user: config.POSTGRES_USER || 'votre_utilisateur',
    host: config.POSTGRES_HOST || 'localhost',
    database: config.POSTGRES_DB || 'votre_base',
    password: config.POSTGRES_PASSWORD || 'votre_mot_de_passe',
    port: Number(config.POSTGRES_PORT) || 5432,
});

const query = async (text, params = []) => {
    try {
        const result = await pool.query(text, params);
        return result.rows;
    } catch (error) {
        console.error('Erreur de requête SQL :', error.message);
        throw new Error(error.message);
    }
};

const getUsersByAccess = async (access) => {
    const queryText = `
      SELECT * FROM user_ui
      WHERE access = $1
      ORDER BY created DESC;
    `;
    const result = await query(queryText, [access]);
    return result;
};

const verifyCredentials = async (name, password) => {
    const queryText = `
      SELECT * FROM user_ui
      WHERE name = $1 AND password = $2;
    `;
    const result = await query(queryText, [name, password]);
    return result[0] || null;
};

// ✅ : implemented (./api/get-all-user.js)
const getAllUsers = async () => {
    const queryText = `
      SELECT * FROM user_ui
      ORDER BY created DESC;
    `;
    const result = await query(queryText);
    return result;
};

const deleteUserById = async (id) => {
    const queryText = `
      DELETE FROM user_ui
      WHERE id = $1
      RETURNING *;
    `;
    const result = await query(queryText, [id]);
    return result[0];
};

const updateAccess = async (id, newAccess) => {
    const queryText = `
      UPDATE user_ui
      SET access = $1
      WHERE id = $2
      RETURNING *;
    `;
    const params = [newAccess, id];
    const result = await query(queryText, params);
    return result[0];
};

const updatePassword = async (id, newPassword) => {
    const queryText = `
      UPDATE user_ui
      SET password = $1
      WHERE id = $2
      RETURNING *;
    `;
    const params = [newPassword, id];
    const result = await query(queryText, params);
    return result[0];
};

const getUserByName = async (name) => {
    const queryText = `
      SELECT * FROM user_ui
      WHERE name = $1;
    `;
    const result = await query(queryText, [name]);
    return result[0] || null;
};

const getUserById = async (id) => {
    const queryText = `
      SELECT * FROM user_ui
      WHERE id = $1;
    `;
    const result = await query(queryText, [id]);
    return result[0] || null;
};

// ✅ : implemented (./api/add-user.js)
const addUser = async (name, password, access = 0) => {
    const queryText = `
      INSERT INTO user_ui (name, password, access)
      VALUES ($1, $2, $3)
      RETURNING *;
    `;
    const params = [name, password, access];
    const result = await query(queryText, params);
    return result[0];
};

// Export des fonctions
export {
    addUser,
    getUserById,
    getUserByName,
    updatePassword,
    updateAccess,
    deleteUserById,
    getAllUsers,
    verifyCredentials,
    getUsersByAccess,
};