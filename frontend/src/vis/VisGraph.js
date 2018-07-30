import React, { Component } from 'react';
import Graph from 'react-graph-vis';

export default class VisGraph extends Component {

    render() {
        var graph = {
            nodes: [
                { id: 1, label: 'Node 1' },
                { id: 2, label: 'Node 2' },
                { id: 3, label: 'Node 3' },
                { id: 4, label: 'Node 4' },
                { id: 5, label: 'Node 5' }
            ],
            edges: [
                { from: 1, to: 2 },
                { from: 1, to: 3 },
                { from: 1, to: 4 },
                { from: 1, to: 5 },
            ]
        };

        var options = {
            layout: {
                // improvedLayout: true,
                hierarchical: {
                    enable: true,
                    direction: 'LR'
                }
            },
            nodes: {
                shape: 'circle'
            },
            edges: {
                smooth: {
                    enable: false,
                }
            },
        }

        return (
            <Graph graph={graph} options={options} style={{ height: '500px' }} />
        );
    }
}