# CorazaModule

## Requirements

Before using this project, please ensure that you meet the following requirements:

1. **Go Installation:**
    - You must have Go installed on your system. If you haven't installed it yet, you can download and install it from the [official Go website](https://golang.org/).
    
2. **System Path Configuration:**
    - Make sure that the directory containing the Go binary is added to your system's PATH environment variable.


## Installation
Follow these steps to install CorazaModule:

1. Clone this repository to your local machine:
    ```
    git clone https://github.com/Banky3011/CorazaModule.git
    ```

2. Navigate to the project directory:
    ```
    cd CorazaModule
    ```

3. Install & Build the module:
    ```
    ./setup.py <operation>
    ```

4. Once the dependencies are installed and built, you can start using CorazaModule in your projects by importing as a module.

## Integrate with OWASP Core Ruleset
    ```go
    func MyWaf() gin.HandlerFunc {
	waf, err := coraza.NewWAF(coraza.NewWAFConfig().
		WithDirectivesFromFile("coraza.conf").
		WithDirectivesFromFile("coreruleset/crs-setup.conf.example").
		WithDirectivesFromFile("coreruleset/rules/*.conf").
		WithDirectives(
			`SecRule REMOTE_ADDR "!@ipMatch 127.0.0.1" "id:1,phase:1,deny,status:403"`,
		))
	if err != nil {
		log.Fatalf("Error creating WAF: %v", err)
	}

	waf, err := coraza.NewWAF(coraza.NewWAFConfig().
		WithDirectives(
		`SecRule REMOTE_ADDR "!@ipMatch 127.0.0.1" "id:1,phase:1,deny,status:403"`,
		))
		if err != nil {
			log.Fatalf("Error creating WAF:", err)
		}
	
	return func(c *gin.Context) {
		tx := waf.NewTransaction()
		defer func() {
			tx.ProcessLogging()
			log.Println("Transaction closed successfully")
			tx.Close()
		}()
	
		tx.ProcessConnection("127.0.0.1", 8080, "127.0.0.1", 12345)
	
		if it := tx.ProcessRequestHeaders(); it != nil {
			log.Printf("Transaction was interrupted with status %d\n", it.Status)
			c.AbortWithStatus(it.Status)
			return
		} else {
			log.Printf("Request Allowed")
		}
	
		c.Next()
	}
}
    ```

## Usage
To use CorazaModule in your project, follow these steps:

1. Import the module into your Python project:
    ```python
    from corazamodule import gomodule
    ```

2. Call the CorazaModule :
    ```python
    gomodule.Runserver()
    ```

