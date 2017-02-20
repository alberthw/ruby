import React from 'react';
import {Input, Icon} from 'antd';

export default class EditableCell extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            value: this.props.value,
            editable: false
        };

        this.handleChange = this
            .handleChange
            .bind(this);
    }

    handleChange = (e) => {
        this.setState({value: e.target.value});
    }

    check = () => {
        this.setState({editable: false});
        if (this.props.onChange) {
            this
                .props
                .onChange(this.state.value);
        }
    }

    edit = () => {
        this.setState({editable: true});
    }
    render() {
        const {value, editable} = this.state;

        return (
            <div className="editable-cell">{editable
                    ? <div className="editable-cell-input-wrapper">
                            <Input value={value} onChange={this.handleChange} onPressEnter={this.check}/>
                            <Icon type="check" className="editable-cell-icon-check" onClick={this.check}/>

                        </div>
                    : <div className="editable-cell-text-wrapper">
                        {value || ' '}
                        <Icon type="edit" className="editable-cell-icon" onClick={this.edit}/>
                    </div>}
            </div>
        );
    }
}