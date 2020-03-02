.PHONY: all

curr_time = $(shell date "+%Y%m%d_%H%M%S")
pprof_path = .pprof/${curr_time}

all: test bench

test:
	go test

bench:
	mkdir -p ${pprof_path}
	go test -bench=BenchmarkMergeSort -run=none -count=5 -benchmem -test.parallel=1 -cpuprofile=cpu_merge.out -blockprofile=block_merge.out -memprofile=mem_merge.out -outputdir=${pprof_path}
	go test -bench=BenchmarkNormalSort -run=none -count=5 -benchmem -test.parallel=1 -cpuprofile=cpu_normal.out -blockprofile=block_normal.out -memprofile=mem_normal.out -outputdir=${pprof_path}

