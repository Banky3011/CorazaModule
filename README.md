# CorazaModule

## Requirements

Before using this project, please ensure that you meet the following requirements:

1. **Go Installation (V 1.22 +):**
    You must have Go installed on your system. If you haven't installed it yet, you can download and install it from the [official Go website](https://golang.org/).
    
2. **System Path Configuration:**
    Make sure that the directory containing the Go binary is added to your system's PATH environment variable.

    Here to add Go path:
    ```
    sudo nano ~/.bashrc
    ```
    Then save:
    ```
    source ~/.bashrc
    ```

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

3. Install & Build the module: After install, restart the ubuntu then you can try to build
    ```
    ./setup.py <operation>
    ```

## Integrate with OWASP Core Ruleset
Core Ruleset can be installed by importing each required file in the following order:

1. Install coraza.conf, coreruleset/crs-setup.conf.example, coreruleset/rules/*.conf:
    ```
    wget https://raw.githubusercontent.com/corazawaf/coraza/v3/dev/coraza.conf-recommended -O coraza.conf
    git clone https://github.com/coreruleset/coreruleset
    ```

2. integrate the Core Ruleset:
    ```go
    waf, err := coraza.NewWAF(coraza.NewWAFConfig().
		WithDirectivesFromFile("coreruleset/crs-setup.conf").
		WithDirectivesFromFile("coreruleset/modsecurity.conf").
		WithDirectivesFromFile("coreruleset/rules/*.conf"))
    ```


## Usage
To use module in your project, follow these steps:

1. Import the module into your Python project:
    ```python
    from corazamodule import gomodule
    ```

2. Call the module:
    ```python
    gomodule.CorazaWAF()
    ```
3. Pass the necessary parameters to the module such as:
    ```python
        CorazaRequest = gomodule.Request(
                RemoteAddr = request.ip,
                Path = request.path,
                Port = request.port,
                Query = request.query_string,
                HTTPVersion = request.version,
                Method = request.method,
                Headers = headers_str,
                HeaderHost = HeaderHostPy,
                HeaderUserAgent = HeaderUserAgentPy,
                HeaderContentType = HeaderContentTypePy,
                Body = request.Body,
            )

            result = gomodule.CorazaWAF(CorazaRequest)
    ```

4. Run the project
