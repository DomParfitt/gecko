import * as React from 'react';
import { IAutomata } from './../automata/Automata';
import { AutomataGraph } from './AutomataGraph';

export class AutomataViewer extends React.Component<IAutomataViewerProps, IAutomataViewerState> {

    constructor(props: IAutomataViewerProps) {
        super(props);
        this.state = { 
            currentNode: 0,
            flattenEdges: false, 
        };
    }

    public render(): JSX.Element {
        return (
            <div>
                <div>
                    <label>
                        Flatten Edges?
                    <input type="checkbox" onChange={this.setFlattenEdges} />
                    </label>
                </div>
                <button onClick={this.stepBackward}>Back</button>
                <button onClick={this.stepForward}>Forward</button>
                <AutomataGraph automata={this.props.automata} currentNode={this.state.currentNode} flattenEdges={this.state.flattenEdges} />
            </div>
        );
    }

    private setFlattenEdges = (event: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ flattenEdges: !this.state.flattenEdges });
    }

    private stepForward = (event: React.MouseEvent<HTMLButtonElement>) => {
        this.setState({ currentNode: this.state.currentNode + 1 });
    }

    private stepBackward = (event: React.MouseEvent<HTMLButtonElement>) => {
        this.setState({ currentNode: this.state.currentNode - 1 }); 
    }

}

export interface IAutomataViewerProps extends React.ClassAttributes<AutomataViewer> {
    automata: IAutomata
}

// tslint:disable-next-line:no-empty-interface
export interface IAutomataViewerState extends React.ComponentState {
    currentNode: number,
    flattenEdges: boolean
}