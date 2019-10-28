package generators

type Generatorer interface {
	GetName() string
	Run()
	Finalize()
}

type Generatorser interface {
	Add(generator Generatorer) Generatorer
	Find(name string) Generatorer
	Run()
}

type Generators struct {
	Store map[string]Generatorer
	Order []Generatorer
}

func (r *Generators) Add(generator Generatorer) Generatorer {
	r.Order = append(r.Order, generator)
	r.Store[generator.GetName()] = generator
	return generator
}

func (r *Generators) Find(name string) Generatorer {
	return r.Store[name]
}

func (r *Generators) Run() {
	for _, gen := range r.Order {
		gen.Run()
	}

	for i := range r.Order {
		gen := r.Order[len(r.Order)-1-i]
		gen.Finalize()
	}
}

var instance = &Generators{
	Store: map[string]Generatorer{},
}

func GetInstance() Generatorser {
	return instance
}
