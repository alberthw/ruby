import React from 'react';
import {
    Card,
    Input,
    Row,
    Col,
    Button,
    InputNumber
} from 'antd';
import $ from "jquery";


class SetCalibration extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            rmsValue: 0
        };

        this.handleSetCalibration = this
            .handleSetCalibration
            .bind(this);

        this.handleRMSInputChange = this
            .handleRMSInputChange
            .bind(this);
    }
    handleRMSInputChange(e) {
        console.log(e);
        const number = parseInt(e || 0, 10);
        if (isNaN(number)) {
            return;
        }
        this.setState({rmsValue: e});
    }

    handleSetCalibration(e) {
        var url = "/setcalibration";
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: {
                "rms": this.state.rmsValue
            },
            success: function (data) {
                console.log("handleSetCalibration:", data);
                this
                    .props
                    .onResponse(data);
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }
    render() {
        return (
            <Card title="Set the calibration data">
                <b>Enter RMS value :</b>
                <InputNumber onChange={this.handleRMSInputChange} value={this.state.rmsValue}></InputNumber>
                <Button onClick={this.handleSetCalibration}>SetRMS</Button>
            </Card>
        );
    }
}


class CommandCard extends React.Component {
    constructor(props) {
        super(props);

        this.handleButtonClick = this
            .handleButtonClick
            .bind(this);
    }

    handleButtonClick(e) {
        let url = this.props.URL;
        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                console.log(this.props.Title + ":" + data);
                this
                    .props
                    .onResponse(data);
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }

    render() {
        return (
            <Card title={this.props.Title}>
                <Button onClick={this.handleButtonClick}>{this.props.ButtonText}</Button>
            </Card>
        );
    }
}

export default class Calibration extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            result: ""
        }

        this.handleResponse = this
            .handleResponse
            .bind(this);
    }

    handleResponse(e) {
        console.log("handleResponse:" + e);
        this.setState({result: e});
    }
    render() {

        return (
            <Card title="Calibration Data">
                <Row gutter={4}>
                    <Col span={4}>
                        <Row>
                            <Col>
                                <CommandCard
                                    URL="/enterservicemode"
                                    Title="Enter the service mode"
                                    ButtonText="Enter"
                                    onResponse={this.handleResponse}/>
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <CommandCard
                                    URL="/exitservicemode"
                                    Title="Exit the service mode"
                                    ButtonText="Exit"
                                    onResponse={this.handleResponse}/>
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <CommandCard
                                    URL="/startcalibration"
                                    Title="Start to calibrate the data"
                                    ButtonText="Start"
                                    onResponse={this.handleResponse}/>
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <SetCalibration onResponse={this.handleResponse}/>
                            </Col>
                        </Row>
                    </Col>
                    <Col span={20}>
                        <Card title="Response">
                            <Input
                                type="textarea"
                                value={this.state.result}
                                autosize={{
                                minRows: 20,
                                maxRows: 200
                            }}
                                readOnly/>
                        </Card>
                    </Col>
                </Row>

            </Card>
        );
    }
}
