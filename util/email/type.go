package email

type NexsoftWelcomingParam struct {
    Email string `field:"Email"`
    Name string `field:"Name"`
    CompanyName string `field:"CompanyName"`
}

type GrologWelcomingParam struct {
    UserID string `field:"UserID"`
    UserName string `field:"UserName"`
}
