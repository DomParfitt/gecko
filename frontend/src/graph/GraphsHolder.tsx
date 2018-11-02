import * as React from 'react';
import { IAbstractSyntaxTree } from 'src/ast/AbstractSyntaxTree';
import { IAutomata } from 'src/automata/Automata';
import { ASTGraph } from './ASTGraph';
import { AutomataGraph } from './AutomataGraph';
import { ISwitcherData, Switcher } from './Switcher';

export class GraphsHolder extends React.Component<IGraphsHolderProps, IGraphsHolderState> {

    constructor(props: IGraphsHolderProps) {
        super(props);
    }

    public componentWillMount() {
        this.showAutomataGraph();
    }

    public render(): JSX.Element {
        return (
            <div>
                {this.state.currentGraph()}
                <Switcher graphs={this.switcherData()}/>
            </div>
        );
    }

    private switcherData(): ISwitcherData[] {
        const data = [];
        data.push({name: 'Automata', selector: this.showAutomataGraph})
        data.push({name: 'AST', selector: this.showASTGraph})
        return data;
    }

    private showAutomataGraph = () => {
        this.setState({'currentGraph': this.renderAutomataGraph});
    }

    private showASTGraph = () => {
        this.setState({'currentGraph': this.renderASTGraph});
    }

    private renderAutomataGraph = (): JSX.Element => {
        return (
            <AutomataGraph automata={this.props.automata}/>
        );
    }

    private renderASTGraph = (): JSX.Element => {
        return (
            <ASTGraph ast={this.props.ast} />
        );
    }
}

export interface IGraphsHolderState extends React.ComponentState {
    currentGraph: () => JSX.Element
}

export interface IGraphsHolderProps extends React.ClassAttributes<GraphsHolder> {
    automata: IAutomata
    ast: IAbstractSyntaxTree
}
