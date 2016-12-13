function TestRemoteServerConnection(url) {
    var result = false;
    $.ajax({
        url: url,
        dataType: "json",
        cache: false,
        async: false,
        success: function (data) {
            console.log(data);
            if(data == "200"){
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
    var result = null;
    $.ajax({
        url: "/remoteserver",
        dataType: "json",
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

function UpdateRemoteServer(id, ip, status) {
    var result = false;
    $.ajax({
        url: "/remoteserver",
        dataType: "json",
        type: "POST",
        async: false,
        data: {
            Id : id,
            Remoteserver : ip,
            Isconnected : status
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

class RemoteFileServer extends React.Component {

    constructor(props) {
        super(props);

        var result = GetRemoteServer();

        this.state = {
            Id: result.Id,
            Remoteserver: result.Remoteserver,
            Contentfolder:"UserContent",
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

    handleRemoteServerIpChange(e) {
        this.setState({ Remoteserver: e.target.value });
    }

    handleIsConnectedChange(e) {
        this.setState({ Isconnected: e.target.value });
    }

    handleCheckClick(e) {
        var url = "/testremoteserver";
        var isConnected = TestRemoteServerConnection(url);
        console.log("server connection :", isConnected);
        this.setState({
            Isconnected: isConnected
        });
        UpdateRemoteServer(this.state.Id, this.state.Remoteserver, isConnected);
    }

    render() {
        const remoteServer = this.state.Remoteserver;
        const isConnected = this.state.Isconnected;
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