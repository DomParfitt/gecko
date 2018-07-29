import { Automata } from './automata.service';
import { Injectable } from '@angular/core';
import { HttpClient } from '../../node_modules/@angular/common/http';
import { Observable } from '../../node_modules/rxjs';

@Injectable({
  providedIn: 'root'
})
export class AutomataService {

  constructor(private http: HttpClient) { }

  getAutomata(pattern: string): Observable<Automata> {
    return this.http
      .get<Automata>('http://localhost:8080/pattern/' + pattern);
  }
}

export interface Automata {
  CurrentState: number;
  TerminalStates: number[];
  Transitions: any;
}
