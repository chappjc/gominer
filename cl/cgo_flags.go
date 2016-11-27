package cl

/*
#cgo CFLAGS: -I /opt/intel/opencl-sdk/include/CL
#cgo windows CFLAGS: -IC:/appsdk/include
#cgo !darwin LDFLAGS: -L/opt/intel/opencl-sdk/lib64 -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL
#cgo windows LDFLAGS: -LC:/appsdk/lib/x86_64
*/
import "C"
