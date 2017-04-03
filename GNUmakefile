CC ?= gcc5 -fPIC -L/opt/cuda/lib64
CXX ?= g++5 -fPIC -L/opt/cuda/lib64
NVCC ?= nvcc -gencode arch=compute_20,code=sm_21 -gencode arch=compute_30,code=sm_30 -gencode arch=compute_35,code=sm_35 -gencode arch=compute_50,code=sm_50 -gencode arch=compute_52,code=compute_52 -gencode arch=compute_61,code=sm_61 -D_FORCE_INLINES -Xcompiler -fPIC -L/opt/cuda/lib64 -std=c++11
AR ?= ar
# -o is gnu only so this needs to be smarter; it does work because on darwin it
#  fails which is also not windows.
ARCH:=$(shell uname -o)

.DEFAULT_GOAL := build

ifeq ($(ARCH),Msys)
nvidia:
endif

# Windows needs additional setup and since cgo does not support spaces in
# in include and library paths we copy it to the correct location.
#
# Windows build assumes that CUDA V7.0 is installed in its default location.
#
# Windows gominer requires nvml.dll and decred.dll to reside in the same
# directory as gominer.exe.
ifeq ($(ARCH),Msys)
obj: nvidia
	mkdir nvidia
	cp -r /c/Program\ Files/NVIDIA\ GPU\ Computing\ Toolkit/* nvidia
	cp -r /c/Program\ Files/NVIDIA\ Corporation/NVSMI nvidia
else
obj:
endif
	mkdir obj

ifeq ($(ARCH),Msys)
obj/decred.dll: obj sph/blake.c decred.cu
	$(NVCC) --shared --optimize=3 --compiler-options=-GS-,-MD -I. -Isph decred.cu sph/blake.c -o obj/decred.dll
else
obj/decred.a: obj sph/blake.c decred.cu
	$(NVCC) -L/opt/cuda/lib64 --lib --optimize=3 -I. decred.cu sph/blake.c -o obj/decred.a
endif

ifeq ($(ARCH),Msys)
build: obj/decred.dll
else
build: obj/decred.a
endif
	go build -tags 'cuda'

ifeq ($(ARCH),Msys)
install: obj/decred.dll
else
install: obj/decred.a
endif
	go install -tags 'cuda'

clean:
	rm -rf obj
	go clean
ifeq ($(ARCH),Msys)
	rm -rf nvidia
endif
