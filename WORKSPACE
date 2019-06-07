load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)

BAZEL_BUILD_TOOLS_SHA = "f27d1753c8b3210d9e87cdc9c45bc2739ae2c2db"

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "cdaac537b56375f658179ee2f27813cac19542443f4722b6730d84e4125355e6",
    strip_prefix = "buildtools-" + BAZEL_BUILD_TOOLS_SHA,
    url = "https://github.com/bazelbuild/buildtools/archive/" + BAZEL_BUILD_TOOLS_SHA + ".zip",
)

ATLASSIAN_BAZEL_TOOLS_SHA = "f04c7c0ec6851ea378f90c74c855bcb87d9793b5"

http_archive(
    name = "com_github_atlassian_bazel_tools",
    sha256 = "281aba7b549daaf418302cb0a51e8358d6557a03185019d6b521d9a12744c957",
    strip_prefix = "bazel-tools-" + ATLASSIAN_BAZEL_TOOLS_SHA,
    urls = ["https://github.com/atlassian/bazel-tools/archive/" + ATLASSIAN_BAZEL_TOOLS_SHA + ".zip"],
)

git_repository(
    name = "io_bazel_rules_docker",
    commit = "5d35fee42a513ac83dc23c32ba1f0b029fd75e0f",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    shallow_since = "1559569106 -0700",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

load("@com_github_atlassian_bazel_tools//goimports:deps.bzl", "goimports_dependencies")

goimports_dependencies()

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

go_repository(
    name = "com_github_stripe_skycfg",
    commit = "80fe5c6480773463ac32f81613e210c5db45c3f4",
    importpath = "github.com/stripe/skycfg",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "7b8349ac747c6a24702b762d2c4fd9266cf4f1d6",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "net_starlark_go",
    commit = "854ab5fc397e48ac689f94c4084e4b1842ef3c46",
    importpath = "go.starlark.net",
)
