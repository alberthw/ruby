function SendCommand(command) {
    var url = "/command";
    $.ajax({
        url: url,
        dataType: "json",
        type: "POST",
        cache: false,
        async: false,
        data: {
            "command": command
        },
        success: function (data) {
            console.log(data);
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });
}

function string2Bin(str) {
    var result = [];
    for (var i = 0; i < str.length; i++) {
        result.push(str.charCodeAt(i));
    }
    return result;
}


class CommandButton extends React.Component {
    constructor(props) {
        super(props);

        this.handleButtonClick = this.handleButtonClick.bind(this);
    }

    handleButtonClick(e) {
        SendCommand(this.props.command);
    }

    render() {
        return (
            <button type="button" value={this.props.command} className="btn btn-default" onClick={this.handleButtonClick}>{this.props.command}</button>
        );
    }
}

class SerialCommand extends React.Component {
    constructor(props) {
        super(props);

        this.handleInputChange = this.handleInputChange.bind(this);
        this.handleSendButtonClick = this.handleSendButtonClick.bind(this);
        this.handleOutputChange = this.handleOutputChange.bind(this);

        this.handleSetSysConfigButtonClick = this.handleSetSysConfigButtonClick.bind(this);

        this.state = {
            input: "",
            output: ""
        };
    }

    handleInputChange(e) {
        this.setState({
            input: e.target.value
        });
    }


    handleSendButtonClick(e) {
        SendCommand(this.state.input);
    }


    handleSetSysConfigButtonClick(e) {
        var url = "/setsysconfig";
        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                console.log(data);
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
    }

    handleSetHwConfigButtonClick(e) {
        var url = "/sethwconfig";
        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                console.log(data);
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
    }

    componentDidMount() {
        this.timer = setInterval(() => { this.getSerialOutput() }, 500);
    }

    componentWillUnmount() {
        clearInterval(this.timer);
    }


    getSerialOutput() {

        document.getElementById("taOutput").scrollTop = document.getElementById("taOutput").scrollHeight;
        var url = "/getreceivecommands";
        //       $("#taOutput").scrollTop = $("#taOutput").scrollHeight;
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: {
                "limit": 500
            },
            success: function (data) {
                let result = "";
                data.forEach(function (row) {
                    result += row.Info;
                });
                this.setState({
                    output: result
                })
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });


    }

    handleOutputChange(e) {
        document.getElementById("taOutput").scrollTop = document.getElementById("taOutput").scrollHeight;
        this.setState({
            output: e.target.value
        });
    }


    render() {

        return (
            <div>
                <div className="input-group  col-md-8">
                    <span className="input-group-addon">Command:</span>
                    <input type="text" className="form-control" value={this.state.input} onChange={this.handleInputChange}></input>
                    <span className="input-group-btn">
                        <button type="button" className="btn btn-default" onClick={this.handleSendButtonClick}>Send</button>
                        <CommandButton command="ver.get"/>
                        <CommandButton command="commands" />
                    </span>
                </div>
                <div>
                    <textarea id="taOutput" cols="100" rows="30" value={this.state.output} onChange={this.handleOutputChange} readOnly></textarea>
                </div>
            </div>
        );
    }

}


ReactDOM.render(
    <SerialCommand />,
    document.getElementById("command"));