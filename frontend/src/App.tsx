import * as React from 'react';
import './App.css'
import { IAbstractSyntaxTree } from './ast/AbstractSyntaxTree';
import { getFlattenedEdges, IAutomata } from './automata/Automata';
import { ASTGraph } from './graph/ASTGraph';
import { AutomataGraph } from './graph/AutomataGraph';
import { TextInput } from './input/TextInput';

class App extends React.Component<IAppProps, IAppState> {

    constructor(props: any) {
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
                currentNode: 0,
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
            flattenEdges: false,
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
                <label>
                    Flatten Edges?
                    <input type="checkbox" onChange={this.setFlattenEdges}/>
                </label>
                <div>Pattern: {this.state.pattern}</div>
                <div>Input: {this.state.input}</div>
                <div>Matches: {this.state.matches.toString()}</div>
                <AutomataGraph automata={this.state.automata} flattenEdges={this.state.flattenEdges} />
                <ASTGraph ast={this.state.ast} />
                {/* <GraphsHolder automata={this.state.automata} ast={this.state.ast} /> */}
            </div>
        );
    }

    private handlePatternChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const pattern = event.target.value;
        this.setState({ 'pattern': pattern });
        this.requestPattern(pattern);
    }

    private handlePatternClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        this.requestPattern(this.state.pattern);
    }

    private handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const input = event.target.value;
        this.setState({ 'input': input });
    }

    private setFlattenEdges = (event: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ flattenEdges: !this.state.flattenEdges });
    }

    private requestPattern(pattern: string) {
        this.log(pattern);
        fetch("http://localhost:8080/pattern/" + encodeURI(pattern))
            .then((resp) => resp.json())
            .then(
                (data) => {
                    this.log(data);
                    if (this.state.flattenEdges) {
                        getFlattenedEdges(data.automata);
                    }
                    this.setState({ automata: data.automata, ast: data.ast })
                },
                (error) => {
                    this.log("Gecko Server Unavailable. " + error)
                }
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
    flattenEdges: boolean
}

export default App;