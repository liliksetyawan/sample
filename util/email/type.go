package email

type NexsoftWelcomingParam struct {
    CompanyName string `field:"CompanyName"`
    Name string `field:"Name"`
    Email string `field:"Email"`
}

type GrologWelcomingParam struct {
    UserID string `field:"UserID"`
    UserName string `field:"UserName"`
}
