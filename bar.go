package bar

import (
	"fmt"
	"strings"
)

const (
	defaultGraph = "█"
	defaultStep  = 1
)

type Bar struct {
	cur         int    //当前进度位置
	total       int    //总进度
	step        int    //步长
	percent     int    //百分比
	graph       string //显示符号
	progressBar string //进度条
}

func (bar *Bar) NewBar(start, total int) {
	bar.cur = start
	bar.total = total
	bar.step = defaultStep
	bar.percent = bar.getPercent()
	bar.graph = defaultGraph
}

func (bar *Bar) SetGraph(graph string) {
	if bar.graph == "" {
		bar.graph = defaultGraph
	} else {
		bar.graph = graph
	}
}

func (bar *Bar) SetStep(step int) {
	if bar.step < 1 {
		bar.step = defaultStep
	} else {
		bar.step = step
	}
}

func (bar *Bar) Play() {
	bar.percent = bar.getPercent()
	bar.progressBar = strings.Repeat(bar.graph, bar.percent/2)

	fmt.Printf("\r[%-50s]%3d%%  %8d/%d", bar.progressBar, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) getPercent() int {
	return int(float32(bar.cur) / float32(bar.total) * 100)
}

func (bar *Bar) Increment() {
	if bar.cur+bar.step > bar.total {
		bar.cur = bar.total
	} else {
		bar.cur += bar.step
	}
}

func (bar *Bar) Finish() {
	fmt.Println()
}
