package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Update interface {
	Type() UpdateType
	BindOperator()
	Validate() Operator
}

type NopUpdate struct{}

func (NopUpdate) Type() UpdateType {
	return updateTypeNotSupported
}

func (NopUpdate) BindOperator() {
	log.Println("update received but no need to handle further")
}

func (NopUpdate) Validate() Operator {
	return new(OperatorForNop)
}

type BotInChatStatusChangedUpdate struct {
}

func (update BotInChatStatusChangedUpdate) Type() UpdateType {
	return updateTypeBotSelfChanged
}

func (update BotInChatStatusChangedUpdate) BindOperator() {
}

func (update BotInChatStatusChangedUpdate) Validate() Operator {
	var operator OperatorForBotInChatStatusChangedUpdate
	operator.SetOperatorFunc()
	return &operator
}

type MemberAddUpdate struct {
	Inviter *tgbotapi.User
	Users   []tgbotapi.User
}

func (MemberAddUpdate) Type() UpdateType {
	return updateTypeMemberAdded
}

func (MemberAddUpdate) BindOperator() {
}

func (MemberAddUpdate) Validate() Operator {
	var operator OperatorForMemberAddUpdate
	operator.SetOperatorFunc()
	return &operator
}
