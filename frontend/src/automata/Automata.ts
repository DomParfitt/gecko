import { IAutomata } from './Automata';
// Interface for Automata objects used by the frontend
export interface IAutomata {
    currentNode: number,
    nodes: INode[],
    edges: IEdge[]
}

export function isAutomata(object: any): object is IAutomata {
    const automata = object as IAutomata;
    return automata.currentNode !== undefined 
            && automata.nodes !== undefined 
            && automata.edges !== undefined;
}

// Interface for Edge objects used by the frontend's
// internal Automata representation
export interface IEdge {
    from: number,
    to: number,
    label: string
}

// Interface for Node objects used by the frontend's
// internal Automata representation
export interface INode {
    id: number,
    isTerminal: boolean
}