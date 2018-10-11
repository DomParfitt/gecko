package server

import (
	"encoding/json"
	"fmt"
	"github.com/DomParfitt/gecko/core"
	"github.com/DomParfitt/gecko/core/automata"
	"github.com/DomParfitt/gecko/server/api"
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

	response := &api.PatternResponse{}
	response.Automata = transformAutomata(compiler.Exe)
	response.AST = compiler.Ast.Transform()

	json, err := json.Marshal(response)
	fmt.Printf("%s", json)
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
	ok, steps := compiler.Exe.ExecuteStep(input)
	if err != nil {
		fmt.Fprintf(w, "Error")
	} else {
		json, err := json.Marshal(&api.MatchResponse{Pattern: pattern, Input: input, Result: ok, Steps: steps})
		if err != nil {
			fmt.Fprintf(w, "Error")
		}
		fmt.Printf("Returning result %s", json)
		fmt.Fprintf(w, "%s", json)
	}
}

func transformAutomata(exe *automata.FiniteState) api.Automata {
	a := api.Automata{CurrentNode: 0}

	states := []int{}
	edges := []api.Edge{}
	for from, transition := range exe.Transitions {
		if !contains(states, from) {
			states = append(states, from)
		}
		for ch, to := range transition {
			if !contains(states, to) {
				states = append(states, to)
			}

			edge := api.Edge{From: from, To: to, Label: string(ch)}
			edges = append(edges, edge)
		}
	}

	nodes := []api.Node{}
	for _, state := range states {
		node := api.Node{ID: state, IsTerminal: contains(exe.TerminalStates, state)}
		nodes = append(nodes, node)
	}

	a.Nodes = nodes
	a.Edges = edges

	return a
}

func contains(array []int, value int) bool {
	for _, present := range array {
		if present == value {
			return true
		}
	}

	return false
}
