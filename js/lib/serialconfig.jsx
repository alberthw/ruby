import React from "react";
import {Input, Row, Col, Button} from "antd";
import $ from "jquery";

export default class SerialConfig extends React.Component {
    constructor(props) {
        super(props);
        this.handleSerialNameChange = this
            .handleSerialNameChange
            .bind(this);
        this.handleStatusChange = this
            .handleStatusChange
            .bind(this);
        this.handleConnectClick = this
            .handleConnectClick
            .bind(this);

        this.state = {
            serialName: "",
            isConnected: false,
            id: 0
        };
    }

    handleSerialNameChange(e) {
        this.setState({serialName: e.target.value});
    }

    handleStatusChange(e) {
        this.setState({isConnected: e.target.value});
    }
    closeSerial() {
        var result = false;
        var url = "closeserial";
        $.ajax({
            url: url,
            dataType: "json",
            async: false,
            success: function (data) {
                result = true;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    handleConnectClick(e) {
        this.closeSerial();
        const data = this.state;
        $.ajax({
            url: this.props.url,
            dataType: "json",
            type: "POST",
            async: false,
            data: data,
            success: function (data) {
                //      result = true;
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }

    componentDidMount() {
        this.getSerialConfig(this.props.url);
        this.timer = setInterval(() => {
            this.getLatestSettings(this.props.url)
        }, 10000);
    }

    componentWillUnmount() {
        clearInterval(this.timer);
    }

    getSerialConfig(url) {

        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            success: function (data) {
                console.log(data);
                this.setState({serialName: data.SerialName, id: data.ID})
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }

    getLatestSettings(url) {
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            success: function (data) {
                console.log(data);
                this.setState({isConnected: data.IsConnected, id: data.ID})
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });

    }

    render() {
        const serialName = this.state.serialName;
        const isConnected = this.state.isConnected;
        return (
            <div>
                <Row>
                    <Col span={8}>
                        <Input
                            placeholder="COM"
                            addonBefore="Serial Name:"
                            value={serialName}
                            onChange={this.handleSerialNameChange}/>
                    </Col>
                    <Col span={8}>
                        <Input
                            addonBefore="Status:"
                            value={isConnected}
                            onChange={this.handleStatusChange}
                            disabled/>
                    </Col>
                    <Col span={8}>
                        <Button onClick={this.handleConnectClick}>Check</Button>
                    </Col>
                </Row>
            </div>
        );

    }
}