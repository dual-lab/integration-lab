"""
This module contains utils function used inside the repo build process
"""

load("@rules_jvm_external//:defs.bzl", "artifact")

def wrapped_artifact(group_id, artifact_id, repo_name = None):
  """
  Prefix the maven dep with the maven workspace name

  Args:
    group_id:     The maven group id
    articat_id:   The maven artifact id
    repo_name:    The repo name used in maven_install rule
  Return:
    The maven workspace for the specified jar deps
  """
  return artifact(group_id, artifact_id, repo_name)
