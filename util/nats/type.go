package nats

type NatsUtil interface {
    PublishNexsoftSubject1(msg []byte) error
    PublishNexsoftSubject2(msg []byte) error
    PublishGromartSubject_gromart(msg []byte) error
    Close()
}

type natsConnection struct {
    conn *nats.Conn
    js   nats.JetStreamContext
}

type NatsConnectionConfig struct {
    Host string
    Port int
}
