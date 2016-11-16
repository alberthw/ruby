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
        this.handleDeviceNameChange = this.handleDeviceNameChange.bind(this);

        this.state = {
            Serialname: "",
            Isconnected: false,
            Devicename: "",
            Id: 0
        };
    }

    handleSerialNameChange(e) {
        this.setState({
            Serialname: e.target.value
        });
    }

    handleDeviceNameChange(e) {
        this.setState({
            Devicename: e.target.value
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
        this.timer = setInterval(() => { this.getLatestSettings(this.props.url) }, 1000);
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
                    Devicename: data.Devicename,
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
        const deviceName = this.state.Devicename;
        return (
            <div>
                <a>Serial Name:</a>
                <input type="text" placeholder="com1" value={serialName} onChange={this.handleSerialNameChange}></input>
                <a>Status:</a>
                <input type="text" value={isConnected} onChange={this.handleStatusChange} readOnly></input>

                <input type="button" value="Connect" onClick={this.handleConnectClick}></input>
                <br />
                <a>Device Name:</a>
                <input type="text" value={deviceName} onChange={this.handleDeviceNameChange} readOnly></input>
            </div>
        );

    }
}

ReactDOM.render(
    <SerialConfig url="/config" />,
    document.getElementById("serialConfig")
);