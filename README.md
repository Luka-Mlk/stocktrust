# stocktrust
Homework project for Software design and Architecture

## Dependencies 

This project requires:
        golang 1.23+
        docker
        any postgre gui of your choice

## Config
 On linux: 
 ``` bash
export DATABASE_USER=any user of your choice
export DATABASE_PASSWORD=any password of your choice
export DATABASE_HOST=127.0.0.1
export DATABAES_PORT=5432
export DATABASE_NAME=stocktrust
export NUM_THREADS=desired amount of threads
 ```

 On Windos: 
 
 - You will need to hardcode in the databse user, password, host, port and name.
 - You will need to remove error handling for environment handling and explicitly say how many threads you want to use ex. threadsInt := 20

 ## How to Start
 - You will need to start the server through docker compose up or through the docker.desktop app
 - You will need to connect the server to the postgre gui of your choice.
 - You will need to creat the company and history_record tables with the files in sql-exports
 - You will need to open the terminal and type go run main.go then click enter

 Your database should be configured now.
