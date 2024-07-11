package main

import (
	"awesomeProject/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Amount  int
}

func (cmd *Command) Do() error {
	switch cmd.Cmd {
	case "create":
		return cmd.create()
	case "get":
		return cmd.get()
	case "delete":
		return cmd.delete()
	case "change-amount":
		return cmd.changeAmount()
	case "change-name":
		return cmd.changeName()
	default:
		return fmt.Errorf("unknown command: %s", cmd.Cmd)
	}
}

func (cmd *Command) create() error {
	conn, err := grpc.NewClient(fmt.Sprintf("0.0.0.0:%v", cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()
	c := proto.NewHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.CreateAccount(ctx, &proto.CreateAccountRequest{Name: cmd.Name, Amount: int32(cmd.Amount)})
	if err != nil {
		return err
	}

	return nil
}

func (cmd *Command) get() error {
	conn, err := grpc.NewClient(fmt.Sprintf("0.0.0.0:%v", cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()
	c := proto.NewHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.GetAccount(ctx, &proto.GetAccountRequest{Name: cmd.Name})
	if err != nil {
		return err
	}

	fmt.Printf("Account name and amount: %s, %v", res.Name, res.Amount)

	return nil
}

func (cmd *Command) delete() error {
	conn, err := grpc.NewClient(fmt.Sprintf("0.0.0.0:%v", cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()
	c := proto.NewHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: cmd.Name})
	if err != nil {
		return err
	}

	return nil
}

func (cmd *Command) changeAmount() error {
	conn, err := grpc.NewClient(fmt.Sprintf("0.0.0.0:%v", cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()
	c := proto.NewHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.ChangeAccountAmount(ctx, &proto.ChangeAmountRequest{Name: cmd.Name, NewAmount: int32(cmd.Amount)})
	if err != nil {
		return err
	}

	return nil
}

func (cmd *Command) changeName() error {
	conn, err := grpc.NewClient(fmt.Sprintf("0.0.0.0:%v", cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()
	c := proto.NewHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.ChangeAccountName(ctx, &proto.ChangeNameRequest{Name: cmd.Name, NewName: cmd.NewName})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newNameVal := flag.String("new-name", "", "new name of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NewName: *newNameVal,
		Amount:  *amountVal,
	}

	if err := cmd.Do(); err != nil {
		panic(err)
	}
}
