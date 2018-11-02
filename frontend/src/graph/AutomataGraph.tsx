import * as React from 'react';
import {  getFlattenedEdges, IAutomata, sortEdges } from 'src/automata/Automata';
import { Graphviz } from './Graphviz';

export class AutomataGraph extends React.Component<IAutomataGraphProps, any> {

    public render(): JSX.Element {
        return(
            <Graphviz dot={this.generateDot()} />
        );
    }

    private generateDot(): string {
        let dot = 'digraph { rankdir="LR";\n';

        for(const node of this.props.automata.nodes) {
            dot += node.id +' [';

            if (node.isTerminal) {
                dot += 'shape="doublecircle" ';
            } else {
                dot += 'shape="circle" ';
            }

            if (node.id === this.props.currentNode) {
                dot += 'fillcolor="red" style="filled" ';
            }

            dot += '];\n';
        }

        sortEdges(this.props.automata);
        let edges = this.props.automata.edges;
        if (this.props.flattenEdges) {
            edges = getFlattenedEdges(this.props.automata);
        }

        for(const edge of edges) {
            dot += edge.from + '->' + edge.to + ' [label="' + edge.label + '"];\n';
        }

        dot += '}';
        return dot;
    }

}

export interface IAutomataGraphProps extends React.ClassAttributes<AutomataGraph> {
    automata: IAutomata,
    currentNode?: number,
    flattenEdges?: boolean,
}
