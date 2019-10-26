package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introdução a go Golang",
		Description: "Criando uma camama de serviço com Golang",
	},
}

func createEvent(response http.ResponseWriter, request *http.Request) {
	//Altera o retorno da request para application/json
	response.Header().Set("content-type", "application/json")
	//Cria um novo evento basedo no "tipo" event
	var newEvent event
	//Lê o body para requisição (request), em caso positivo,
	//os dados são transferidos como array de bytes para reqBody.
	//Se houver erro, o conteúdo de err não é nulo (nil)
	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Fprintf(response, "Informe dados o evento para cadastro.")
	}
	//Transfere o array de bytes no tipo event
	json.Unmarshal(reqBody, &newEvent)

	//Adiciona um evento no vetor
	events = append(events, newEvent)
	//Returna o status criado 201 para o cliente
	response.WriteHeader(http.StatusCreated)
	//Retorna o json de evento
	json.NewEncoder(response).Encode(newEvent)
}

func getOneEvent(response http.ResponseWriter, request *http.Request) {
	//Recebe o ID informado como path parameter
	eventID := mux.Vars(request)["id"]

	//pesquisa o vetor events
	for _, singleEvent := range events {
		//procura o event com ID dentro do vetor
		if singleEvent.ID == eventID {
			//retorna o resutado da pesquisa
			json.NewEncoder(response).Encode(singleEvent)
		}
	}
}

func getAllEvents(response http.ResponseWriter, request *http.Request) {
	//Altera o retorno da request para application/json
	response.Header().Set("content-type", "application/json")
	//retorna o vetor de registros
	json.NewEncoder(response).Encode(events)
}

func updateEvent(response http.ResponseWriter, request *http.Request) {
	//Altera o retorno da request para application/json
	response.Header().Set("content-type", "application/json")

	//Recebe o ID informado como path parameter
	eventID := mux.Vars(request)["id"]
	//Declara uma variável do tipo evento para atualização
	var updatedEvent event
	//Lê o body para requisição (request), em caso positivo,
	//os dados são transferidos como array de bytes para reqBody.
	//Se houver erro, o conteúdo de err não é nulo (nil)
	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Fprintf(response, "Informe os dados do evento.")
	}

	//Faz cópia do body recebido na requisição para a variável updateEvent
	json.Unmarshal(reqBody, &updatedEvent)

	//Pesquisa e atualiza o registro pelo ID.
	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(response).Encode(singleEvent)
		}
	}
}

func deleteEvent(response http.ResponseWriter, request *http.Request) {
	eventID := mux.Vars(request)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...) //remove índice
			fmt.Fprintf(response, "The event com ID %v foi deletado com sucesso.", eventID)
		}
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
