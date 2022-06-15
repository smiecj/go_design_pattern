package builder_

// director: 使用 builder 并生成实际数据
type Director struct {
}

func (d *Director) Construct(builder builder) {
	builder.makeTitle("teacher_day")
	builder.makeString("morning")
	builder.makeItems([]string{"early reading", "Chinese", "math"})
	builder.makeString("afternoon")
	builder.makeItems([]string{"physis", "chemical"})
	builder.makeString("evening")
	builder.makeItems([]string{"lesson prepare", "perusal"})
	builder.close()
}
