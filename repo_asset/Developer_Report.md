## How to get better in this project

### Security

* Current Strategy

Restrict Query Access
  ```
  // use regex to match string
  datePattern := `^\d{4}-\d{2}-\d{2}$`
  ```
Parameterized query (for insert data)
  ```
  query := fmt.Sprintf("INSERT INTO %s(lons, lats, counts) VALUES ($1, $2, $3)", tableName)
  ```

* Improve Strategy

SQL password management (Also can apply on API Keys)
  ```
  // Currently no security of password
  connStr := "user=USERNAME dbname=DBNAME password=PASSWORD sslmode=disable"

  // Potentially Method
  1. Set all password (or API keys or anything confidential) in one file, ex. PASSWORD.txt
  2. Define read-in function
  
  func readPassword(filepath string) string {
  	// read Password file
  	...
  	return password
  }

  3. PASSWORD.txt MUST NOT BE SEND INTO GIT REPO !!!
  ```

Restricted Query Frequency
  ```
  Potential Principle: Maximum number of N query for per user in M minute
  Violation : Ban from query for X minutes
  ```
Simple Example
  ```
  Potential Principle: Maximum number of 10 query for per user in 1 minute
  Violation : Ban from query for 10 minutes

  // run the Example
  go run restricted_query_frequency.go
  ```