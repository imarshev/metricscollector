package server

type Counter struct {
	name  string
	value int64
}

type Gauge struct {
	name  string
	value float64
}

type Metric interface {
	GetValue() interface{}
	SetName(name string)
	UpdateValue(value interface{})
}

func (counter *Counter) UpdateValue(value interface{}) {
	if v, ok := value.(int64); ok {
		counter.value += v
	}
}

func (gauge *Gauge) UpdateValue(value interface{}) {
	if v, ok := value.(float64); ok {
		gauge.value = v
	}
}

func (counter *Counter) SetName(name string) {
	counter.name = name
}

func (gauge *Gauge) SetName(name string) {
	gauge.name = name
}

func (counter *Counter) GetValue() interface{} {
	return counter.value
}

func (gauge *Gauge) GetValue() interface{} {
	return gauge.value
}
