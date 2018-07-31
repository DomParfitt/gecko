import React, {Component} from 'react';
import { select } from 'd3-selection';
import {} from 'd3-graphviz';

export default class Graph extends Component {

    componentDidMount() {
        select('#graphDiv')
            .graphviz()
            .renderDot(
                'digraph { rankdir="LR"; a -> b; a -> c; b -> a}'
            );
    }

    render() {

        return(
            <div id="graphDiv" ref="graphDiv"></div>
        );
    }
}