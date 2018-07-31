import React, {Component} from 'react';
import { select } from 'd3-selection';
import {} from 'd3-graphviz';

export default class Graph extends Component {

    constructor(props) {
        super(props);
        this.state = {
            currentNode: 0,
            nodes: [
                {id: 0, isTerminal: false},
                {id: 1, isTerminal: true},
            ],
            edges: [{from: 0, to: 1, label: 'a'}]
        };
    }

    componentDidMount() {
        select('#graphDiv')
            .graphviz()
            .height(500)
            .width(500)
            .fit(true)
            .zoom(false)
            .renderDot(this.generateDot());
    }

    render() {
        return(
            <div id="graphDiv" ref="graphDiv"></div>
        );
    }

    generateDot() {
        let dot = 'digraph { rankdir="LR";\n';

        for(let i = 0; i < this.state.nodes.length; i++) {
            const node = this.state.nodes[i];
            dot += node.id +' [';

            if (node.isTerminal) {
                dot += 'shape="doublecircle" ';
            } else {
                dot += 'shape="circle" ';
            }

            if (i === this.state.currentNode) {
                dot += 'fillcolor="red" style="filled" ';
            }

            dot += '];\n';
        }

        for(let i = 0; i < this.state.edges.length; i++) {
            const edge = this.state.edges[i];
            dot += edge.from + '->' + edge.to + ' [label="' + edge.label + '"];\n';
        }

        dot += '}';
        return dot;
    }
}