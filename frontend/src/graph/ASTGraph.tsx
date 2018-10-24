import * as React from 'react';
import { IAbstractSyntaxTree } from 'src/ast/AbstractSyntaxTree';
import Graphviz from './Graphviz';

class ASTGraph extends React.Component<IASTGraphProps, any> {

    constructor(props: IASTGraphProps) {
        super(props);
    }

    public render(): JSX.Element {
        return(
            <Graphviz dot={this.generateDot()}/>
        );
    }

    private generateDot(): string {
        let dot = 'digraph {\n';
        dot += this.generateDotNodes(this.props.ast);
        dot += '}';
        // tslint:disable-next-line:no-console
        console.log(dot);
        return dot;
    }

    private generateDotNodes(ast: IAbstractSyntaxTree): string {
        return this.generateDotNodesHelper(ast, 0).dot;
    }

    private generateDotNodesHelper(ast: IAbstractSyntaxTree, count: number): {dot: string, count: number} {
        const rootName = count++;
        let dot = rootName +  ' [label="' + ast.label + '"';
        if (ast.children.length === 0) {
            dot +=  ' fillcolor="lightblue" style="filled"]\n';
            return {dot, count};
        } 

        dot += ']\n';
        let newCount = count;
        for (const child of ast.children) {
            const childName = newCount;
            const result = this.generateDotNodesHelper(child, newCount++);
            dot += result.dot;
            dot += rootName + ' -> ' + childName + '\n';
            newCount = result.count;
        }

        return {dot, count: newCount};
    }
}

export interface IASTGraphProps extends React.ClassAttributes<ASTGraph> {
    ast: IAbstractSyntaxTree
}

export default ASTGraph;