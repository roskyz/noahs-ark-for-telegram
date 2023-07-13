package main

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Operator interface {
	NewUpdatePointer(tgUpdate *tgbotapi.Update) Update
	DO(ctx context.Context, update Update) error
	SetOperatorFunc()
}

type OperatorForNop struct { /* do nothing only print some info */
}

func (operator *OperatorForNop) NewUpdatePointer(*tgbotapi.Update) Update {
	return new(NopUpdate)
}

func (operator *OperatorForNop) DO(ctx context.Context, update Update) error {
	defer update.BindOperator()
	return nil
}

func (operator *OperatorForNop) SetOperatorFunc() {}

type OperatorForBotInChatStatusChangedUpdate struct {
	handleBotInChatStatusChangedUpdate func(ctx context.Context, update *BotInChatStatusChangedUpdate) error
}

func (operator *OperatorForBotInChatStatusChangedUpdate) NewUpdatePointer(tgUpdate *tgbotapi.Update) Update {
	//TODO implement me
	panic("implement me")
}

func (operator *OperatorForBotInChatStatusChangedUpdate) DO(ctx context.Context, update Update) error {
	return operator.handleBotInChatStatusChangedUpdate(ctx, update.(*BotInChatStatusChangedUpdate))
}

func (operator *OperatorForBotInChatStatusChangedUpdate) SetOperatorFunc() {
	operator.handleBotInChatStatusChangedUpdate = func(ctx context.Context, update *BotInChatStatusChangedUpdate) error {
		return nil
	}
}

type OperatorForMemberAddUpdate struct {
	handleMemberAddUpdate func(ctx context.Context, update *MemberAddUpdate) error
}

func (*OperatorForMemberAddUpdate) NewUpdatePointer(tgUpdate *tgbotapi.Update) Update {
	return new(MemberAddUpdate)
}

func (operator *OperatorForMemberAddUpdate) DO(ctx context.Context, update Update) error {
	defer update.BindOperator()
	return operator.handleMemberAddUpdate(ctx, update.(*MemberAddUpdate))
}

func (operator *OperatorForMemberAddUpdate) SetOperatorFunc() {
	operator.handleMemberAddUpdate = func(ctx context.Context, update *MemberAddUpdate) error {
		return nil
	}
}
