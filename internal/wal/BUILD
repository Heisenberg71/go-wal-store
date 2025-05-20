load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "wal_lib",
    srcs = ["wal.go"],
    importpath = "simplewal/internal/wal",
    visibility = ["//:__subpackages__"],
    deps = [
        "//proto:wal_go_proto",
        "@org_golang_google_protobuf//proto",
    ],
)
