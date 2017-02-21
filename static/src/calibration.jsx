import React from 'react';
import {
    Card,
    Row,
    Col,
    Button,
    InputNumber,
    message
} from 'antd';
import $ from "jquery";

class CalibrationTable extends React.Component {
    render() {
        return (
            <Table title={() => "Calibration Data"}></Table>
        );
    }
}

export default class Calibration extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            rmsValue: 0
        }

        this.handleRMSInputChange = this
            .handleRMSInputChange
            .bind(this);

        this.handleStartCalibration = this
            .handleStartCalibration
            .bind(this);

        this.handlePrintCalibration = this
            .handlePrintCalibration
            .bind(this);

        this.handleSetCalibration = this
            .handleSetCalibration
            .bind(this);
    }

    handleStartCalibration(e) {
        var url = "/startcalibration";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                console.log("handleStartCalibration:", data);
                result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    handlePrintCalibration(e) {
        var url = "/printcalibration";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                console.log("handlePrintCalibration:", data);
                result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    handleSetCalibration(e) {
        var url = "/setcalibration";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data : {
                "rms" : this.state.rmsValue
            },
            success: function (data) {
                console.log("handleSetCalibration:", data);
                result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    handleRMSInputChange(e) {
        console.log(e);
        const number = parseInt(e || 0, 10);
        if (isNaN(number)) {
            return;
        }
        this.setState({rmsValue: e});
    }
    render() {

        return (
            <Card title="Calibration Data">
                <Row>
                    <b>Start to calibrate the data :
                    </b>
                    <Button onClick={this.handleStartCalibration}>Start</Button>
                </Row>
                <Row>
                    <b>Print the calibration data :
                    </b>
                    <Button onClick={this.handlePrintCalibration}>Print</Button>
                </Row>
                <Row>
                    <b>Enter RMS value :
                    </b>
                    <InputNumber onChange={this.handleRMSInputChange} value={this.state.rmsValue}></InputNumber>
                    <Button onClick={this.handleSetCalibration}>SetRMS</Button>
                </Row>

            </Card>
        );
    }
}
