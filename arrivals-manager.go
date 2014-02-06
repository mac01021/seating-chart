package main



type ArrivalsManager struct{
	Chart Chart

	arrivals []Guest
	Announce chan Guest

	registrars map[int]*Registrar
	hireChan chan *Registrar
	fireChan chan *Registrar
	nbHired int
}

func NewArrivalsManager(chart Chart, arrivals []Guest) *ArrivalsManager{
	man := new(ArrivalsManager)
	man.Chart = chart
	man.arrivals = arrivals
	man.Announce = make(chan Guest)
	man.registrars = make(map[int]*Registrar)
	man.hireChan = make(chan *Registrar)
	man.fireChan = make(chan *Registrar)
	return man
}

func ToManager(m *ArrivalsManager) chan Guest{
	return m.Announce
}



func (man *ArrivalsManager) Run() {
	go func(){
		for{
			select{
				case r := <-man.hireChan:
					man.registrars[man.nbHired] = r
					r.Id = man.nbHired
					r.boss = man
					for _, g := range man.arrivals {
						ToRegistrar(r) <- g
					}
					man.nbHired += 1
				case r := <-man.fireChan:
					close(ToRegistrar(r))
					delete(man.registrars, r.Id)
				case g := <-ToManager(man):
					man.arrivals = append(man.arrivals, g)
					for _, r := range man.registrars{
						ToRegistrar(r) <- g
					}
			}
		}
	}()
}

func (man *ArrivalsManager) Hire(r *Registrar) *Registrar{
	man.hireChan <- r
	return r
}

func (man *ArrivalsManager) Fire(r *Registrar){
	man.fireChan <- r
}


