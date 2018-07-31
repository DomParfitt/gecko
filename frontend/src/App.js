import React, { Component } from 'react';
import './App.css';
import Graph from './graph/Graph';

class App extends Component {

  constructor(props) {
    super(props);

    this.state = {
      currentNode: 0,
      nodes: [
        { id: 0, isTerminal: false },
        { id: 1, isTerminal: true },
        { id: 2, isTerminal: false },
      ],
      edges: [
        { from: 0, to: 1, label: 'a' },
        { from: 1, to: 2, label: 'b' },
        { from: 2, to: 1, label: 'd' },
      ]
    };

    this.handleAutomataData =this.handleAutomataData.bind(this);
  }

  handleClick(pattern) {
    console.log(pattern);
    fetch("http://localhost:8080/pattern/" + pattern)
      .then((resp) => resp.json())
      .then((data) => { console.log(data); this.handleAutomataData(data) })
  }

  handleAutomataData(automata) {

    const newNodes = [];
    for(const state of automata.States) {
      const terminal = automata.TerminalStates.includes(state);
      newNodes.push({id: state, isTerminal: terminal});
    }

    const newEdges = [];
    const transitions = automata.Transitions;
    for(const from in transitions) {
      for(const over in transitions[from]) {
        const to = transitions[from][over];
        newEdges.push({from: from, to: to, label: over});
      }
    }

    this.setState({
      currentNode: automata.CurrentState,
      nodes: newNodes.sort(),
      edges: newEdges
    });
    console.log(this.state);
  }

  render() {
    console.log("parent rendered");
    return (
      <div className="App">
        <input ref="inputBox" type="text" placeholder="Enter a pattern"></input>
        <button onClick={() => this.handleClick(this.refs.inputBox.value)}>Enter</button>
        <Graph currentNode={this.state.currentNode} nodes={this.state.nodes} edges={this.state.edges} />
      </div>
    );
  }
}

export default App;
