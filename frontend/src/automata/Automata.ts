// Interface for Automata objects used by the frontend
export interface IAutomata {
    nodes: INode[],
    edges: IEdge[],
}

export function isAutomata(object: any): object is IAutomata {
    const automata = object as IAutomata;
    return automata.nodes !== undefined
        && automata.edges !== undefined;
}

export function getFlattenedEdges(automata: IAutomata): IEdge[] {
    sortEdges(automata);
    const newEdges: IEdge[] = [];
    const checkedEdges: IEdge[] = [];

    for (const edge of automata.edges) {

        if (checkedEdges.indexOf(edge) >= 0) {
            continue;
        }

        const labels: string[] = []
        labels.push(edge.label);
        for (const otherEdge of automata.edges) {
            if (edge.from === otherEdge.from && edge.to === otherEdge.to && edge.label !== otherEdge.label) {
                checkedEdges.push(otherEdge);
                labels.push(otherEdge.label);
            }
        }
        let label = labels.toString();
        if (labels.length > 5) {
            label = labels[0] + ","+ labels[1] +"..." + labels[labels.length-1];
        }
        const newEdge: IEdge = { from: edge.from, to: edge.to, label };
        newEdges.push(newEdge);
    }

    return newEdges;
}

export function sortEdges(automata: IAutomata) {
    automata.edges.sort(compare);
}

function compare(first: IEdge, second: IEdge): number {
    if (first.from < second.from) {
        return -1;
    }

    if (first.to < second.to) {
        return -1;
    }

    if (first.label < second.label) {
        return -1;
    }

    return 1;
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