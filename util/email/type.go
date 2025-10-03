package email

type NexsoftWelcomingParam struct {
    Email string `field:"Email"`
    Name string `field:"Name"`
    CompanyName string `field:"CompanyName"`
}

type GrologWelcomingParam struct {
    UserName string `field:"UserName"`
    UserID string `field:"UserID"`
}
