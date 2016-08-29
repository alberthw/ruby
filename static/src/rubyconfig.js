function closeSerial() {
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

function serialControl(url, data) {
    var result = false;
    $.ajax({
        url: url,
        dataType: "json",
        type: "POST",
        async: false,
        data: data,
        success: function (data) {
            //               alert(data);
            result = true;
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });
    return result;
}

var SendCommands = React.createClass({
    handleClick: function () {
        if (this.props.Isconnected == false) {
            return;
        }
        var url = this.props.url;
        var data = {
            command: this.props.data
        }
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            async: false,
            data: data,
            success: function (data) {
                //          alert(data);
                //        console.log(data);
            },
            error: function (xhr, status, err) {
                alert(status);
                console.error(url, status, err.toString());
            }
        });
    },
    render: function () {
        return (
            <input type="button" value={this.props.data} onClick={this.handleClick}></input>
        )
    }
});

var CommandBox = React.createClass({
    render: function () {
        return (
            <div className="commandBox">
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="commands"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="ver.get"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="service.mode"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="service.menu"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="submode.exit"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="viadc.print"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="top.on"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="top.off"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="led.off"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="beep.on"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="beep.off"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="image.upload"/>
                <SendCommands Isconnected={this.props.Isconnected} url={this.props.url} data="image.update"/>
            </div>
        );
    }
});

var SerialControl = React.createClass({
    render: function () {
        return (
            <div className="serialControl">

                <CommandBox url="command" Isconnected={this.props.Isconnected}/>

            </div>
        )
    }
});

var SerialBox = React.createClass({
    getInitialState: function () {
        return {
            Serialname: "",
            Serialbaud:"",
            Isconnected:false
        };
    },
    componentDidMount: function () {
        this.loadConfigFromServer();
        setInterval(this.loadStatusFromServer, this.props.pollInterval);
    },
    loadConfigFromServer: function () {
        $.ajax({
            url: this.props.url,
            dataType: "json",
            cache: false,
            async: false,
            success: function (data) {
                this.setState({
                    Serialname: data.Serialname,
                    Serialbaud:data.Serialbaud,
                    Isconnected:data.Isconnected,
                    Id:data.Id
                })
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    loadStatusFromServer: function () {
        $.ajax({
            url: this.props.url,
            dataType: "json",
            cache: false,
            async: false,
            success: function (data) {
                this.setState({
                    Isconnected: data.Isconnected
                });
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    postConfigToServer(data) {
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

    },
    handleSerialNameChange: function (e) {
        this.setState({
            Serialname : e.target.value
        });
    },
    handleSerialBaudChange: function (e) {
        this.setState({
            Serialbaud: e.target.value
        });
    },
    handleConnectChange: function (e) {
        this.setState({
            Isconnected: e.target.value
        });
    },
    handleUpdateClick: function (e) {
        this.postConfigToServer(this.state);
        closeSerial();
    },
    render: function () {
        return (
            <div>
                <span>Serial Name: </span> <input type="text" value={this.state.Serialname} onChange={this.handleSerialNameChange}/>
                <span>Serial Baud: </span> <input type="text" value={this.state.Serialbaud} onChange={this.handleSerialBaudChange}/>
                <span>Connect: </span> <input type="text" value={this.state.Isconnected} onChange={this.handleConnectChange} readOnly/>

                <input type="button" value="Update" onClick={this.handleUpdateClick}></input>

                <SerialControl Isconnected={this.state.Isconnected} />
                

            </div>
        )

    }
});

ReactDOM.render(<SerialBox url="/config" pollInterval={1000}/>, document.getElementById("rubyconfig"));

