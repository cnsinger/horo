package service

import (
	"fmt"
	"horo/model"
	"horo/storage"
	"horo/util"
	"strconv"
	"strings"
	"time"
)

const DefaultTimeFormat = "Mon, 02 Jan 2006 03:04PM"

func AddTimer(context []string) {
	s := storage.Instance()
	horoTimer := parseContext(context)
	s.InsertTimer(horoTimer)
	now := time.Now()
	fmt.Printf("[TIMER] End: %s, Delay: %s\n",
		horoTimer.DoneAt.Format(DefaultTimeFormat),
		horoTimer.DoneAt.Sub(now),
	)
}

func ShowTimer() {
	s := storage.Instance()
	s.Delete()
	horoTimers := s.Query()
	for _, h := range horoTimers {
		now := time.Now()
		if now.After(h.DoneAt) {
			s.Update([]int{h.Id})
			// update
			alert(h)
			continue
		}
		start := now.Unix() - h.InsertAt.Unix()
		desc := fmt.Sprintf("[#%d: %s]", h.Id, h.Context)
		util.ShowBar(desc, int(start), h.Length)
		alert(h)
	}
}

func DaemonRun() {
	select {}
}

func parseContext(context []string) model.HoroTimer {
	var length int
	for i := range context {
		if context[i] == "" {
			continue
		}
		if context[i][0] >= byte('0') && context[i][0] <= byte('9') {
			fields := strings.Split(context[i], ":")
			switch len(fields) {
			case 1: // 秒
				s, _ := strconv.Atoi(fields[0])
				length = s
			case 2: // 分秒
				m, _ := strconv.Atoi(fields[0])
				s, _ := strconv.Atoi(fields[1])
				length = m*60 + s
			case 3: // 时分秒
				h, _ := strconv.Atoi(fields[0])
				m, _ := strconv.Atoi(fields[1])
				s, _ := strconv.Atoi(fields[2])
				length = h*3600 + m*60 + s
			default:
			}
		}
	}
	now := time.Now()
	d := time.Duration(length) * time.Second
	return model.HoroTimer{
		Context:  strings.Join(context, " "),
		InsertAt: now,
		Length:   length,
		DoneAt:   now.Add(d),
	}
}

func alert(h *model.HoroTimer) {
}
