function loadConfigFromServer(url) {
    var result = {
        id: null,
        name: "",
        baud: null,
        connect: false
    };
    $.ajax({
        url: url,
        dataType: "json",
        cache: false,
        async: false,
        success: function (data) {
            result.id = data.Id;
            result.name = data.Serialname;
            result.baud = data.Serialbaud;
            result.connect = data.Isconnected;
        },
        error: function (xhr, status, err) {
            console.error(this.props.url, status, err.toString());
        }
    });
    return result;
}

function postConfigToServer(url, data) {
    var result = false;
    $.ajax({
        url: url,
        dataType: "json",
        type: "POST",
        async: false,
        data: data,
        success: function (data) {
            result = true;
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });
    return result;
}

var OpenSerial = React.createClass({
    handleClick: function () {
        var url = this.props.url;
        var data = {
            name: this.props.name,
            baud: this.props.baud,
            id: this.props.id,
            connect: true
        };
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            async: false,
            data: data,
            success: function (data) {
                alert(data);
                refresh();
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
       history.go(0);
    },
    render: function () {
        return (
            <input type="button" value="Connect" ref="OpenInput" onClick={this.handleClick}></input>
        )
    }
});

var CloseSerial = React.createClass({
    handleClick: function () {
        var url = this.props.url;
        var data = {
            id: this.props.id,
            connect: false
        };
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            async: false,
            data: data,
            success: function (data) {

                alert(data);
                //        console.log(data);
            },
            error: function (xhr, status, err) {
                alert(status);
                console.error(url, status, err.toString());
            }
        });
        history.go(0);
    },
    render: function () {
        return (
            <input type="button" value="Disconnect" onClick={this.handleClick}></input>
        )
    }
});

var SendCommands = React.createClass({
    handleClick: function () {
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
                alert(data);
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
                <SendCommands url={this.props.url} data="commands"/>
                <SendCommands url={this.props.url} data="ver.get"/>
                <SendCommands url={this.props.url} data="service.mode"/>
                <SendCommands url={this.props.url} data="service.menu"/>
                <SendCommands url={this.props.url} data="submode.exit"/>
                <SendCommands url={this.props.url} data="viadc.print"/>
                <SendCommands url={this.props.url} data="top.on"/>
                <SendCommands url={this.props.url} data="top.off"/>
                <SendCommands url={this.props.url} data="led.off"/>
                <SendCommands url={this.props.url} data="beep.on"/>
                <SendCommands url={this.props.url} data="beep.off"/>
                <SendCommands url={this.props.url} data="image.upload"/>
                <SendCommands url={this.props.url} data="image.update"/>
            </div>
        );
    }
});

var SerialControl = React.createClass({
    render: function () {
        return (
            <div className="serialControl">

                <CommandBox url="command" />

            </div>
        )
    }
});

var SerialItems = React.createClass({
    handleChange: function () {
        this.props.onSerialInput(this.refs.nameInput.value, this.refs.baudInput.value, this.refs.connectInput.value);
    },
    render: function () {
        return (
            <div>
                <span>Serial Name: </span> <input type="text" value={this.props.name} ref="nameInput" onChange={this.handleChange}/>
                <span>Serial Baud: </span> <input type="text" value={this.props.baud} ref="baudInput" onChange={this.handleChange}/>
                <span>Connect: </span> <input type="text" value={this.props.connect} ref="connectInput" onChange={this.handleChange} />
            </div>
        );
    }
})

var SerialLaunch = React.createClass({
    handleChange:function(){


    },
    render: function () {
        return (
            <div>
                
            </div>
        );

    }
});

var SerialBox = React.createClass({
    getInitialState: function () {
        return loadConfigFromServer(this.props.url);
    },
    /*
    handleSerialNameChange: function (e) {
        this.setState({
            name: e.target.value
        });
    },
    handleSerialBaudChange: function (e) {
        this.setState({
            baud: e.target.value
        });
    },
    handleSConnectChange: function (e) {
        this.setState({
            connect: e.target.value
        });
    },
    */
    handleSerialInput: function (name, baud, connect) {
        if (name != null) {
            this.setState({
                name: name
            });
        }
        if (baud != null) {
            this.setState({
                baud: baud
            });
        }
        if (connect != null) {
            this.setState({
                connect: connect
            });
        }
    },
    handleUpdateClick: function (e) {
        e.preventDefault();
        var d = this.state.id;
        var n = this.state.name.trim();
        var b = this.state.baud;
        var c = this.state.connect;
        postConfigToServer(this.props.url, { id: d, name: n, baud: b, connect: c });
        this.getInitialState();
    },
    render: function () {
        var status = this.state.connect.toString();
        return (
            <div>
                <SerialItems name={this.state.name} baud={this.state.baud} connect={this.state.connect} onSerialInput={this.handleSerialInput} />
                <input type="button" value="Update" onClick={this.handleUpdateClick}></input>
                <OpenSerial url="openserial" name={this.state.name} baud={this.state.baud} id={this.state.id} connect={this.state.connect} onSerialInput={this.handleSerialInput} />
                <CloseSerial url="closeserial" id={this.state.id} connect={this.state.connect} onSerialInput={this.handleSerialInput}/>


                <SerialControl name={this.state.name} baud={this.state.baud} />
            </div>
        )

    }
});

ReactDOM.render(<SerialBox url="/config" />, document.getElementById("rubyconfig"));
