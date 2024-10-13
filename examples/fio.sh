#!/bin/sh

MODE="$1"

NUMJOBS=8
IODEPTH=16

if [ "$MODE" = "" ]; then 
    MODE = "small_randown_rw"
fi

if [ "$MODE" = "small_read" ]; then 
    fio --name=test-small-read --direct=1 --ioengine=libaio --group_reporting --rw=read --iodepth=$(IODEPTH) --bs=4k --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "small_write" ]; then 
    fio --name=test-small-write  --direct=1 --ioengine=libaio --group_reporting --rw=write --iodepth=$(IODEPTH) --bs=4k --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "small_randread" ]; then 
    fio --name=test-small-randread --direct=1 --ioengine=libaio --group_reporting --rw=randread --iodepth=$(IODEPTH) --bs=4k --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "small_randwrite" ]; then 
    fio --name=test-small-randwrite --direct=1 --ioengine=libaio --group_reporting --rw=randwrite --iodepth=$(IODEPTH) --bs=4k --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "small_randown_rw" ]; then 
    fio --name=test-small-rw --direct=1 --ioengine=libaio --group_reporting --rw=randrw --rwmixwrite=30 --iodepth=$(IODEPTH) --bs=4k --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "large_read" ]; then 
    fio --name=test-large-read --direct=1 --ioengine=libaio --group_reporting --rw=read --iodepth=$(IODEPTH) --bs=1m --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "large_write" ]; then 
    fio --name=test-large-write --direct=1 --ioengine=libaio --group_reporting --rw=write --iodepth=$(IODEPTH) --bs=1m --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "large_randread" ]; then 
    fio --name=test-large-randread --direct=1 --ioengine=libaio --group_reporting --rw=randread --iodepth=$(IODEPTH) --bs=1m --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "large_randwrite" ]; then 
    fio --name=test-large-randwrite --direct=1 --ioengine=libaio --group_reporting --rw=randwrite --iodepth=$(IODEPTH) --bs=1m --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi

if [ "$MODE" = "large_randown_rw" ]; then 
    fio --name=test-large-rand-rw --direct=1 --ioengine=libaio --group_reporting --rw=randrw --rwmixwrite=30 --iodepth=$(IODEPTH) --bs=1m --filename=/dev/sda --size=1G --numjobs=$(NUMJOBS) --runtime=120
fi