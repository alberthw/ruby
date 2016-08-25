function loadConfigFromServer(url) {
    var result = null;
    $.ajax({
        url: url,
        dataType: "json",
        cache: false,
        async: false,
        success: function (data) {
            result = data;
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

function serialControl(url, data){
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
        if (this.props.Isconnected == false){
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
        return loadConfigFromServer(this.props.url);
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
    /*  
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
    */
    handleUpdateClick: function (e) {
        e.preventDefault();
        postConfigToServer(this.props.url, this.state);
        this.getInitialState();
    },
    handleConnectClick:function(){
        var data = {
            Id:this.state.Id,
            Serialname:this.state.Serialname.trim(),
            Serialbaud :this.state.Serialbaud
        };
        var result = serialControl("openserial", this.state);
        if (result == true){
            this.setState({
                Isconnected:true
            });
        }
    },
    handleDisconnectClick:function(){
        var data = {
            Id:this.state.Id
        };
        var result = serialControl("closeserial", data);
        if (result == true){
            this.setState({
                Isconnected:false
            });
        }
    },
    render: function () {
        return (
            <div>
                <span>Serial Name: </span> <input type="text" value={this.state.Serialname} ref="nameInput" onChange={this.handleSerialNameChange}/>
                <span>Serial Baud: </span> <input type="text" value={this.state.Serialbaud} ref="baudInput" onChange={this.handleSerialBaudChange}/>
                <span>Connect: </span> <input type="text" value={this.state.Isconnected} ref="connectInput" onChange={this.handleConnectChange} readOnly/> 

                <input type="button" value="Update" onClick={this.handleUpdateClick}></input>
                <input type="button" value="Connect" onClick={this.handleConnectClick}></input>
                <input type="button" value="Disconnect"  onClick={this.handleDisconnectClick}></input>

                <SerialControl Isconnected={this.state.Isconnected} />
            </div>
        )

    }
});

ReactDOM.render(<SerialBox url="/config" />, document.getElementById("rubyconfig"));
