import { graphviz, GraphvizOptions } from 'd3-graphviz';
import * as React from 'react';

class Graphviz extends React.Component<IGraphvizProps, any> {

    private static count = 0;
    private static defaultOptions: GraphvizOptions = {
        fit: true,
        height: 500,
        width: 500,
        zoom: false,
    };

    private id: string;

    constructor(props: IGraphvizProps) {
        super(props);
        this.id = "graphviz" + Graphviz.count;
        Graphviz.count++;
    }

    public render(): JSX.Element {
        return (
            <div id={this.id} />
        );
    }

    public componentDidMount() {
        this.loadGraph();
    }

    public componentDidUpdate() {
        this.loadGraph();
    }

    private loadGraph() {
        graphviz('#' + this.id)
            .options(this.props.options || Graphviz.defaultOptions)
            .renderDot(this.props.dot);
    }

}

export interface IGraphvizProps extends React.ClassAttributes<Graphviz> {
    dot: string,
    options?: GraphvizOptions,
}

export default Graphviz;