import { Component, Injectable } from '@angular/core';
import { HttpClient } from '../../node_modules/@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
@Injectable()
export class AppComponent {
  title = 'Gecko';
  result: Automata;

  constructor(private http: HttpClient) { }

  handleClick(pattern: string) {
    // let json;
    this.http
      .get<Automata>('http://localhost:8080/pattern/' + pattern)
      .subscribe((data) => {
        this.result = { ...data };
        console.log(this.result);
      });
  }
}

export interface Automata {
  CurrentState: number;
  TerminalStates: number[];
  Transitions: any;
}
