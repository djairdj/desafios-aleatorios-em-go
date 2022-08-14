package functions

import (
	"aleatorios/functions"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Veiculo interface {
	LigarDesligar()
	MostrarNome()
}
type Carro struct {
	Name   string `json:"name"`
	Ano    int    `json:"ano"`
	Ligado bool   `json:"power"`
}

func (c Carro) MostrarNome() {
	fmt.Println(c.Name)
}
func (c Carro) LigarDesligar() {
	c.Ligado = !c.Ligado
}
func NomeVeiculo(veiculo Veiculo) {
	veiculo.MostrarNome()
}

type Task struct {
	Name string
	Done bool
}

func MainTreino() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/tasks", TaskHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{Done: true, Name: "Aprender Go"}
	j, _ := json.Marshal(task)
	t := template.Must(template.ParseFiles("task.html"))
	t.Execute(w, task)
	fmt.Println()
	w.Write(j)
}
func MainTreino2() {
	canal := make(chan int)
	for i := 1; i <= 1; i++ {
		go Work(canal)
	}
	for i := 1; i <= 35; i++ {
		canal <- i
	}
}
func Work(chanel chan int) {
	for i := range chanel {
		//fmt.Printf("%t \n", i)
		res := functions.FibonacciIterativo(i)
		fmt.Printf("Fib2 de %v = %v\n", i, res)
		//time.Sleep(time.Second * 2)
	}
}
