import * as React from 'react';
import { IAutomata, isAutomata } from './automata/Automata';
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

  }

  public render(): JSX.Element {
    return (
      <div className="App">
        <h1>Welcome to Gecko!</h1>
        <div>
          <input type="text" placeholder="Enter a pattern" onChange={
            // tslint:disable-next-line:jsx-no-lambda
            (event) => {
              const pattern = event.target.value;
              this.setState({'pattern': pattern});
              this.requestAutomata(pattern);
            }
          }/>
          <button onClick={
            // tslint:disable-next-line:jsx-no-lambda
            () => this.requestAutomata(this.state.pattern)
            } >Enter</button>
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

  private requestAutomata(pattern: string) {
    // tslint:disable-next-line:no-console
    console.log(pattern);
    fetch("http://localhost:8080/pattern/" + encodeURI(pattern))
      .then((resp) => resp.json())
      .then(
        (data) => {
          // tslint:disable-next-line:no-console
          console.log(data);
          if (isAutomata(data)) {
            this.handleAutomataData(data);
          }
        },
        (error) => {
          // tslint:disable-next-line:no-console
          console.log("Gecko Server Unavailable.", error)
        }
      );
  }

  private handleAutomataData(automata: IAutomata) {
    automata.nodes.sort();
    this.setState({"automata": automata});
  }

}

export interface IAppState extends React.ComponentState {
  input: string,
  pattern: string,
  matches: boolean,
  ast: any,
  automata: IAutomata 
}

export default App;