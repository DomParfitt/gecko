import * as React from 'react';
import { IAutomata } from './automata/Automata';
import Graph from './graph/Graph';

class App extends React.Component<any, IAppState> {

  constructor(props: any) {
    super(props);

    this.state = {
      ast: {},
      automata: {
        currentNode: 0,
        edges: [
          { from: 0, to: 1, label: 'a' },
          { from: 1, to: 2, label: 'b' },
          { from: 2, to: 1, label: 'c' },
        ],
        nodes: [
          { id: 0, isTerminal: false },
          { id: 1, isTerminal: true },
          { id: 2, isTerminal: false },
        ],
      },
      input: "",
      matches: false,
      pattern: "abc",
    };

    // this.handleAutomataData = this.handleAutomataData.bind(this);
  }

  public render() {
    return (
      <div className="App">
        <h1>Welcome to Gecko!</h1>
        <div>
          <input type="text" placeholder="Enter a pattern" />
          <button>Enter</button>
        </div>
        <div>
          <input type="text" placeholder="Enter an input" />
          <button>Enter</button>
        </div>
        <div>Pattern: {this.state.pattern}</div>
        <div>Input: {this.state.input}</div>
        <div>Matches: {this.state.matches.toString()}</div>
        <Graph currentNode={this.state.automata.currentNode} nodes={this.state.automata.nodes} edges={this.state.automata.edges} />
      </div>
    );
  }

  // private requestAutomata(pattern: string) {
  //   fetch("http://localhost:8080/pattern/" + encodeURI(pattern))
  //     .then((resp) => resp.json())
  //     .then(
  //       (data) => {
  //         this.setState(state => state.pattern = pattern);
  //         this.handleAutomataData(data);
  //       },
  //       (error) => {
  //         console.log("Gecko Server Unavailable.")
  //       }
  //     );
  // }

  // private handleAutomataData(automata: any) {

  //   const newNodes = [];
  //   for (const state of automata.States) {
  //     const terminal = automata.TerminalStates.includes(state);
  //     newNodes.push({ id: state, isTerminal: terminal });
  //   }

  //   const newEdges = [];
  //   const transitions = automata.Transitions;
  //   for (const from in transitions) {
  //     for (const over in transitions[from]) {
  //       const to = transitions[from][over];
  //       newEdges.push({ from: from, to: to, label: over });
  //     }
  //   }

  //   const newAutomata: IAutomata = {
  //     currentNode: automata.CurrentState,
  //     nodes: newNodes.sort(),
  //     edges: newEdges
  //   };

  //   this.setState({automata: newAutomata});
  // }

}

export interface IAppState extends React.ComponentState {
  input: string,
  pattern: string,
  matches: boolean,
  ast: any,
  automata: IAutomata 
}

export default App;