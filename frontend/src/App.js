import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

import Graph from 'react-graph-vis';

class App extends Component {
  handleClick(pattern) {
    console.log(pattern);
  }
  render() {
    var graph = {
      nodes: [
          {id: 1, label: 'Node 1'},
          {id: 2, label: 'Node 2'},
          {id: 3, label: 'Node 3'},
          {id: 4, label: 'Node 4'},
          {id: 5, label: 'Node 5'}
        ],
      edges: [
          {from: 1, to: 2},
          {from: 1, to: 3},
          {from: 2, to: 4},
          {from: 2, to: 5}
        ]
    };
    return (
      <div className="App">
        <input ref="inputBox" type="text" placeholder="Enter a pattern"></input>
        <button onClick={() => this.handleClick(this.refs.inputBox.value)}>Enter</button>
        <Graph graph={graph}/>
      </div>
    );
  }
}

export default App;
