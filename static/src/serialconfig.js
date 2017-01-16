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

class SerialConfig extends React.Component {
    constructor(props) {
        super(props);
        this.handleSerialNameChange = this.handleSerialNameChange.bind(this);
        this.handleStatusChange = this.handleStatusChange.bind(this);
        this.handleConnectClick = this.handleConnectClick.bind(this);

        this.state = {
            Serialname: "",
            Isconnected: false,
            Id: 0
        };
    }

    handleSerialNameChange(e) {
        this.setState({
            Serialname: e.target.value
        });
    }

    handleStatusChange(e) {
        this.setState({
            Isconnected: e.target.value
        });
    }

    handleConnectClick(e) {
        closeSerial();
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
        this.timer = setInterval(() => { this.getLatestSettings(this.props.url) }, 10000);
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
                this.setState({
                    Serialname: data.Serialname,
                    Id: data.Id
                })
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
                this.setState({
                    Isconnected: data.Isconnected,
                    Id: data.Id
                })
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });

    }

    render() {
        const serialName = this.state.Serialname;
        const isConnected = this.state.Isconnected;
        return (
            <div className="input-group">
                <span className="input-group-addon">Serial Name:</span>
                <input type="text" className="form-control" placeholder="com1" value={serialName} onChange={this.handleSerialNameChange}></input>
                <span className="input-group-addon">Status:</span>
                <input type="text" className="form-control" value={isConnected} onChange={this.handleStatusChange} readOnly></input>
                <span className="input-group-btn">
                    <button type="button" className="btn btn-default" onClick={this.handleConnectClick}>Connect</button>
                </span>
            </div>
        );

    }
}

ReactDOM.render(
    <SerialConfig url="/config" />,
    document.getElementById("serialConfig")
);