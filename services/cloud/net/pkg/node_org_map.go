package pkg

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
)

const orgNetMappingKeyPrefix = "map:"

type NodeOrgMap struct {
	etcd *clientv3.Client
}
type OrgNet struct {
	Org     string
	Network string
}

func NewNodeToOrgMap(config *Config) *NodeOrgMap {
	client, err := clientv3.New(clientv3.Config{
		DialTimeout: config.DialTimeoutSecond,
		Endpoints:   []string{config.EtcdHost},
	})
	if err != nil {
		logrus.Fatalf("Cannot connect to etcd: %v", err)
	}

	return &NodeOrgMap{
		etcd: client,
	}
}

func (n *NodeOrgMap) Add(ctx context.Context, nodeId string, org string, network string) error {
	nodeIdKey := formatMappKey(nodeId)
	_, err := n.etcd.Put(ctx, nodeIdKey, org+"."+network)
	if err != nil {
		return fmt.Errorf("failed to add record to db. Error: %v", err)
	}
	return nil
}

func (n *NodeOrgMap) List(ctx context.Context) (map[string]OrgNet, error) {
	vals, err := n.etcd.Get(ctx, orgNetMappingKeyPrefix, clientv3.WithPrefix())

	if err != nil {
		return nil, fmt.Errorf("failed to get record from db. Error: %v", err)
	}

	res := map[string]OrgNet{}
	for _, val := range vals.Kvs {
		c := strings.Split(string(val.Value), ".")
		if len(c) != 2 {
			logrus.Errorf("failed to parse org.network structure for '%s' with value '%s'", string(val.Key), string(val.Value))
		}

		res[strings.TrimPrefix(string(val.Key), orgNetMappingKeyPrefix)] = OrgNet{
			Org:     c[0],
			Network: c[1],
		}
	}

	return res, nil
}

func formatMappKey(nodeId string) string {
	return orgNetMappingKeyPrefix + nodeId
}