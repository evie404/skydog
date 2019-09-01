load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "313f2c7a23fecc33023563f082f381a32b9b7254f727a7dd2d6380ccc6dfe09b",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.3/rules_go-0.19.3.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.3/rules_go-0.19.3.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
    ],
)

git_repository(
    name = "com_google_protobuf",
    commit = "09745575a923640154bcf307fba8aedff47f240a",
    remote = "https://github.com/protocolbuffers/protobuf",
    shallow_since = "1558721209 -0700",
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

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

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
    commit = "24246af8ff5f56427b506403ba62aadbcf5413a0",
    importpath = "go.starlark.net",
)

go_repository(
    name = "com_github_zorkian_go_datadog_api",
    commit = "5f6a486f4345634ed789aa21bab35c924387df55",
    importpath = "github.com/zorkian/go-datadog-api",
)
