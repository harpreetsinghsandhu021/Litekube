package node

// Represents a node in the cluster, which can be a machine or a virtual machine.
// It containts information about the node's resources and its current state.
type Node struct {
	Name          string // Unique identifier for the node
	Ip            string // IP address of the node, which can be used to communicate with it
	Cores         int    // Number of CPU cores available on the node
	Memory        int    // The amount of RAM available
	Disk          int    // The amount of disk space available on the node
	DiskAllocated int    // The amount of disk space allocated to tasks, in bytes
	Role          string // The role of the node, such as "worker" or "manager"
	TaskCount     int    // The number of tasks currently running on the node
}
