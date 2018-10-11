import { graphviz } from 'd3-graphviz';
import * as React from 'react';
import { IAutomata } from 'src/automata/Automata';

class AutomataGraph extends React.Component<IAutomataGraphProps, any> {

    public render(): JSX.Element {
        return(
            <div id="automataGraphDiv" />
        );
    }

    public componentDidMount() {
        this.loadGraph();
    }
    
    
    public componentDidUpdate() {
        this.loadGraph();
    }

    private loadGraph() {
        graphviz('#automataGraphDiv')
            // .graphviz({useWorker: false})
            .height(500)
            .width(500)
            .fit(true)
            .zoom(false)
            .renderDot(this.generateDot());
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

            if (node.id === this.props.automata.currentNode) {
                dot += 'fillcolor="red" style="filled" ';
            }

            dot += '];\n';
        }

        for(const edge of this.props.automata.edges) {
            dot += edge.from + '->' + edge.to + ' [label="' + edge.label + '"];\n';
        }

        dot += '}';
        return dot;
    }
}

export interface IAutomataGraphProps extends React.ClassAttributes<AutomataGraph> {
    automata: IAutomata
}

export default AutomataGraph;