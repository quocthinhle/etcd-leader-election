package election

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type Election struct {
	id  string
	cli *clientv3.Client
}

func NewElection(cli *clientv3.Client, id string) *Election {
	return &Election{
		id:  id,
		cli: cli,
	}
}

func (e *Election) ElectLeader() error {
	session, err := concurrency.NewSession(e.cli, concurrency.WithTTL(3))
	if err != nil {
		return err
	}

	election := concurrency.NewElection(session, "/leader-election/")

	fmt.Println("Im trying to be the leader")

	if err := election.Campaign(context.Background(), e.id); err != nil {
		return err
	}

	fmt.Println("I am the leader now")

	return nil
}
