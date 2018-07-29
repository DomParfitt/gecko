import { Component, Injectable } from '@angular/core';
import { AutomataService, Automata } from './automata.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
@Injectable()
export class AppComponent {
  title = 'Gecko';
  result: Automata;

  constructor(private automata: AutomataService) { }

  handleClick(pattern: string) {
    this.automata.getAutomata(pattern)
      .subscribe((data) => {
        this.result = { ...data };
      });
  }
}

