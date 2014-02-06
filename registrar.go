package main


type Registrar struct{
	Id int
	boss *ArrivalsManager
	announce chan Guest
}

func ToRegistrar(r *Registrar) chan Guest {
	return r.announce
}

func NewRegistrar() *Registrar{
	r := new(Registrar)
	r.announce = make(chan Guest)
	return r
}

func (r *Registrar) Use(b *Browser) {
	fromBrowser := FromBrowser(b)
	toBrowser := ToBrowser(b)
	fromManager := ToRegistrar(r)
	toManager := ToManager(r.boss)
	toBrowser <- jsonable(r.boss.Chart)
	for{
		select {
			case g, ok := <-fromBrowser:
				if ok {
					toManager <- g
				} else {
					r.boss.Fire(r)
					fromBrowser = nil
				}
			case g, ok := <-fromManager:
				if ok {
					toBrowser <- g.Name()
				} else {
					close(toBrowser)
					return
				}
		}
	}
}




