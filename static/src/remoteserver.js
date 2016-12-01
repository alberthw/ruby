function TestRemoteServerConnection(url) {
    var result = false;
    $.ajax({
        url: url,
        dataType: "json",
        //       cache: false,
        async: false,
        success: function (data) {
            if (data == "ok") {
                result = true;
            }
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });
    return result;
}

function GetRemoteServer() {
    var result = {
        Id: 0,
        Remoteserver: ""
    };
    $.ajax({
        url: "/remoteserver",
        dataType: "json",
        //       cache: false,
        async: false,
        success: function (data) {
            result.Id = data.Id;
            result.Remoteserver = data.Remoteserver;
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });
    return result;
}

function UpdateRemoteServer(ip, id) {
    var result = false;
    $.ajax({
        url: "/remoteserver",
        dataType: "json",
        type: "POST",
        async: false,
        data: {
            Id: id,
            Remoteserver: ip
        },
        success: function (data) {
            console.log(data);
            result = true;

        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });

    return result;


}

class RemoteFileServer extends React.Component {

    constructor(props) {
        super(props);

        var result = GetRemoteServer();

        this.state = {
            Id: result.Id,
            remoteServer: result.Remoteserver,
            isConnected: false
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

    handleRemoteServerIpChange(e) {

        this.setState({ remoteServer: e.target.value });
    }

    handleIsConnectedChange(e) {
        this.setState({ isConnected: e.target.value });
    }

    handleCheckClick(e) {
        const remoteServer = this.state.remoteServer;
        const id = this.state.Id;
        var url = "http://" + this.state.remoteServer + "/test";
        this.setState({
            isConnected: TestRemoteServerConnection(url)
        });
        UpdateRemoteServer(remoteServer, id);
    }

    render() {
        const remoteServer = this.state.remoteServer;
        const isConnected = this.state.isConnected;
        return (
            <div>
                <a>Remote File Server : </a>
                <input type="text" value={remoteServer} onChange={this.handleRemoteServerIpChange}></input>
                <a>Connected : </a>
                <input type="text" value={isConnected} onChange={this.handleIsConnectedChange} readOnly></input>
                <input type="button" value="Check" onClick={this.handleCheckClick}></input>
            </div>
        );
    }

}

ReactDOM.render(<RemoteFileServer />, document.getElementById("remoteserver"));