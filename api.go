package main

// All the structs defined here are from https://core.telegram.org/bots/api#available-types
type (
	TelegramResponse struct {
		Ok          bool        `json:"ok"`
		Result      interface{} `json:"result"`
		Description string      `json:"description"`
		ErrorCode   int32       `json:"error_code"`
	}

	User struct {
		ID           int    `json:"id"`
		IsBot        *bool  `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	}

	Chat struct {
		ID                       int        `json:"id"`
		Type                     string     `json:"type"`
		Title                    string     `json:"title"`
		Username                 string     `json:"username"`
		FirstName                string     `json:"first_name"`
		LastName                 string     `json:"last_name"`
		AllMembersAdministrators *bool      `json:"all_members_are_administrators"`
		Photo                    *ChatPhoto `json:"photo"`
		Description              string     `json:"description"`
		InviteLink               string     `json:"invite_link"`
		PinnedMessage            *Message   `json:"pinned_message"`
		StickerSetName           string     `json:"sticker_set_name"`
		CanSetSticketSet         *bool      `json:"can_set_sticker_set"`
	}

	Message struct {
		ID                    int                `json:"message_id"`
		From                  *User              `json:"from"`
		Date                  int64              `json:"date"`
		Chat                  *Chat              `json:"chat,omitempty"`
		ForwardFrom           *User              `json:"forward_from,omitempty"`
		ForwardFromChat       *Chat              `json:"forward_from_chat,omitempty"`
		ForwardFromMessageID  int                `json:"forward_from_message_id"`
		ForwardSignature      string             `json:"forward_signature"`
		ForwardDate           int64              `json:"forward_date"`
		ReplyToMessage        *Message           `json:"reply_to_message,omitempty"`
		EditDate              int64              `json:"edit_date"`
		MediaGroupID          string             `json:"media_group_id"`
		AuthorSignature       string             `json:"author_signature"`
		Text                  string             `json:"text"`
		Entities              *[]MessageEntity   `json:"entities,omitempty"`
		CaptionEntities       *[]MessageEntity   `json:"caption_entities,omitempty"`
		Audio                 *Audio             `json:"audio,omitempty"`
		Document              *Document          `json:"document,omitempty"`
		Game                  *Game              `json:"game,omitempty"`
		Photo                 *[]PhotoSize       `json:"photo,omitempty"`
		Sticker               *Sticker           `json:"sticker,omitempty"`
		Video                 *Video             `json:"video,omitempty"`
		Voice                 *Voice             `json:"voice,omitempty"`
		VideoNote             *VideoNote         `json:"video_note,omitempty"`
		Caption               string             `json:"caption,omitempty"`
		Contact               *Contact           `json:"contact,omitempty"`
		Location              *Location          `json:"location,omitempty"`
		Venue                 *Venue             `json:"venue,omitempty"`
		NewChatMembers        *[]User            `json:"new_chat_members,omitempty"`
		NewChatMember         *User              `json:"new_chat_member,omitempty"`
		LeftChatMember        *User              `json:"left_chat_member,omitempty"`
		NewChatTitle          string             `json:"new_chat_title"`
		NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo"`
		DeleteChatPhoto       *bool              `json:"delete_chat_photo"`
		GroupChatCreated      *bool              `json:"group_chat_created"`
		SuperGroupChatCreated *bool              `json:"supergroup_chat_created"`
		ChannelChatCreated    *bool              `json:"channel_chat_created"`
		MigrateToChatID       int                `json:"migrate_to_chat_id,omitempty"`
		MigrateFromChatID     int                `json:"migrate_from_chat_id,omitempty"`
		PinnedMessage         *Message           `json:"pinned_message"`
		Invoice               *Invoice           `json:"invoice"`
		SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
		ConnectedWebsite      string             `json:"connected_website"`
	}

	MessageEntity struct {
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Length int    `json:"length"`
		URL    string `json:"url"`
		User   *User  `json:"user"`
	}

	PhotoSize struct {
		FileID   string `json:"file_id"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		FileSize int    `json:"file_size"`
	}

	Audio struct {
		FileID    string `json:"file_id"`
		Duration  int    `json:"duration"`
		Performer string `json:"performer"`
		Title     string `json:"title"`
		MimeType  string `json:"mime_type"`
		FileSize  int    `json:"file_size"`
	}

	Document struct {
		FileID   string     `json:"file_id"`
		Thumb    *PhotoSize `json:"thumb"`
		FileName string     `json:"file_name"`
		MimeType string     `json:"mime_type"`
		FileSize int        `json:"file_size"`
	}

	Video struct {
		FileID   string     `json:"file_id"`
		Width    int        `json:"width"`
		Height   int        `json:"height"`
		Duration int        `json:"duration"`
		Thumb    *PhotoSize `json:"thumb"`
		MimeType string     `json:"mime_type"`
		FileSize int        `json:"file_size"`
	}

	Voice struct {
		FileID   string `json:"file_id"`
		Duration int    `json:"duration"`
		MimeType string `json:"mime_type"`
		FileSize int    `json:"file_size"`
	}

	VideoNote struct {
		FileID   string     `json:"file_id"`
		Length   int        `json:"length"`
		Duration int        `json:"duration"`
		Thumb    *PhotoSize `json:"thumb"`
		MimeType string     `json:"mime_type"`
		FileSize int        `json:"file_size"`
	}

	Contact struct {
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		UserID      int    `json:"user_id"`
	}

	Location struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	}

	Venue struct {
		Location     *Location `json:"location"`
		Title        string    `json:"title"`
		Address      string    `json:"address"`
		FoursquareID string    `json:"foursquare_id"`
	}

	UserProfilePhotos struct {
		TotalCount int          `json:"total_count"`
		Photos     *[]PhotoSize `json:"photos"`
	}

	ReplyKeyboardMarkup struct {
		Keyboard        [][]KeyboardButton `json:"keyboard"`
		ResizeKeyboard  *bool              `json:"resize_keyboard,omitempty"`
		OneTimeKeyboard *bool              `json:"one_time_keyboard,omitempty"`
		Selective       *bool              `json:"selective,omitempty"`
	}

	KeyboardButton struct {
		Text            string `json:"text"`
		RequestContact  *bool  `json:"request_contact,omitempty"`
		RequestLocation *bool  `json:"request_location,omitempty"`
	}

	ReplyKeyboardRemove struct {
		RemoveKeyboard *bool `json:"remove_keyboard,omitempty"`
		Selective      *bool `json:"selective,omitempty"`
	}

	InlineKeyboardMarkup struct {
		InlineKeyboard []InlineKeyboardButton `json:"inline_keyboard"`
	}

	InlineKeyboardButton struct {
		Text                         string        `json:"text"`
		URL                          string        `json:"url"`
		CallbackData                 string        `json:"callback_data"`
		SwitchInlineQuery            string        `json:"switch_inline_query"`
		SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"`
		CallbackGame                 *CallbackGame `json:"callback_game"`
		Pay                          *bool         `json:"pay"`
	}

	CallbackQuery struct {
		ID              string   `json:"id"`
		From            *User    `json:"from"`
		Message         *Message `json:"message"`
		InlineMessageID string   `json:"inline_message_id"`
		ChatInstance    string   `json:"chat_instance"`
		Data            string   `json:"data"`
		GameShortName   string   `json:"game_short_name"`
	}

	ForceReply struct {
		ForceReply *bool `json:"force_reply"`
		Selective  *bool `json:"selective"`
	}

	ChatPhoto struct {
		SmallFileID string `json:"small_file_id"`
		BigFileID   string `json:"big_file_id"`
	}

	ChatMember struct {
		User                  *User  `json:"user"`
		Status                string `json:"status"`
		UntilDate             int64  `json:"until_date,omitempty"`
		CanBeEdited           *bool  `json:"can_be_edited,omitempty"`
		CanChangeInfo         *bool  `json:"can_change_info,omitempty"`
		CanPostMessages       *bool  `json:"can_post_messages,omitempty"`
		CanEditMessages       *bool  `json:"can_edit_messages,omitempty"`
		CanDeleteMessages     *bool  `json:"can_delete_messages,omitempty"`
		CanInviteUsers        *bool  `json:"can_invite_users,omitempty"`
		CanRestrictMembers    *bool  `json:"can_restrict_members,omitempty"`
		CanPinMessages        *bool  `json:"can_pin_messages,omitempty"`
		CanPromoteMembers     *bool  `json:"can_promote_members,omitempty"`
		CanSendMessages       *bool  `json:"can_send_messages,omitempty"`
		CanSendMediaMessages  *bool  `json:"can_send_media_messages,omitempty"`
		CanSendOtherMessages  *bool  `json:"can_send_other_messages,omitempty"`
		CanAddWebPagePreviews *bool  `json:"can_add_web_page_previews,omitempty"`
	}

	ResponseParameters struct {
		MigrateToChatID int64 `json:"migrate_to_chat_id"`
		RetryAfter      int   `json:"retry_after"`
	}

	InputMediaPhoto struct {
		Type      string `json:"type"`
		Media     string `json:"media"`
		Caption   string `json:"caption"`
		ParseMode string `json:"parse_mode"`
	}

	InputMediaVideo struct {
		Type              string `json:"type"`
		Media             string `json:"media"`
		Caption           string `json:"caption"`
		ParseMode         string `json:"parse_mode"`
		Width             int    `json:"width"`
		Height            int    `json:"height"`
		Duration          int    `json:"duration"`
		SupportsStreaming *bool  `json:"supports_streaming"`
	}

	Game struct {
		Title       string           `json:"title"`
		Description string           `json:"description"`
		Photo       *[]PhotoSize     `json:"photo"`
		Text        string           `json:"text"`
		TextEntries *[]MessageEntity `json:"text_entities"`
		Animation   *Animation       `json:"animation"`
	}

	Animation struct {
		FileID   string     `json:"file_id"`
		Thumb    *PhotoSize `json:"thumb"`
		FileName string     `json:"file_name"`
		MimeType string     `json:"mime_type"`
		FileSize int        `json:"file_size"`
	}

	Sticker struct {
		FileID       string        `json:"file_id"`
		Width        int           `json:"width"`
		Height       int           `json:"height"`
		Thumb        *PhotoSize    `json:"thumb"`
		Emoji        string        `json:"emoji"`
		SetName      string        `json:"set_name"`
		MaskPosition *MaskPosition `json:"mask_position"`
		FileSize     int           `json:"file_size"`
	}

	StickerSet struct {
		Name          string     `json:"name"`
		Title         string     `json:"title"`
		ContainsMasks *bool      `json:"contains_masks"`
		Stickers      *[]Sticker `json:"stickers"`
	}

	MaskPosition struct {
		Point  string  `json:"point"`
		XShift float32 `json:"x_shift"`
		YShift float32 `json:"y_shift"`
		Scale  float32 `json:"scale"`
	}

	Invoice struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		StartParameter string `json:"start_parameter"`
		Currency       string `json:"currency"`
		TotalAmount    int    `json:"total_amount"`
	}

	SuccessfulPayment struct {
		Currency                string     `json:"currency"`
		TotalAmount             int        `json:"total_amount"`
		InvoicePayload          string     `json:"invoice_payload"`
		ShippingOptionID        string     `json:"shipping_option_id"`
		OrderInfo               *OrderInfo `json:"order_info"`
		TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
		ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
	}

	OrderInfo struct {
		Name            string           `json:"name"`
		PhoneNumber     string           `json:"phone_number"`
		Email           string           `json:"email"`
		ShippingAddress *ShippingAddress `json:"shipping_address"`
	}
	ShippingAddress struct {
		CountryCode string `json:"country_code"`
		State       string `json:"state"`
		City        string `json:"city"`
		StreetLine1 string `json:"street_line_1"`
		StreetLine2 string `json:"street_line_2"`
		PostCode    string `json:"post_code"`
	}

	CallbackGame struct {
		UserID             int   `json:"user_id"`
		Score              int   `json:"score"`
		Force              *bool `json:"force"`
		DisableEditMessage *bool `json:"disable_edit_message"`
		ChatID             int   `json:"chat_id"`
		MessageID          int   `json:"message_id"`
		InlineMessageID    int   `json:"inline_message_id"`
	}

	Update struct {
		UpdateID           int                 `json:"update_id"`
		Message            *Message            `json:"message"`
		EditedMessage      *Message            `json:"edited_message"`
		ChannelPost        *Message            `json:"channel_post"`
		EditedChannelPost  *Message            `json:"edited_channel_post"`
		InlineQuery        *InlineQuery        `json:"inline_query"`
		ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
		CallbackQuery      *CallbackQuery      `json:"callback_query"`
		ShippingQuery      *ShippingQuery      `json:"shipping_query"`
		PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
	}

	InlineQuery struct {
		ID       string    `json:"id"`
		From     *User     `json:"from"`
		Location *Location `json:"location"`
		Query    string    `json:"query"`
		Offset   string    `json:"offset"`
	}

	ChosenInlineResult struct {
		ID       string    `json:"id"`
		From     *User     `json:"from"`
		Location *Location `json:"location"`
		Query    string    `json:"query"`
		Offset   string    `json:"offset"`
	}

	ShippingQuery struct {
		ID              string           `json:"id"`
		From            *User            `json:"from"`
		InvoicePayload  string           `json:"invoice_payload"`
		ShippingAddress *ShippingAddress `jsong:"shipping_address"`
	}

	PreCheckoutQuery struct {
		ID               string     `json:"id"`
		From             *User      `json:"from"`
		Currency         string     `json:"currency"`
		TotalAmount      int        `json:"total_amount"`
		InvoicePayload   string     `json:"invoice_payload"`
		ShippingOptionID string     `json:"shipping_option_id"`
		OrderInfo        *OrderInfo `json:"order_info"`
	}

	MessageOut struct {
		ChatID                int         `json:"chat_id"`
		Text                  string      `json:"text"`
		ParseMode             string      `json:"parse_mode"`
		DisableWebPagePreview *bool       `json:"disable_web_page_preview,omitempty"`
		DisableNotification   *bool       `json:"disable_notification,omitemptyn"`
		ReplyToMessageID      int         `json:"reply_to_message_id,omitempty"`
		ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	}
)
