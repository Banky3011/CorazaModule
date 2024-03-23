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
    ```python
    from corazamodule import gomodule
    print(1+1)
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

