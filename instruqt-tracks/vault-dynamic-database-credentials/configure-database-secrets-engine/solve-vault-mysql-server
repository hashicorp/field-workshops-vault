#!/bin/bash -l

#Enable bash history
HISTFILE=/root/.bash_history
set -o history

# Configure a connection for the database secrets engine
vault write lob_a/workshop/database/config/wsmysqldatabase \
  plugin_name=mysql-database-plugin \
  connection_url="{{username}}:{{password}}@tcp(localhost:3306)/" \
  allowed_roles="workshop-app","workshop-app-long" \
  username="hashicorp" \
  password="Password123"

# Login to MySQL server with root creds
# Simulating success
echo -e "mysql -u hashicorp -pPassword123" >> /root/.bash_history

# Rotate the root credentials
vault write -force lob_a/workshop/database/rotate-root/wsmysqldatabase

# Login to MySQL server with root creds
# Simulating failure
echo -e "mysql -u hashicorp -pPassword123" >> /root/.bash_history

# Configure a role for the database secrets engine
vault write lob_a/workshop/database/roles/workshop-app-long \
  db_name=wsmysqldatabase \
  creation_statements="CREATE USER '{{name}}'@'%' IDENTIFIED BY '{{password}}';GRANT ALL ON my_app.* TO '{{name}}'@'%';" \
  default_ttl="1h" \
  max_ttl="24h"

# Configure a role for the database secrets engine
vault write lob_a/workshop/database/roles/workshop-app \
  db_name=wsmysqldatabase \
  creation_statements="CREATE USER '{{name}}'@'%' IDENTIFIED BY '{{password}}';GRANT ALL ON my_app.* TO '{{name}}'@'%';" \
  default_ttl="3m" \
  max_ttl="6m"

exit 0
