package nats

import (
    "fmt"
    "github.com/nexsoft-git/nexcommon/util"
)

func NewNatsUtil(
    nexsoft NatsConnectionConfig,
    gromart NatsConnectionConfig,
) (
    NatsUtil,
    error,
) {
    result := &natsUtil{}

    connNexsoft, jsNexsoft, err := util.ConnectNatsJs(fmt.Sprintf("%s:%d", nexsoft.Host, nexsoft.Port))
    if err != nil {
        return nil, err
    }

    result.nexsoft = natsConnection{
        conn: connNexsoft,
        js:   jsNexsoft,
    }

    connGromart, jsGromart, err := util.ConnectNatsJs(fmt.Sprintf("%s:%d", gromart.Host, gromart.Port))
    if err != nil {
        return nil, err
    }

    result.gromart = natsConnection{
        conn: connGromart,
        js:   jsGromart,
    }

    err = util.CreateStream(
        result.nexsoft.js,
        "streamName",
        []string{
            "subject1",
        },
    )
    if err != nil {
        return nil, err
    }

    err = util.CreateStream(
        result.nexsoft.js,
        "streamName",
        []string{
            "subject2",
        },
    )
    if err != nil {
        return nil, err
    }

    err = util.CreateStream(
        result.gromart.js,
        "streamName",
        []string{
            "subject_gromart",
        },
    )
    if err != nil {
        return nil, err
    }

    return result, nil
}

type natsUtil struct {
    nexsoft natsConnection
    gromart natsConnection
}

func (n *natsUtil) PublishNexsoftSubject1(msg []byte) error {
    _, err := n.nexsoft.js.Publish("subject1", msg)
    return err
}

func (n *natsUtil) PublishNexsoftSubject2(msg []byte) error {
    _, err := n.nexsoft.js.Publish("subject2", msg)
    return err
}

func (n *natsUtil) PublishGromartSubject_gromart(msg []byte) error {
    _, err := n.gromart.js.Publish("subject_gromart", msg)
    return err
}

func (n *natsUtil) Close() {

    if n.nexsoft.conn != nil {
        n.nexsoft.conn.Close()
    }

    if n.gromart.conn != nil {
        n.gromart.conn.Close()
    }
}