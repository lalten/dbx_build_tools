from __future__ import annotations

from typing import Dict, List

###
### General configuration.
###
BUILD_INPUT = "BUILD.in"
DEFAULT_BUILD = "BUILD"
ALT_BUILD = "BUILD.bazel"

###
### Python-related configuration.
###
PY_EXTENSION_RULE_TYPES = ()

PY_RULE_TYPES = (
    "dbx_py_library",
    "dbx_py_binary",
    "dbx_py_compiled_binary",
    "dbx_py_pytest_test",
    "dbx_py_compiled_pytest_test",
)

PY_LIBRARY_RULE_TYPES = ("dbx_py_library", "py_library")

WELL_KNOWN_PIP_DIRS = ("pip", "thirdparty")

WELL_KNOWN_PY_EXTENSION_DIRS = ()

PY_LOAD_STATEMENT = """
load('@dbx_build_tools//build_tools/py:py.bzl', 'dbx_py_library', 'dbx_py_binary', 'dbx_py_pytest_test')
"""

EXTERNAL_PIP_MODULE_TARGETS: Dict[str, List[str]] = {}

PIP_DEFAULT_EXCLUDES = ["test", "tests", "testing", "SelfTest", "Test", "Tests"]
PIP_GEN_RULE_TYPES = ("dbx_py_pypi_piplib", "dbx_py_local_piplib")
PIP_RULE_TYPES = PIP_GEN_RULE_TYPES + ("dbx_py_piplib_alias",)
PIP_LOAD_STATEMENT = "load('@dbx_build_tools//build_tools/py:py.bzl', %s)" % (
    ", ".join([repr(t) for t in PIP_RULE_TYPES])
)

###
### Go-related configuration.
###
GO_RULE_TYPES = ("dbx_go_binary", "dbx_go_library", "dbx_go_test")
GO_TEST_RULE = "dbx_go_test"

WHITELISTED_GO_SRCS_PATHS: List[str] = []
