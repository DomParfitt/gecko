import * as React from 'react';
import { IAutomata } from './../automata/Automata';
import { AutomataGraph } from './AutomataGraph';

export class AutomataViewer extends React.Component<IAutomataViewerProps, IAutomataViewerState> {

    constructor(props: IAutomataViewerProps) {
        super(props);
        this.state = {
            charPointer: 0,
            currentNode: 0,
            flattenEdges: false,
            previousNode: 0,
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
                <button disabled={this.isBackwardButtonDisabled()} onClick={this.stepBackward}>Back</button>
                <button disabled={this.isForwardButtonDisabled()} onClick={this.stepForward}>Forward</button>
                <AutomataGraph
                    automata={this.props.automata}
                    currentNode={this.state.currentNode}
                    previousNode={this.state.previousNode}
                    flattenEdges={this.state.flattenEdges}
                    consumed={this.props.input ? this.props.input[this.state.charPointer] : undefined}
                />
            </div>
        );
    }


    private setFlattenEdges = (event: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ flattenEdges: !this.state.flattenEdges });
    }

    private stepForward = (event: React.MouseEvent<HTMLButtonElement>) => {
        if (!this.props.input) {
            return;
        }

        for (const edge of this.props.automata.edges) {
            if (edge.from === this.state.currentNode
                && edge.label === this.props.input[this.state.charPointer]) {
                this.setState({
                    charPointer: this.state.charPointer + 1,
                    currentNode: edge.to,
                    previousNode: edge.from,
                });
                break;
            }
        }
    }

    private stepBackward = (event: React.MouseEvent<HTMLButtonElement>) => {
        if (!this.props.input) {
            return;
        }

        for (const edge of this.props.automata.edges) {
            if (edge.to === this.state.currentNode
                && edge.label === this.props.input[this.state.charPointer - 1]) {
                this.setState({
                    charPointer: this.state.charPointer - 1,
                    currentNode: edge.from,
                    previousNode: edge.from,
                });
                break;
            }
        }
    }

    private isForwardButtonDisabled(): boolean {
        if (!this.props.input) {
            return true;
        }
        return this.state.charPointer >= this.props.input.length;
    }

    private isBackwardButtonDisabled(): boolean {
        if (!this.props.input) {
            return true;
        }

        return this.state.charPointer <= 0;
    }

}

export interface IAutomataViewerProps extends React.ClassAttributes<AutomataViewer> {
    automata: IAutomata,
    input?: string,
}

export interface IAutomataViewerState extends React.ComponentState {
    charPointer: number,
    currentNode: number,
    flattenEdges: boolean,
    previousNode: number,
}