import * as React from 'react';
import './App.css'
import { IAbstractSyntaxTree } from './ast/AbstractSyntaxTree';
import { IAutomata } from './automata/Automata';
import { ASTGraph } from './graph/ASTGraph';
import { AutomataViewer } from './graph/AutomataViewer';
import { TextInput } from './input/TextInput';

class App extends React.Component<IAppProps, IAppState> {

    constructor(props: IAppProps) {
        super(props);

        this.state = {
            ast: {
                children: [
                    {
                        children: [
                            {
                                children: [],
                                label: 'a'
                            },
                            {
                                children: [],
                                label: 'b'
                            }
                        ],
                        label: ''
                    },
                    {
                        children: [],
                        label: '|'
                    },
                    {
                        children: [],
                        label: 't'
                    }
                ],
                label: 'root',
            },
            automata: {
                edges: [
                    { from: 0, to: 1, label: 'a' },
                    { from: 1, to: 2, label: 'b' },
                    { from: 2, to: 1, label: 'c' },
                ],
                nodes: [
                    { id: 0, isTerminal: false },
                    { id: 1, isTerminal: true },
                    { id: 2, isTerminal: false },
                ],
            },
            input: "",
            matches: false,
            pattern: "abc",
        };
    }

    public render(): JSX.Element {
        return (
            <div className="App">
                <h1>Welcome to Gecko!</h1>
                <TextInput placeholder="Enter a pattern" onChangeHandler={this.handlePatternChange} onClickHandler={this.handlePatternClick} hideButton={true} />
                <TextInput placeholder="Enter an input" onChangeHandler={this.handleInputChange} />
                <div>Pattern: {this.state.pattern}</div>
                <div>Input: {this.state.input}</div>
                <div>Matches: {this.state.matches.toString()}</div>
                <AutomataViewer automata={this.state.automata} steps={[0,1,2,1,2,1]}/>
                <ASTGraph ast={this.state.ast} />
                {/* <GraphsHolder automata={this.state.automata} ast={this.state.ast} /> */}
            </div>
        );
    }

    private handlePatternChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const pattern = event.target.value;
        this.setState({ pattern });
        this.requestPattern(pattern);
    }

    private handlePatternClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        this.requestPattern(this.state.pattern);
    }

    private handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const input = event.target.value;
        const pattern = this.state.pattern;
        this.setState({ input });
        this.requestMatch(pattern, input);
    }

    private requestPattern(pattern: string) {
        this.log(pattern);
        fetch("http://localhost:8080/pattern/" + encodeURI(pattern))
            .then((resp) => resp.json())
            .then(
                (data) => {
                    this.log(data);
                    this.setState({ automata: data.automata, ast: data.ast })
                },
                (error) => {
                    this.log("Gecko Server Unavailable. " + error)
                }
            );
    }

    private requestMatch(pattern: string, input: string) {
        fetch("http://localhost:8080/match/" + encodeURI(pattern) + "/" + encodeURI(input))
            .then((resp) => resp.json())
            .then(
                (data) => {
                    this.log(data);
                    this.setState({matches: data.result});
                },
                (error) => {this.log("Gecko Server Unavailable. " + error)}
            );
    }

    private log(message: string) {
        if (this.props.debugMode) {
            // tslint:disable-next-line:no-console
            console.log(message);
        }
    }

}

export interface IAppProps extends React.ClassAttributes<App> {
    debugMode: boolean;
}

export interface IAppState extends React.ComponentState {
    input: string,
    pattern: string,
    matches: boolean,
    ast: IAbstractSyntaxTree,
    automata: IAutomata,
}

export default App;