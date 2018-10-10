import { IEdge } from './Edge';
import { INode } from './Node';

export interface IAutomata {
    currentNode: number,
    nodes: INode[],
    edges: IEdge[]
}