//go:build test
// +build test

package config

// Admin constants
const SUPER_ADMIN_EMAIL_ADDRESS = "frontend@lastdisco.com"
const SUPER_ADMIN_PHONE_NUMBER = "72233445588"
const SUPER_ADMIN_FIRST_NAME = "issa"
const SUPER_ADMIN_LAST_NAME = "paintin'"
const SUPER_ADMIN_PASSWORD = "lastdisco12345"

// API-Related stuff
var PROD_CRUD_ROLES = []string{"ADMIN", "MODERATOR"}
var USER_CRUD_ROLES = []string{"ADMIN", "MODERATOR"}

// Program-Wide defines
const TESTING = true
const PRODUCTION = false

// Authentication
const JWT_SECRET = "lastdisco-jwt-secret"

// PostgresDB constants
const PG_USER = "lastdisco_admin_db"
const PG_NAME = "postgres"
const PG_PASS = "Eequo2quAiBok9su"

const PG_USR_NAME_MAXLEN = 75
