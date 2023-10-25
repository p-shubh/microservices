#!/bin/bash

# check the postgres Sql is already installed
if ! command psql --version &>/dev/null; then

    # install the postgres sql and psql tool
    echo "Installing th postgresSQL and psql tool..."

    # add installation command based on the operating system on the server
    # for the ubuntu
    sudo apt-get update
    sudo apt-get install -y postgresql

    # check if the installation was successfull or not for installing the postgres
    if ! command psql --version &>/dev/null; then
        echo "Failed to install PostgreSQL and psql. Please install them manually."
        exit 1
    fi
fi

# Database connection parameters
DB_NAME="new_database"
DB_USER="new_user"
DB_PASSWORD="user_password"

# Check if the database already exists
if psql -lqt | cut -d \| -f 1 | grep -qw $DB_NAME; then
    echo "Database $DB_NAME already exists."
else
    # Create the database
    psql -U postgres -c "CREATE DATABASE $DB_NAME"
    echo "Database $DB_NAME created successfully."

    # Create the user with a password
    psql -U postgres -d $DB_NAME -c "CREATE USER $DB_USER WITH ENCRYPTED PASSWORD '$DB_PASSWORD'"
    echo "User $DB_USER with password created successfully."

    # Grant privileges (if needed)
    # psql -U postgres -d $DB_NAME -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER"
fi
