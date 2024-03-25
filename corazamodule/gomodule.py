
# python wrapper for package example/hello/gomodule within overall package corazamodule
# This is what you import to use the package.
# File is generated by gopy. Do not edit.
# gopy build -output=./corazamodule -name=corazamodule ./gomodule

# the following is required to enable dlopen to open the _go.so file
import os,sys,inspect,collections
try:
	import collections.abc as _collections_abc
except ImportError:
	_collections_abc = collections

cwd = os.getcwd()
currentdir = os.path.dirname(os.path.abspath(inspect.getfile(inspect.currentframe())))
os.chdir(currentdir)
from . import _corazamodule
from . import go

os.chdir(cwd)

# to use this code in your end-user python file, import it as follows:
# from corazamodule import gomodule
# and then refer to everything using gomodule. prefix
# packages imported by this package listed below:




# ---- Types ---


#---- Enums from Go (collections of consts with same type) ---


#---- Constants from Go: Python can only ask that you please don't change these! ---


# ---- Global Variables: can only use functions to access ---


# ---- Interfaces ---


# ---- Structs ---


# ---- Slices ---


# ---- Maps ---


# ---- Constructors ---


# ---- Functions ---
def RunServer(goRun=False):
	"""RunServer() """
	_corazamodule.gomodule_RunServer(goRun)


