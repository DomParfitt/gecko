package server

import (
	"encoding/json"
	"fmt"
	"github.com/DomParfitt/gecko/automata"
	"github.com/DomParfitt/gecko/compiler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Run the server
func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/pattern/{pattern}", patternHandler)
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func patternHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type")

	vars := mux.Vars(r)
	pattern := vars["pattern"]
	fmt.Printf("Received request for pattern: %s\n", pattern)

	compiler := compiler.New()
	compiler.Compile(pattern)
	json, err := marshal(compiler.Exe)
	if err != nil {
		fmt.Fprintf(w, "Error")
	} else {
		fmt.Fprintf(w, "%s", json)
	}
}

func marshal(exe *automata.FiniteState) ([]byte, error) {
	a := &jsonAutomata{
		CurrentState:   exe.CurrentState,
		TerminalStates: exe.TerminalStates,
	}

	transitions := make(map[int]map[string]int)
	for from, transition := range exe.Transitions {
		_, ok := transitions[from]
		if !ok {
			transitions[from] = make(map[string]int)
		}
		for ch, to := range transition {
			transitions[from][string(ch)] = to
		}
	}

	a.Transitions = transitions

	return json.Marshal(a)
}

type jsonAutomata struct {
	CurrentState   int
	TerminalStates []int
	Transitions    map[int]map[string]int
}
