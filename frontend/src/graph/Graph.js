import React, {Component} from 'react';
import { select } from 'd3-selection';
import {} from 'd3-graphviz';

class Graph extends Component {

    render() {
        return(
            <div id="graphDiv" ref="graphDiv"></div>
        );
    }

    loadGraph() {
        select('#graphDiv')
            .graphviz({useWorker: false})
            .height(500)
            .width(500)
            .fit(true)
            .zoom(false)
            .renderDot(this.generateDot());
    }

    componentDidMount() {
        this.loadGraph();
    }

    componentDidUpdate() {
        this.loadGraph();
    }

    generateDot() {
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

export default Graph;