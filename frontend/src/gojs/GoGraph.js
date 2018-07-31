import React, { Component } from 'react';
import { GojsDiagram } from 'react-gojs';
import * as go from 'gojs'
import { ToolManager, Diagram } from 'gojs';
import './MyDiagram.css';

export default class GoGraph extends Component {

    constructor(props) {
        super(props);
        this.createDiagram = this.createDiagram.bind(this);
        this.state = {
            model: {
                nodeDataArray: [
                    { key: '1' },
                    { key: '2' },
                    { key: '3' },
                    { key: '4' },
                ],
                linkDataArray: [
                    { from: '1', to: '2' },
                    { from: '3', to: '2' },
                    { from: '2', to: '4' },
                    { from: '3', to: '4' },
                ]
            }
        };
    }

    render() {

        return (
            <GojsDiagram
                key="gojsDiagram"
                diagramId="myDiagramDiv"
                model={this.state.model}
                createDiagram={this.createDiagram}
                className="myDiagram"
            />
        );
    }

    createDiagram(diagramId) {
        const $ = go.GraphObject.make;

        const myDiagram = $(go.Diagram, diagramId, {
            initialContentAlignment: go.Spot.LeftCenter
        });

        myDiagram.nodeTemplate = $(
            go.Node,
            'Auto',
            $(go.Shape, 'Circle', { strokeWidth: 0 }, new go.Binding('fill', 'color')),
            $(go.TextBlock, { margin: 8 }, new go.Binding('text', 'key'))
        );

        return myDiagram;
    }
}