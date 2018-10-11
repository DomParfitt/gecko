import * as React from 'react';

class Switcher extends React.Component<ISwitcherProps, any> {

    constructor(props: any) {
        super(props);
    }

    public render(): JSX.Element {
        const list = this.props.graphs.map((graph, index) => <button key={index} onClick={graph.selector}>{graph.name}</button>);
        return (
            <div>
                {list}
            </div>
        );
    }

}

export interface ISwitcherProps extends React.ClassAttributes<Switcher> {
    graphs: ISwitcherData[];
}

export interface ISwitcherData {
    name: string;
    selector: () => void;
}


export default Switcher;