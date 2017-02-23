import React from 'react';
import {Row, Col, Card, Button, Input} from 'antd';
import $ from "jquery";

export default class ValidateConfig extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            isConfigValidated: false
        }

        this.handleValidateButtonClick = this
            .handleValidateButtonClick
            .bind(this);
        this.handleUpdateButtonClick = this
            .handleUpdateButtonClick
            .bind(this);
    }
    handleValidateButtonClick(e) {
        var url = "/validateconfig";
        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                console.log(data);
                //result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
    }

    handleUpdateButtonClick(e) {
        this.getVersions();
        this.getLastKnownVersions();
        window
            .location
            .reload();
    }

    getConfigValidateStatus() {
        var url = "/config";
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            success: function (data) {
                console.log("config:", data);
                this.setState({isConfigValidated: data.IsConfigValidated})
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }

    getVersions() {
        var url = "/getversion";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                //            console.log(data);
                result = data;

            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    getLastKnownVersions() {
        var url = "/getlastknownversion";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                //            console.log(data);
                result = data;

            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    componentWillMount() {
        this.getConfigValidateStatus();
    }

    render() {
        return (
            <Card title="Configuration Validation">
                <Row>
                    <Col span={2}>
                        <Input
                            addonBefore="Validate Result:"
                            value={this.state.isConfigValidated.toString()}
                            onChange={this.handleIsValidatedChange}
                            disabled/>
                    </Col>
                    <Col span={2}>
                        <Button onClick={this.handleValidateButtonClick}>Validate</Button>
                    </Col>
                    <Col span={2}>
                        <Button onClick={this.handleUpdateButtonClick}>Update</Button>
                    </Col>
                </Row>
            </Card>
        );
    }
}