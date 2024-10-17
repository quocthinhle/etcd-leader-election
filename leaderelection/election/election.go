package election

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type Election struct {
	candidateId string
	cli         *clientv3.Client
}

func NewElection(cli *clientv3.Client, id string) *Election {
	return &Election{
		candidateId: id,
		cli:         cli,
	}
}

func (e *Election) ElectLeader() error {
	session, err := concurrency.NewSession(e.cli, concurrency.WithTTL(3))
	if err != nil {
		return err
	}

	election := concurrency.NewElection(session, "/leader-election")

	node, err := election.Leader(context.Background())
	if err != nil {
		return err
	}

	currentLeader := string(node.Kvs[0].Value)
	if currentLeader == e.candidateId {
		// Resign
	} else {
		fmt.Println("Current leader is: ", currentLeader)
	}

	if err := election.Campaign(context.Background(), e.candidateId); err != nil {
		return err
	}

	fmt.Println("I am the leader now with key: ", string(node.Kvs[0].Value))

	return nil
}
