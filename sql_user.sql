drop user central_client;
create user central_client with encrypted password 'central_client';
grant all privileges on database centraldb to central_client;
grant all privileges on all tables in schema public to central_client;
