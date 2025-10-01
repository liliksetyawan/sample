package email

type NexsoftWelcomingParam struct {
    Name string `field:"Name"`
    Email string `field:"Email"`
    CompanyName string `field:"CompanyName"`
}

type GrologWelcomingParam struct {
    UserID string `field:"UserID"`
    UserName string `field:"UserName"`
}
