package hubs

import "sync"

type connInfo struct {
	connectionId string
	userId       int64
}

type connectionManager struct {
	connections []connInfo
}

var locker sync.RWMutex

var ConnectionManager = &connectionManager{
	connections: make([]connInfo, 10),
}

func (c *connectionManager) Add(userId int64, connectionId string) {
	locker.Lock()
	defer locker.Unlock()
	c.connections = append(c.connections, connInfo{
		connectionId: connectionId,
		userId:       userId,
	})
}

func (c *connectionManager) Delete(connectionId string) {
	locker.Lock()
	defer locker.Unlock()
	for i, con := range c.connections {
		if con.connectionId == connectionId {
			c.connections = append(c.connections[0:i], c.connections[i+1:]...)
			break
		}
	}
}

func (c *connectionManager) GetConnectionIds(userId int64) []string {
	locker.RLock()
	defer locker.RUnlock()
	var connectionIds []string
	for _, con := range c.connections {
		if con.userId == userId {
			connectionIds = append(connectionIds, con.connectionId)
		}
	}
	return connectionIds
}
