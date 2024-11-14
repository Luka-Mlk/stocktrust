# Stocktrust
Homework project for **Software Design and Architecture**

## Dependencies

This project requires the following:

- **Golang 1.23+**
- **Docker**
- **Any PostgreSQL GUI** of your choice (e.g., pgAdmin, DBeaver)

## Config

### On Linux

To configure your environment, run the following commands in your terminal:

```bash
export DATABASE_USER="your_user_choice"
export DATABASE_PASSWORD="your_password_choice"
export DATABASE_HOST="127.0.0.1"
export DATABASE_PORT="5432"
export DATABASE_NAME="stocktrust"
export NUM_THREADS="desired_number_of_threads"
```

### On Windows

1. **Hardcode the Database Configurations**:
   - You will need to manually set the **database user**, **password**, **host**, **port**, and **name** in your code.

2. **Thread Configuration**:
   - Remove environment variable error handling.
   - Explicitly define the number of threads you want to use, e.g.:
     ```go
     threadsInt := 20
     ```

## How to Start

### Step-by-Step Instructions:

1. **Start the Server**:
   - Use the following command in your terminal to start the server:
     ```bash
     docker-compose up
     ```
   - Alternatively, you can use the Docker Desktop app to start the server.

2. **Connect to PostgreSQL**:
   - Connect the server to the PostgreSQL GUI of your choice (e.g., pgAdmin or DBeaver).

3. **Create Tables**:
   - Create the necessary tables (`company` and `history_record`) by using the SQL files found in the `sql-exports` folder.

4. **Run the Application**:
   - Open your terminal, navigate to the project directory, and run the following command:
     ```bash
     go run main.go
     ```

After completing these steps, your database should be configured and the server should be up and running.




##Execution of program

In the following image is the time needed for our program to scrape and save all needed information from the Macedonian stock exchange.
![Alt text]time_executed.JPG
