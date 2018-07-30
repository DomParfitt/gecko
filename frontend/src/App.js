import React, { Component } from 'react';
import './App.css';
import VisGraph  from './vis/VisGraph';

class App extends Component {
  handleClick(pattern) {
    console.log(pattern);
    fetch("http://localhost:8080/pattern/" + pattern)
      .then((resp) => resp.json())
      .then((data) => console.log(data))
  }
  render() {

    return (
      <div className="App">
        <input ref="inputBox" type="text" placeholder="Enter a pattern"></input>
        <button onClick={() => this.handleClick(this.refs.inputBox.value)}>Enter</button>
        <VisGraph />
      </div>
    );
  }
}

export default App;
