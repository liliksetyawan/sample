package email

import (
    "github.com/nexsoft-git/nexcommon/util"
)

func NewEmailUtility(
    nexsoft util.EmailUtilityParam,
    grolog util.EmailUtilityParam,
) *EmailUtility {
    return &EmailUtility{
        nexsoftUtil: util.NewEmailUtility(nexsoft),
        grologUtil: util.NewEmailUtility(grolog),
    }
}

type EmailUtility struct {
    nexsoftUtil *util.EmailUtility
    grologUtil *util.EmailUtility
}

func (e *EmailUtility) SendNexsoftWelcomingEmail(
    param NexsoftWelcomingParam,
    receiver string,
    alias string,
) {
    message := e.GetNexsoftWelcomingTemplate()
    message = e.nexsoftUtil.CompleteEmail(&param, message)

    e.nexsoftUtil.NewEmailBuilder().
        Receiver(receiver).
        Subject("Welcome to nexsoft").
        HTMLMessage(message)
}

func (e *EmailUtility) SendNexsoftReminderEmail(
    receiver string,
    alias string,
) {
    message := e.GetNexsoftReminderTemplate()

    e.nexsoftUtil.NewEmailBuilder().
        Receiver(receiver).
        Subject("Don't forget about nexsoft").
        HTMLMessage(message)
}

func (e *EmailUtility) SendGrologWelcomingEmail(
    param GrologWelcomingParam,
    receiver string,
    alias string,
) {
    message := e.GetGrologWelcomingTemplate()
    message = e.grologUtil.CompleteEmail(&param, message)

    e.grologUtil.NewEmailBuilder().
        Receiver(receiver).
        Subject("Welcome to grolog").
        Message(message)
}

func (e *EmailUtility) SendGrologReminderEmail(
    receiver string,
    alias string,
) {
    message := e.GetGrologReminderTemplate()

    e.grologUtil.NewEmailBuilder().
        Receiver(receiver).
        Subject("Don't forget about grolog").
        Message(message)
}