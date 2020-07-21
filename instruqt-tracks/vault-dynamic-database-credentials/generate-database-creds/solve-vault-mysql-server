#!/bin/bash -l

#Enable bash history
HISTFILE=/root/.bash_history
set -o history

# Generate creds for the workshop-app-long role using the Vault API
curl --header "X-Vault-Token: root" "http://localhost:8200/v1/lob_a/workshop/database/creds/workshop-app-long" | jq

# Generate creds for the workshop-app-long role using the Vault CLI
vault read lob_a/workshop/database/creds/workshop-app-long

# Generate creds for the workshop-app role using the Vault CLI
{ read USER; read PASSWORD; } < <(vault read -format=json lob_a/workshop/database/creds/workshop-app | jq -r '.data | .username,.password')

# Login to the MySQL database with the mysql utility
mysql -u ${USER} -p${PASSWORD} -e "show databases"
if [ $? -ne 0 ]; then
    fail-message "mysql was unable to connect to the MySQL server."
fi

exit 0
