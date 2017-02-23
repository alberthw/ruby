import React from 'react';
import {Input, Row, Col, Button} from "antd";
import $ from "jquery";

export default class RepositorySetting extends React.Component {

    constructor(props) {
        super(props);

        var result = this.GetRepositorySetting();

        this.state = {
            Id: result.Id,
            Remoteserver: result.Remoteserver,
            Isconnected: result.Isconnected
        };

        this.handleRemoteServerIpChange = this
            .handleRemoteServerIpChange
            .bind(this);
        this.handleIsConnectedChange = this
            .handleIsConnectedChange
            .bind(this);
        this.handleCheckClick = this
            .handleCheckClick
            .bind(this);

    }

    TestRemoteServerConnection(url) {
        var result = false;
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            success: function (data) {
                console.log(data);
                if (data == "200") {
                    result = true;
                }
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    GetRepositorySetting() {
        var result = null;
        $.ajax({
            url: "/reposetting", dataType: "json",
            //       cache: false,
            async: false,
            success: function (data) {
                result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    UpdateRepositorySetting(id, ip, status) {
        var result = false;
        $.ajax({
            url: "/reposetting",
            dataType: "json",
            type: "POST",
            async: false,
            data: {
                Id: id,
                Remoteserver: ip,
                Isconnected: status
            },
            success: function (data) {
                //          console.log(data);
                result = true;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });

        return result;
    }

    handleRemoteServerIpChange(e) {
        this.setState({Remoteserver: e.target.value});
    }

    handleIsConnectedChange(e) {
        this.setState({Isconnected: e.target.value});
    }

    handleCheckClick(e) {
        this.UpdateRepositorySetting(this.state.Id, this.state.Remoteserver, isConnected);
        var url = "/testremoteserver";
        var isConnected = this.TestRemoteServerConnection(url);
        console.log("server connection :", isConnected);
        this.setState({Isconnected: isConnected});
        this.UpdateRepositorySetting(this.state.Id, this.state.Remoteserver, isConnected);
    }

    render() {
        const remoteServer = this.state.Remoteserver;
        const isConnected = this.state.Isconnected;
        return (
            <div>
                <Row>
                    <Col span={8}>
                        <Input
                            placeholder="COM"
                            addonBefore="Remote File Server:"
                            value={remoteServer}
                            onChange={this.handleRemoteServerIpChange}/>
                    </Col>
                    <Col span={8}>
                        <Input
                            addonBefore="Connected:"
                            value={isConnected}
                            onChange={this.handleIsConnectedChange}
                            disabled/>
                    </Col>
                    <Col span={8}>
                        <Button onClick={this.handleCheckClick}>Check</Button>
                    </Col>
                </Row>
            </div>
        );
    }

}