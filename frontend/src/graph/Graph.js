import React, {Component} from 'react';
import { select } from 'd3-selection';
import {} from 'd3-graphviz';

class Graph extends Component {

    constructor(props) {
        super(props);
    }

    render() {
        console.log("child rendered");
        return(
            <div id="graphDiv" ref="graphDiv"></div>
        );
    }

    loadGraph() {
        select('#graphDiv')
            .graphviz()
            .height(500)
            .width(500)
            .fit(true)
            .zoom(false)
            .renderDot(this.generateDot());
    }

    componentDidMount() {
        console.log("child mounted");
        this.loadGraph();
    }

    componentDidUpdate() {
        console.log("child updated");
        this.loadGraph();
    }

    generateDot() {
        console.log(this.props);
        let dot = 'digraph { rankdir="LR";\n';

        for(let i = 0; i < this.props.nodes.length; i++) {
            const node = this.props.nodes[i];
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

        for(let i = 0; i < this.props.edges.length; i++) {
            const edge = this.props.edges[i];
            dot += edge.from + '->' + edge.to + ' [label="' + edge.label + '"];\n';
        }

        dot += '}';
        return dot;
    }
}

export default Graph;