



class SerialCommand extends React.Component {
    constructor(props) {
        super(props);

        this.handleInputChange = this.handleInputChange.bind(this);
        this.handleSendButtonClick = this.handleSendButtonClick.bind(this);
        this.handleOutputChange = this.handleOutputChange.bind(this);

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

        $.ajax({
            url: "/command",
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: {
                "command": this.state.input
            },
            success: function (data) {
                console.log(data);
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
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
        //       $("#taOutput").scrollTop = $("#taOutput").scrollHeight;
        $.ajax({
            url: "/getreceivecommands",
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
                <div className="input-group  col-md-4">
                    <span className="input-group-addon">Command:</span>
                    <input type="text" className="form-control" value={this.state.input} onChange={this.handleInputChange}></input>
                    <span className="input-group-btn">
                        <button type="button" className="btn btn-default" onClick={this.handleSendButtonClick}>Send</button>
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