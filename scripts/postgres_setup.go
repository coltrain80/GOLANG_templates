// postgres_setup.go
/*
PostgreSQL Setup Script in Go
------------------------------

This script installs PostgreSQL on Ubuntu, configures it with initial settings,
and creates a generic database.

Usage:
    Run the script with sudo or as a root user to install and configure PostgreSQL.

Example:
    sudo go run postgres_setup.go
*/

package main

import (
    "fmt"
    "os"
    "os/exec"
)

// Helper function to run a system command and handle errors
func runCommand(command string) {
    fmt.Printf("Running: %s\n", command)
    cmd := exec.Command("bash", "-c", command)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("Command failed: %s\n", err)
        os.Exit(1)
    }
}

// Main function to install and configure PostgreSQL
func main() {
    if os.Geteuid() != 0 {
        fmt.Println("This script must be run as root. Please use sudo.")
        os.Exit(1)
    }

    // Step 1: Update the package list
    runCommand("apt-get update")

    // Step 2: Install PostgreSQL
    runCommand("apt-get install -y postgresql postgresql-contrib")

    // Step 3: Start PostgreSQL service
    runCommand("systemctl start postgresql")

    // Step 4: Enable PostgreSQL service to start on boot
    runCommand("systemctl enable postgresql")

    // Step 5: Configure PostgreSQL
    configurePostgres()
}

// Function to configure PostgreSQL with initial settings and create a generic database
func configurePostgres() {
    postgresPassword := "password123" // Replace with a secure password
    databaseName := "generic_db"
    userName := "generic_user"
    userPassword := "user_password" // Replace with a secure password

    commands := []string{
        fmt.Sprintf(`sudo -u postgres psql -c "ALTER USER postgres PASSWORD '%s';"`, postgresPassword),
        fmt.Sprintf(`sudo -u postgres psql -c "CREATE ROLE %s WITH LOGIN PASSWORD '%s';"`, userName, userPassword),
        fmt.Sprintf(`sudo -u postgres psql -c "CREATE DATABASE %s OWNER %s;"`, databaseName, userName),
        fmt.Sprintf(`sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE %s TO %s;"`, databaseName, userName),
    }

    for _, command := range commands {
        runCommand(command)
    }

    fmt.Println("PostgreSQL installed and configured successfully.")
    fmt.Printf("Database '%s' created with user '%s'.\n", databaseName, userName)
}
