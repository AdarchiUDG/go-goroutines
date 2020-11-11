package main

import (
	"fmt"
	"time"
)

type Process struct {
	Signal chan bool
	End chan bool

	Terminated bool
}

func (p *Process) Start(id uint64) {
	show := false
	go func() {
		i := uint64(0)
		for {
			select {
				case <- p.End:
					return
				case show = <- p.Signal:

				default:
					if show {
						fmt.Println(id, ":", i)
					}
					i = i + 1
					time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}

func (p *Process) Stop() {
	if !p.Terminated {
		p.End <- true
		p.Terminated = true
	}
}

func (p *Process) Show() {
	if !p.Terminated {
		p.Signal <- true
	}
}

func (p *Process) Hide() {
	if !p.Terminated {
		p.Signal <- false
	}
}

func main() {
	var option uint
	var lastId uint64

	var processes []Process

	for {
		fmt.Println("1) Agregar Proceso")
		fmt.Println("2) Mostrar Procesos")
		fmt.Println("3) Terminar Proceso")
		fmt.Println("0) Salir")
		fmt.Scanf("%d", &option)
		fmt.Scanln()
		switch option {
			case 1:
				p := Process {
					Signal: make(chan bool),
					End: make(chan bool) }
				p.Start(lastId)
				processes = append(processes, p)
				lastId++
			case 2:
				for _, p := range processes {
					p.Show()
				}
				fmt.Scanln()	
				for k, p := range processes {
					fmt.Println(k)
					p.Hide()
				}
			case 3:
				var id uint64
				fmt.Scanf("%d", &id)
				fmt.Scanln()
				if (id < lastId) {
					processes[id].Stop()
					fmt.Println("Se elimino el proceso", id)
				}
			case 0:
				return
		}
	}
	
}