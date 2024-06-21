## Decentralized File Storage System in Golang

This repository implements a decentralized file storage system built with Golang. The project is designed to be educational and provides a foundational understanding of how to build such a system.

**Getting Started**

This project requires some familiarity with Golang concepts and TCP networking. 

**Learning Objectives:**

* Build a basic TCP library in Golang.
* Implement functionalities for storing and retrieving files in a decentralized network.
* Integrate encryption for secure file storage.

**Project Breakdown:**

The project is divided into several parts, each focusing on a specific aspect of the decentralized file storage system:

1. **TCP Library:** We will start by building a foundation with a custom TCP library in Golang. This library will handle communication between peers in the network.
2. **File Storage:** We will implement functionalities to store and retrieve files on the local system. The system will also handle retrieving files from other peers if they are not available locally.
3. **Encryption:** To ensure secure storage of files on peers, we will integrate encryption functionalities into the system.
4. **Further Exploration (Optional):** Consider implementing a function to delete files. This function should not only delete the file locally but also broadcast a message to the network instructing other peers to do the same.

**Resources:**

* The video tutorial for this project can be found here: [link to the video](https://www.youtube.com/watch?v=bymQakvTY40) (This readme assumes the video is the primary resource)
