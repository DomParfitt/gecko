import { Automata } from './automata.service';
import { Injectable } from '@angular/core';
import { HttpClient } from '../../node_modules/@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AutomataService {

  constructor(private http: HttpClient) { }

  getAutomata(pattern: string): Automata {
    let automata: Automata;
    this.http
      .get<Automata>('http://localhost:8080/pattern/' + pattern)
      .subscribe((data) => {
        automata = { ...data };
      });
      return automata;
  }
}

export interface Automata {
  CurrentState: number;
  TerminalStates: number[];
  Transitions: any;
}
