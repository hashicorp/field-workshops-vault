#!/bin/bash -l

set -e

grep -q 'vault.*write.*lob_a/workshop/database/config/wsmysqldatabase.*plugin_name=mysql-database-plugin.*connection_url="{{username}}:{{password}}@tcp(localhost:3306)/".*allowed_roles="workshop-app","workshop-app-long".*username="hashicorp".*password="Password123"' /root/.bash_history || fail-message "You haven't configured the wsmysqldatabase connection yet."

grep -q "mysql.*-u.*hashicorp.*-pPassword123" /root/.bash_history || fail-message "You haven't logged into the MySQL server with the 'hashicorp' username yet."

grep -q "vault.*write.*-force.*lob_a/workshop/database/rotate-root/wsmysqldatabase" /root/.bash_history || fail-message "You haven't rotated the root credentials for the wsmysqldatabase connection yet."

root_login_attempts=$(grep "mysql.*-u.*hashicorp.*-pPassword123" /root/.bash_history | wc -l)
if [ $root_login_attempts -lt 2 ]; then
  echo "You have not tried to login to the MySQL server a second time yet (after rotating the root credentials)."
  exit 1
fi

grep -q "vault.*write.*lob_a/workshop/database/roles/workshop-app-long.*db_name=wsmysqldatabase.*creation_statements=\"CREATE USER '{{name}}'@'%' IDENTIFIED BY '{{password}}';GRANT ALL ON my_app.* TO '{{name}}'@'%';\".*default_ttl=\"1h\".*max_ttl=\"24h\"" /root/.bash_history || fail-message "You haven't configured the workshop-app-long role yet."

grep -q "vault.*write.*lob_a/workshop/database/roles/workshop-app.*db_name=wsmysqldatabase.*creation_statements=\"CREATE USER '{{name}}'@'%' IDENTIFIED BY '{{password}}';GRANT ALL ON my_app.* TO '{{name}}'@'%';\".*default_ttl=\"3m\".*max_ttl=\"6m\"" /root/.bash_history || fail-message "You haven't configured the workshop-app role yet."

exit 0
