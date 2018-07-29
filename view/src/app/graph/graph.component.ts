import { Automata } from './../automata.service';
import { EdgeComponent } from './edge/edge.component';
import { VertexComponent } from './vertex/vertex.component';
import { Component } from '@angular/core';

@Component({
  selector: 'app-graph',
  templateUrl: './graph.component.html',
  styleUrls: ['./graph.component.scss']
})
export class GraphComponent {

  vertices: VertexComponent[];
  edges: EdgeComponent[];

  constructor() {
    this.vertices = [];
    this.edges = [];
    // this.edges.push(new EdgeComponent());
  }

  displayAutomata(automata: Automata) {
    automata.States.forEach((state, index) => {
      const v = this.createVertex(state, automata.TerminalStates);
      v.x = index * 100;
      this.vertices.push(v);
    });

    console.log(this.vertices);

  }

  private createVertex(label: number, terminals: number[]): VertexComponent {
    const v = new VertexComponent();
    v.label = String(label);
    console.log('Label: ' + label + ' isTerminal = ' + terminals.includes(label));
    if (terminals.includes(label)) {
      v.isTerminal = true;
    }
    return v;
  }

  private createEdge(): EdgeComponent {
    return null;
  }

}
