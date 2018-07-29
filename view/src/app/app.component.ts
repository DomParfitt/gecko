import { GraphComponent } from './graph/graph.component';
import { Component, Injectable, ViewChild, AfterViewInit } from '@angular/core';
import { AutomataService, Automata } from './automata.service';

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

  constructor(private automata: AutomataService) { }

  handleClick(pattern: string) {
    this.automata.getAutomata(pattern)
      .subscribe((data) => {
        this.result = { ...data };
        this.graph.displayAutomata(this.result);
      });
  }

  ngAfterViewInit(): void {
    // throw new Error('Method not implemented.');
  }
}

