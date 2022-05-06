"""
Contains all the maven deps for the project
"""

load("@rules_jvm_external//:specs.bzl", "maven")

_QUARKUS_VERSION = "2.7.3.Final"
_QUARKUS_GROUP_ID = "io.quarkus"

_JUNIT_VERSION = "5.8.2"
_JUNIT_JUPITER_GROUP_ID = "org.junit.jupiter"
_JUNIT_PLATFORM_GROUP_ID = "org.junit.platform"

def quarkus_core_deps(version = _QUARKUS_VERSION):
    """
    Genere the querkus artifact deps
    Args:
      version:    quarkus version, optional
    Return:
      List of maven artifact
    """

    deps = []

    # Resteasy reactive
    deps.append(maven.artifact(
        version,
        group = _QUARKUS_GROUP_ID,
        artifact = "quarkus-resteasy-reactive",
    ))

    # quarkus arc
    deps.append(maven.artifact(
        version,
        group = _QUARKUS_GROUP_ID,
        artifact = "quarkus-arc",
    ))

    return deps

def quarkus_test_deps(version = _QUARKUS_VERSION):
    """
    Collect quarkus test deps
    Args:
      version:      quarkus version, optional
    Return:
      List of maven artifact
    """

    deps = []

    # quarkus junit 5
    deps.append(maven.artifact(
        version,
        group = _QUARKUS_GROUP_ID,
        artifact = "quarkus-junit5",
    ))

    # rest assured
    deps.append(maven.artifact(
        group = "io.rest-assured",
        artifact = "rest-assured",
        version = "4.5.1",
    ))

    # add junit deps
    _JUNIT_JUPITER_ARTIFACT_ID_LIST = [
        "junit-jupiter-api",
        "junit-jupiter-engine",
        "junit-jupiter-params",
    ]
    deps.extend([maven.artifact(
        group = _JUNIT_JUPITER_GROUP_ID,
        artifact = t,
        version = _JUNIT_VERSION,
    ) for t in _JUNIT_JUPITER_ARTIFACT_ID_LIST])

    # add junit platform deps
    _JUNIT_PLATFORM_ARTIFACT_ID_LIST = [
        "junit-platform-commons",
        "junit-platform-console",
        "junit-platform-engine",
        "junit-platform-launcher",
        "junit-platform-suite-api",
    ]
    deps.extend([maven.artifact(
        group = _JUNIT_PLATFORM_GROUP_ID,
        artifact = t,
        version = _JUNIT_VERSION,
    ) for t in _JUNIT_PLATFORM_ARTIFACT_ID_LIST])

    # extra deps
    _JUNIT_EXTRA_DEPENDENCIES = [
        ("org.apiguardian", "apiguardian-api", "1.0.0"),
        ("org.opentest4j", "opentest4j", "1.1.1"),
    ]
    deps.extend(["%s:%s:%s" % t for t in _JUNIT_EXTRA_DEPENDENCIES])

    return deps
