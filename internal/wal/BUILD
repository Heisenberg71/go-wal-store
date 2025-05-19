load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "wal_lib",
    srcs = ["wal.go"],
    importpath = "simplewal/internal/wal",
    visibility = ["//:__subpackages__"],
)

go_test(
    name = "wal_test",
    srcs = ["wal_test.go"],
    deps = [":wal_lib"],
)
