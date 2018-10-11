import { graphviz } from 'd3-graphviz';
import * as React from 'react';
import { IAbstractSyntaxTree } from 'src/ast/AbstractSyntaxTree';

class ASTGraph extends React.Component<IASTGraphProps, any> {

    private nodeCount: number;

    constructor(props: IASTGraphProps) {
        super(props);
        this.nodeCount = 0;
    }

    public render(): JSX.Element {
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
    
    private generateDot(): string {
        let dot = 'digraph {\n';
        dot += this.generateDotNodes(this.props.ast);
        dot += '}';
        this.nodeCount = 0;
        // tslint:disable-next-line:no-console
        console.log(dot);
        return dot;
    }

    private generateDotNodes(ast: IAbstractSyntaxTree): string {
        const rootName = 'node' + this.nodeCount;
        let dot = rootName +  ' [label="' + ast.label + '"]\n';
        if (ast.children.length === 0) {
            return dot;
        } 

        for (const child of ast.children) {
            this.nodeCount++;
            const childName = 'node' + this.nodeCount;
            dot += this.generateDotNodes(child);
            dot += rootName + ' -> ' + childName + '\n';
        }

        return dot;
    }
}

export interface IASTGraphProps extends React.ClassAttributes<ASTGraph> {
    ast: IAbstractSyntaxTree
}

export default ASTGraph;