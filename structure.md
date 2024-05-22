# Project Structure

## BIN
The **bin** directory will contain the compiled application binaries, ready for deployment to the production server.

## CMD/API
The **cmd/api** directory will contain the application-specific code for the API. This will include the code running the server, reading, and writing HTTP requests, and managing authentication.

## INTERNAL
The **internal** directory will contain various ancillary packages used by the API. It will contain the code for interaction with the database, doing data validation, sending emails, and so on. Basically, any code which isn't application-specific but can potentially be reused will live here. Go code under **cmd/api** will import the packages in the **internal** directory (but never the other way around).

## MIGRATIONS
The **migrations** directory will contain the SQL migration files for the database.

## REMOTE
The **remote** directory will contain the configuration files and setup scripts for the production server.

## GO.MOD
The **go.mod** file will declare the project dependencies, versions, and module path.

## MAKEFILE
The **makefile** will contain recipes for automating common administrative tasks - like auditing Go code, building binaries, and executing database migrations.

