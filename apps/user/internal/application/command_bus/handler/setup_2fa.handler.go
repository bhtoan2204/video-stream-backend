package handler

import (
	"context"
	"encoding/base64"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

type Setup2FACommandHandler struct {
	userService        interfaces.UserServiceInterface
	userSettingService interfaces.UserSettingServiceInterface
}

func NewSetup2FACommandHandler(
	userService interfaces.UserServiceInterface,
	userSettingService interfaces.UserSettingServiceInterface,
) *Setup2FACommandHandler {
	return &Setup2FACommandHandler{
		userService: userService,
	}
}

func (h *Setup2FACommandHandler) Handle(ctx context.Context, cmd *command.Setup2FACommand) (*command.Setup2FACommandResult, error) {
	data, err := h.userService.GetUserById(ctx, &command.GetUserByIdCommand{ID: cmd.UserID})
	if err != nil {
		return nil, err
	}

	opts := totp.GenerateOpts{
		Issuer:      "User",
		AccountName: data.Result.Email,
	}

	key, err := totp.Generate(opts)
	if err != nil {
		return nil, err
	}

	updateUserSettingCmd := &command.UpdateUserSettingsCommand{
		UserID:       cmd.UserID,
		Is2FAEnabled: func(b bool) *bool { return &b }(true),
	}

	_, err = h.userSettingService.UpdateByUserId(ctx, cmd.UserID, updateUserSettingCmd)
	if err != nil {
		return nil, err
	}

	otpURL := key.URL()
	qrCode, err := qrcode.Encode(otpURL, qrcode.Medium, 256)

	if err != nil {
		return nil, err
	}

	base64QR := base64.StdEncoding.EncodeToString(qrCode)

	return &command.Setup2FACommandResult{
		OTPUrl: otpURL,
		QRCode: base64QR,
	}, nil
}
