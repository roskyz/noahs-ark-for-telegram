package main

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var handledUpdates = []Update{
	NopUpdate{}, BotInChatStatusChangedUpdate{}, MemberAddUpdate{},
}

// only have reads once init
var handledOperators = func() map[UpdateType]Operator {
	var operators = make(map[UpdateType]Operator, len(handledUpdates))
	for _, update := range handledUpdates {
		operators[update.Type()] = update.Validate()
	}
	return operators
}()

func (b *Bot) handleUpdate(ctx context.Context, tgUpdate *tgbotapi.Update) {
	defer b.wg.Done()
	operator := handledOperators[whichUpdateType(tgUpdate)]
	update := operator.NewUpdatePointer(tgUpdate)
	if err := operator.DO(ctx, update); err != nil {
		log.Printf("[def-error]: %v\n", err)
	}
}

type UpdateType int

const (
	updateTypeNotSupported UpdateType = iota
	updateTypeMemberAdded
	updateTypeMemberLeft
	updateTypeMessageRecv
	updateTypeBotSelfChanged
)

func whichUpdateType(tgUpdate *tgbotapi.Update) UpdateType {
	if tgUpdate.Message != nil {
		if len(tgUpdate.Message.NewChatMembers) > 0 {
			return updateTypeMemberAdded
		}
		if tgUpdate.Message.LeftChatMember != nil {
			return updateTypeMemberLeft
		}
	}

	if tgUpdate.Message != nil || tgUpdate.EditedMessage != nil {
		return updateTypeMessageRecv
	}

	if tgUpdate.MyChatMember != nil {
		return updateTypeBotSelfChanged
	}

	return updateTypeNotSupported
}
