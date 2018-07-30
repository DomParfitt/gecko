import { Component, OnInit } from '@angular/core';
import { Automata } from '../automata.service';

@Component({
  selector: 'app-nx-graph',
  templateUrl: './nx-graph.component.html',
  styleUrls: ['./nx-graph.component.scss']
})

export class NxGraphComponent implements OnInit {

  hierarchialGraph = { nodes: [], links: [] };

  constructor() {

  }

  ngOnInit() {

  }

  displayAutomata(automata: Automata) {
    automata.States.sort();
    automata.States.forEach((state, index) => {
      const v = {
        id: state,
        label: state
      };
      this.hierarchialGraph.nodes.push(v);
    });
  }
}
