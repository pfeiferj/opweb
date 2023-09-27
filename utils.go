package main

import (
	"capnproto.org/go/capnp/v3"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var BasePath string = GetBasePath()

func check(e error) {
	if e != nil {
		log.Println(e)
		panic(e)
	}
}

// exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MapSlice[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func GetBasePath() string {
	exists, err := Exists("/data/media/0/realdata")
	check(err)
	if exists {
		return "/data/media/0/realdata/"
	} else {
		return "media/"
	}
}

func Segments(route string) []string {
	files, err := os.ReadDir(BasePath)
	check(err)

	segment_dirs := []string{}
	for _, file := range files {
		name := file.Name()

		if strings.HasPrefix(name, route) {
			segment_dirs = append(segment_dirs, name)
		}
	}
	sort.SliceStable(segment_dirs, func(i, j int) bool {
		split_seg_1 := strings.Split(segment_dirs[i], "--")
		seg_idx_1, _ := strconv.Atoi(split_seg_1[len(split_seg_1)-1])
		split_seg_2 := strings.Split(segment_dirs[j], "--")
		seg_idx_2, _ := strconv.Atoi(split_seg_2[len(split_seg_2)-1])
		return seg_idx_1 < seg_idx_2
	})
	return segment_dirs
}

func MonoDatetime(event Event, offset uint64, timezone string) time.Time {
	mono_time := event.LogMonoTime()
	wall_time := mono_time + offset
	t := time.UnixMicro(int64(wall_time / 1000))
	loc, _ := time.LoadLocation(timezone)
	t2 := time.Now().In(loc)
	_, zoneOffset := t2.Zone()
	return t.Add(time.Duration(-1 * zoneOffset)).UTC()
}

func MonoOffset(dat []byte) uint64 {
	offset := uint64(0)
	for offset < uint64(len(dat)) {
		msg, err := capnp.Unmarshal(dat[offset:])
		check(err)

		size, err := msg.TotalSize()
		check(err)
		offset += size

		evt, err := ReadRootEvent(msg)
		if err != nil {
			continue
		}
		if evt.Which() != Event_Which_clocks {
			continue
		}

		clocks, err := evt.Clocks()
		if err != nil {
			continue
		}
		wall := clocks.WallTimeNanos()
		mono := clocks.MonotonicNanos()
		return wall - mono
	}
	return 0
}
