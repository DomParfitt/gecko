import * as React from 'react';

class TextInput extends React.Component<ITextInputProps, any> {

    constructor(props: ITextInputProps) {
        super(props);
    }

    public render(): JSX.Element {
        return (
            <div>
                {this.renderInput()}
                {this.props.hideButton ? null : this.renderButton()}
            </div>
        );
    }

    private renderInput(): JSX.Element {
        return(
            <input type="text" placeholder={this.props.placeholder} onChange={this.props.onChangeHandler} />
        );
    }

    private renderButton(): JSX.Element {
        return(
            <button onClick={this.props.onClickHandler}>{this.props.buttonText || "Enter"} </button>
        );
    }
}

export interface ITextInputProps extends React.ClassAttributes<TextInput> {
    placeholder?: string;
    buttonText?: string;
    hideButton?: boolean
    onChangeHandler?: (event: React.ChangeEvent<HTMLInputElement>) => void;
    onClickHandler?: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

export default TextInput;