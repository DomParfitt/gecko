import React, { Component } from 'react';
import './App.css';
import Graph from './graph/Graph';

class App extends Component {

  constructor(props) {
    super(props);

    this.state = {
      input: "",
      pattern: "abc",
      ast: {},
      automata: {
        currentNode: 0,
        nodes: [
          { id: 0, isTerminal: false },
          { id: 1, isTerminal: true },
          { id: 2, isTerminal: false },
        ],
        edges: [
          { from: 0, to: 1, label: 'a' },
          { from: 1, to: 2, label: 'b' },
          { from: 2, to: 1, label: 'c' },
        ]
      }
    };

    this.handleAutomataData = this.handleAutomataData.bind(this);
  }

  handleClick(pattern) {
    console.log(pattern);
    fetch("http://localhost:8080/pattern/" + pattern)
      .then((resp) => resp.json())
      .then(
        (data) => {
          console.log(data);
          this.setState(state => state.pattern = pattern);
          this.handleAutomataData(data);
        },
        (error) => {
          console.log("Gecko Server Unavailable.")
        }
      );
  }

  handleAutomataData(automata) {

    const newNodes = [];
    for (const state of automata.States) {
      const terminal = automata.TerminalStates.includes(state);
      newNodes.push({ id: state, isTerminal: terminal });
    }

    const newEdges = [];
    const transitions = automata.Transitions;
    for (const from in transitions) {
      for (const over in transitions[from]) {
        const to = transitions[from][over];
        newEdges.push({ from: from, to: to, label: over });
      }
    }

    const newAutomata = {
      currentNode: automata.CurrentState,
      nodes: newNodes.sort(),
      edges: newEdges
    };

    this.setState(state => state.automata = newAutomata);
  }

  render() {
    return (
      <div className="App">
        <h1>Welcome to Gecko!</h1>
        <div>
          <input ref="patternBox" type="text" placeholder="Enter a pattern" onChange={() => this.handleClick(this.refs.patternBox.value)}></input>
          <button onClick={() => this.handleClick(this.refs.patternBox.value)}>Enter</button>
        </div>
        <div>
          <input ref="inputBox" type="text" placeholder="Enter an input" onChange={() => console.log("not yet implemented")}></input>
          <button onClick={() => console.log("not yet implemented")}>Enter</button>
        </div>
        <div>Pattern: {this.state.pattern}</div>
        <div>Input: {this.state.input}</div>
        <Graph currentNode={this.state.automata.currentNode} nodes={this.state.automata.nodes} edges={this.state.automata.edges} />
      </div>
    );
  }
}

export default App;
