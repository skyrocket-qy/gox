Key Components:
Peer Discovery: Peers need to find each other, which can be done through a simple server or list of peers.
File Sharing: Once a peer discovers others, they can request and send files directly to each other.
File Request Handling: The server will facilitate file requests between peers but will not store the files itself. Each peer stores and serves its own files.
Simple Approach:
Peer Server: A Go server to manage file requests from other peers.
Peer Client: Each peer can act as both a client (requesting files) and a server (serving files).
Discovery Mechanism: A simple broadcast of available files to other peers in the network.
File Transfer: Peers can request files from others and download them via a TCP connection.