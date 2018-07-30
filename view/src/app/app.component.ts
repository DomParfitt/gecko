import { GraphComponent } from './graph/graph.component';
import { Component, Injectable, ViewChild, AfterViewInit } from '@angular/core';
import { AutomataService, Automata } from './automata.service';
import { NxGraphComponent } from './nx-graph/nx-graph.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
@Injectable()
export class AppComponent implements AfterViewInit {

  title = 'Gecko';
  result: Automata;

  @ViewChild(GraphComponent)
  graph: GraphComponent;

  @ViewChild(NxGraphComponent)
  nxGraph: NxGraphComponent;

  constructor(private automata: AutomataService) { }

  handleClick(pattern: string) {
    this.automata.getAutomata(pattern)
      .subscribe((data) => {
        this.result = { ...data };
        this.graph.displayAutomata(this.result);
        this.nxGraph.displayAutomata(this.result);
      });
  }

  ngAfterViewInit(): void {
    // throw new Error('Method not implemented.');
  }
}

