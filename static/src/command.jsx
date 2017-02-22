import React from 'react';
import {Card, Input, Row, Col, Button} from 'antd';
import $ from "jquery";

export default class Command extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            command: "",
            response: ""
        };

        this.handleCommandChange = this
            .handleCommandChange
            .bind(this);

        this.handleSendButtonClick = this
            .handleSendButtonClick
            .bind(this);

        this.handleClearButtonClick = this
            .handleClearButtonClick
            .bind(this);
    }

    handleCommandChange(e) {
        this.setState({command: e.target.value});
    }

    handleClearButtonClick(e) {
        this.setState({response: ""});
    }

    handleSendButtonClick(e) {
        //       this.handleClearButtonClick();
        var url = "/sendcommand";
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: {
                "command": this.state.command
            },
            success: function (data) {
                console.log("response :", data);
                this.setState({response: data});

            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }

    render() {
        return (
            <Card title="Command">
                <Row>
                    <Col span={8}>
                        <Input
                            addonBefore="Command:"
                            value={this.state.command}
                            onChange={this.handleCommandChange}
                            onPressEnter={this.handleSendButtonClick}/>
                    </Col>
                    <Col span={8}>
                        <Button onClick={this.handleSendButtonClick}>Send</Button>
                        <Button onClick={this.handleClearButtonClick}>Clear</Button>
                    </Col>
                </Row>
                <Row>
                    <Input
                        type="textarea"
                        value={this.state.response}
                        autosize={{
                        minRows: 2,
                        maxRows: 50
                    }}
                        readOnly/>
                </Row>
            </Card>
        );
    }
}