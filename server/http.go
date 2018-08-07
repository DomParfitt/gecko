package server

import (
	"encoding/json"
	"fmt"
	"github.com/DomParfitt/gecko/core"
	"github.com/DomParfitt/gecko/core/automata"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Serve http content on given port
func Serve(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/pattern/{pattern}", patternHandler)
	router.HandleFunc("/match/{pattern}/{input}", matchHandler)
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func patternHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	pattern := vars["pattern"]
	fmt.Printf("Received request for pattern: %s\n", pattern)

	compiler := core.New()
	compiler.Compile(pattern)
	json, err := marshal(compiler.Exe)
	if err != nil {
		fmt.Fprintf(w, "Error")
	} else {
		fmt.Fprintf(w, "%s", json)
	}
}

func matchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	pattern := vars["pattern"]
	input := vars["input"]
	fmt.Printf("Received request to match input %s against pattern: %s\n", input, pattern)

	compiler := core.New()
	ok, err := compiler.MatchPattern(pattern, input)
	if err != nil {
		fmt.Fprintf(w, "Error")
	} else {
		json, err := json.Marshal(&matchResponse{Pattern: pattern, Input: input, Result: ok})
		if err != nil {
			fmt.Fprintf(w, "Error")
		}
		fmt.Printf("Returning result %s", json)
		fmt.Fprintf(w, "%s", json)
	}
}

func marshal(exe *automata.FiniteState) ([]byte, error) {
	a := &jsonAutomata{
		CurrentState:   exe.CurrentState,
		TerminalStates: exe.TerminalStates,
	}

	states := []int{}
	transitions := make(map[int]map[string]int)
	for from, transition := range exe.Transitions {
		if !contains(states, from) {
			states = append(states, from)
		}
		_, ok := transitions[from]
		if !ok {
			transitions[from] = make(map[string]int)
		}
		for ch, to := range transition {
			if !contains(states, to) {
				states = append(states, to)
			}
			transitions[from][string(ch)] = to
		}
	}

	a.States = states
	a.Transitions = transitions

	return json.Marshal(a)
}

func contains(array []int, value int) bool {
	for _, present := range array {
		if present == value {
			return true
		}
	}

	return false
}

type jsonAutomata struct {
	CurrentState   int
	TerminalStates []int
	States         []int
	Transitions    map[int]map[string]int
}

type matchResponse struct {
	Pattern string
	Input   string
	Result  bool
}
