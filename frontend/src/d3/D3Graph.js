import React, { Component } from 'react';
import { Graph } from 'react-d3-graph';

export default class D3Graph extends Component {
    render() {
        const data = {
            nodes: [
              {id: 1},
              {id: 2},
              {id: 3}
            ],
            links: [
                {source: 1, target: 2},
                {source: 2, target: 3},
                {source: 3, target: 2},
            ]
        };
        return(
            <Graph id='d3-graph' data={data} />
        );
    }
}
