import React, {Component} from 'react';
import { select } from 'd3-selection';
import {} from 'd3-graphviz';

export default class Graph extends Component {

    componentDidMount() {
        select('#graphDiv')
            .graphviz()
            .height(500)
            .width(500)
            .fit(true)
            .zoom(false)
            .renderDot(
                `digraph { rankdir="LR"; 
                    0 -> 1 [label="a"]; 
                    0 -> 2 [label="b"]; 
                    1 -> 0 [label="c"];
                }`
            );
    }

    render() {

        return(
            <div id="graphDiv" ref="graphDiv"></div>
        );
    }
}