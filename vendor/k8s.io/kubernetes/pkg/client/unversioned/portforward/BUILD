package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
    "go_test",
    "cgo_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "portforward.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/kubelet/server/portforward:go_default_library",
        "//pkg/util/httpstream:go_default_library",
        "//pkg/util/runtime:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["portforward_test.go"],
    library = "go_default_library",
    tags = ["automanaged"],
    deps = [
        "//pkg/client/restclient:go_default_library",
        "//pkg/client/unversioned/remotecommand:go_default_library",
        "//pkg/kubelet/server/portforward:go_default_library",
        "//pkg/types:go_default_library",
        "//pkg/util/httpstream:go_default_library",
    ],
)
