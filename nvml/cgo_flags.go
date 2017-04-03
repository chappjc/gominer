package nvml

/*
#cgo !windows LDFLAGS: -L/opt/cuda/lib64 -lnvidia-ml
#cgo windows LDFLAGS: -L../nvidia/NVSMI/ -lnvml
*/
import "C"
