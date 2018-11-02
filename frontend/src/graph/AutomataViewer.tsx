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
                <button disabled={this.isBackwardButtonDisabled()} onClick={this.stepBackward}>Back</button>
                <button disabled={this.isForwardButtonDisabled()} onClick={this.stepForward}>Forward</button>
                <AutomataGraph
                    automata={this.props.automata}
                    currentNode={this.props.steps ? this.props.steps[this.state.currentNode] : undefined}
                    previousNode={this.state.previousNode}
                    flattenEdges={this.state.flattenEdges}
                />
            </div>
        );
    }


    private setFlattenEdges = (event: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ flattenEdges: !this.state.flattenEdges });
    }

    private stepForward = (event: React.MouseEvent<HTMLButtonElement>) => {
        this.setState({
            currentNode: this.state.currentNode + 1,
            previousNode: this.props.steps ? this.props.steps[this.state.currentNode] : undefined,
        });
    }

    private stepBackward = (event: React.MouseEvent<HTMLButtonElement>) => {
        this.setState({
            currentNode: this.state.currentNode - 1,
            previousNode: this.props.steps ? this.props.steps[this.state.currentNode] : undefined,
        });
    }

    private isForwardButtonDisabled(): boolean {
        if (!this.props.steps) {
            return true;
        }
        return this.state.currentNode >= this.props.steps.length - 1;
    }

    private isBackwardButtonDisabled(): boolean {
        if (!this.props.steps) {
            return true;
        }

        return this.state.currentNode <= 0;
    }

}

export interface IAutomataViewerProps extends React.ClassAttributes<AutomataViewer> {
    automata: IAutomata,
    steps?: number[],
}

export interface IAutomataViewerState extends React.ComponentState {
    currentNode: number,
    flattenEdges: boolean,
    previousNode?: number,
}