#!/usr/bin/python3

import sys
import subprocess

__help__ = """
Usage: ./setup.py <operation>

Operations:
- install: Install dependencies
- build: Build module from Go to Python
"""

def check_dependencies_installed():
    try:
        subprocess.check_call(['gopy', 'version'], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
        return True
    except subprocess.CalledProcessError:
        return False

def install_dependencies():
    if not check_dependencies_installed():
        subprocess.check_call(['go', 'get', 'github.com/go-python/gopy@latest'])
        subprocess.check_call(['go', 'install', 'golang.org/x/tools/cmd/goimports@latest'])
    else:
        print("No dependencies to install")

def run_build_script():
    subprocess.check_call(['gopy', 'build', '-output=./corazamodule', '-name=corazamodule', './gomodule'])

def main():
    if len(sys.argv) != 2:
        print(__help__)
        sys.exit(1)

    operation = sys.argv[1]

    if operation == "install":
        install_dependencies()
    elif operation == "build":
        run_build_script()
    else:
        print("Error")
        sys.exit(1)

if __name__ == '__main__':
    main()
