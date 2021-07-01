package main

import (
	shimrunh "github.com/JTS22/shim-runh/shim"
	containerdshim "github.com/containerd/containerd/runtime/v2/shim"
)

func main() {
	containerdshim.Run("io.containerd.runh.v1", shimrunh.New)
}
