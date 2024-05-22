#BIN
The **bin** directory will contain the compiled application binaries, ready for deployment to production server.
#CMD/API
The **cmd/api** directory will the contain the application-specific code for the API. This will include the code running the server, reading, and writing HTTP requests, and managing authentication.
#INTERNAL
The **internal** directory will contain various ancillary packages used by the API. it will containing the code for interaction with the database, doing data validation, sending emails and so on. Basically, any code which isn't application-specific can be potentially be reused will live here. Go code under **cmd/api** will import the packages in the **internal** directy(but never the other way around.)
#MIGRATIONS
The **migrations** directory will contain the SQL migration file for the database.
#REMOTE
The **remote** directory will contain the configuration files and setup scripts for the production server.
#GO.MOD
The **go.mod** file will declare the project dependecies, version, and module path.
#MAKEFILE
The **makefile** will contain recipes for automating common administrative tasks - like auditing Go code, building binaries, and executing database migrations.
