import { graphviz } from 'd3-graphviz';
import * as React from 'react';
import { IEdge, INode } from 'src/automata/Automata';

class Graph extends React.Component<IGraphProps, any> {

    public render() {
        return(
            <div id="graphDiv" />
        );
    }

    public componentDidMount() {
        this.loadGraph();
    }
    
    
    public componentDidUpdate() {
        this.loadGraph();
    }

    private loadGraph() {
        graphviz('#graphDiv')
            // .graphviz({useWorker: false})
            .height(500)
            .width(500)
            .fit(true)
            .zoom(false)
            .renderDot(this.generateDot());
    }

    private generateDot() {
        let dot = 'digraph { rankdir="LR";\n';

        for(const node of this.props.nodes) {
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

        for(const edge of this.props.edges) {
            dot += edge.from + '->' + edge.to + ' [label="' + edge.label + '"];\n';
        }

        dot += '}';
        return dot;
    }
}

export interface IGraphProps extends React.ClassAttributes<Graph> {
    nodes: INode[]
    edges: IEdge[]
    currentNode: number
}

export default Graph;