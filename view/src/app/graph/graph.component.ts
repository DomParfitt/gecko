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
    this.reset();
  }

  private reset() {
    this.vertices = [];
    this.edges = [];
  }

  displayAutomata(automata: Automata) {
    this.reset();
    // automata.States.sort();
    // automata.States.forEach((state, index) => {
    //   const v = this.createVertex(state, automata.TerminalStates);
    //   v.x = index * 100;
    //   this.vertices.push(v);
    // });

    this.createVertices(automata);

    this.createEdges(automata);

    // let i = 1;
    // for (const from in automata.Transitions) {
    //   if (automata.Transitions.hasOwnProperty(from)) {
    //     for (const char in automata.Transitions[from]) {
    //       if (automata.Transitions[from].hasOwnProperty(char)) {
    //         // console.log("From: " + from);
    //         // console.log("Over: " + char);
    //         // console.log("To: " + automata.Transitions[from][char]);
    //         const e = this.createEdge(from, automata.Transitions[from][char], char);
    //         e.y = i * 100;
    //         this.edges.push(e);
    //         i++;
    //       }
    //     }
    //   }
    // }
  }

  private createVertices(automata: Automata) {
    automata.States.sort();
    automata.States.forEach((state, index) => {
      const v = this.createVertex(state, automata.TerminalStates);
      v.x = index * 200;
      if (automata.CurrentState === state) {
        v.isCurrent = true;
      }
      this.vertices.push(v);
    });
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

  private createEdges(automata: Automata) {
    let i = 1;
    for (const from in automata.Transitions) {
      if (automata.Transitions.hasOwnProperty(from)) {
        for (const char in automata.Transitions[from]) {
          if (automata.Transitions[from].hasOwnProperty(char)) {
            const e = this.createEdge(from, automata.Transitions[from][char], char);
            console.log(from);
            e.y = -25;
            e.x = (automata.States.indexOf(Number(from)) * 200) + 100;
            this.edges.push(e);
            i++;
          }
        }
      }
    }
  }

  private createEdge(from: string, to: string, label: string): EdgeComponent {
    const e = new EdgeComponent();
    e.from = String(from);
    e.to = String(to);
    e.label = label;
    return e;
  }

}
